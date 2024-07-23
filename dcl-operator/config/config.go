package dcl_operator_config

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"

	op_common "github.com/witnesschain-com/operator-cli/common"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

type OperatorConfig struct {
	ProverPrivateKeysHex     []string         `json:"prover_private_keys"`
	ProverAddresses          []common.Address `json:"prover_addresses"`
	ProverEncryptedKeys      []string         `json:"prover_encrypted_keys"`
	ChallengerPrivateKeysHex []string         `json:"challenger_private_keys"`
	ChallengerAddresses      []common.Address `json:"challenger_addresses"`
	ChallengerEncryptedKeys  []string         `json:"challenger_encrypted_keys"`
	OperatorPrivateKeyHex    string           `json:"operator_private_key"`
	OperatorAddress          common.Address   `json:"operator_address"`
	OperatorEncryptedKey     string           `json:"operator_encrypted_key"`
	EthRPCUrl                string           `json:"eth_rpc_url"`
	GasLimit                 uint64           `json:"gas_limit"`
	TxReceiptTimeout         uint64           `json:"tx_receipt_timeout"`
	ExpiryInDays             uint64           `json:"expiry_in_days"`
	Endpoint                 string           `json:"external_signer_endpoint"`
	KeyType                  string           `json:"encrypted_key_type"`
	ProverPrivateKeys        []*ecdsa.PrivateKey
	ChallengerPrivateKeys    []*ecdsa.PrivateKey
	OperatorPrivateKey       *ecdsa.PrivateKey
	ChainID                  *big.Int
}

func GetChallengerConfigFromContext(cCtx *cli.Context) *OperatorConfig {
	configFilePath := cCtx.String("config-file")
	fmt.Printf("Using config file path : %s\n", configFilePath)

	data, err := os.ReadFile(configFilePath)
	op_common.CheckError(err, "Error reading json file")

	// Parse the json data into a struct
	var config OperatorConfig = OperatorConfig{ExpiryInDays: 1, TxReceiptTimeout: 300, GasLimit: 300000}
	err = json.Unmarshal(data, &config)
	op_common.CheckError(err, "Error unmarshaling json data")

	if len(config.ChallengerEncryptedKeys) != 0 {
		// get the path from the first key, as others should be same
		// will not work with different paths
		op_common.RetryMounting()
		op_common.ProcessConfigKeyPath(config.ChallengerEncryptedKeys[0], config.KeyType)
		op_common.UseEncryptedKeys(config.KeyType)
	}

	if len(config.ChallengerPrivateKeysHex) != 0 {
		for _, privKey := range config.ChallengerPrivateKeysHex {
			fmt.Println(privKey)
			key, err := crypto.HexToECDSA(privKey)
			op_common.CheckError(err, "unable to convert challenger privatekey")
			config.ChallengerAddresses = append(config.ChallengerAddresses, crypto.PubkeyToAddress(key.PublicKey))
			config.ChallengerPrivateKeys = append(config.ChallengerPrivateKeys, key)
		}
	}

	if len(config.ChallengerEncryptedKeys) != 0 {
		for _, keyPath := range config.ChallengerEncryptedKeys {
			privKey, err := op_common.LoadPrivateKey(keyPath, config.KeyType)
			op_common.CheckError(err, "unable to load challenger encrypted keys")

			config.ChallengerPrivateKeys = append(config.ChallengerPrivateKeys, privKey)
			config.ChallengerAddresses = append(config.ChallengerAddresses, crypto.PubkeyToAddress(privKey.PublicKey))
		}
	}

	if len(config.OperatorEncryptedKey) != 0 {
		priv, err := op_common.LoadPrivateKey(config.OperatorEncryptedKey, config.KeyType)
		if err != nil {
			log.Fatal("unable to retive operator privateKey")
		}
		config.OperatorAddress = crypto.PubkeyToAddress(priv.PublicKey)
		config.OperatorPrivateKey = priv
	}

	if len(config.OperatorPrivateKeyHex) != 0 {
		priv, err := crypto.HexToECDSA(config.OperatorPrivateKeyHex)
		op_common.CheckError(err, "unable to convert operator privateKey")
		config.OperatorAddress = crypto.PubkeyToAddress(priv.PublicKey)
		config.OperatorPrivateKey = priv
	}

	if config.OperatorAddress.Cmp(common.Address{0}) == 0 {
		panic("operatorAddress is zero")
	}

	SetDefaultValues(&config)

	return &config
}

func GetProverConfigFromContext(cCtx *cli.Context) *OperatorConfig {
	configFilePath := cCtx.String("config-file")
	fmt.Printf("Using config file path : %s\n", configFilePath)

	data, err := os.ReadFile(configFilePath)
	op_common.CheckError(err, "Error reading json file")

	// Parse the json data into a struct
	var config OperatorConfig = OperatorConfig{ExpiryInDays: 1, TxReceiptTimeout: 300, GasLimit: 300000}
	err = json.Unmarshal(data, &config)
	op_common.CheckError(err, "Error unmarshaling json data")

	if len(config.ProverEncryptedKeys) != 0 {
		// get the path from the first key, as others should be same
		// will not work with different paths
		op_common.RetryMounting()
		op_common.ProcessConfigKeyPath(config.ProverEncryptedKeys[0], config.KeyType)
		op_common.UseEncryptedKeys(config.KeyType)
	}

	if len(config.ProverPrivateKeysHex) != 0 {
		for _, privKey := range config.ProverPrivateKeysHex {
			fmt.Println(privKey)
			key, err := crypto.HexToECDSA(privKey)
			op_common.CheckError(err, "unable to convert prover privatekey")
			config.ProverAddresses = append(config.ProverAddresses, crypto.PubkeyToAddress(key.PublicKey))
			config.ProverPrivateKeys = append(config.ProverPrivateKeys, key)
		}
	}

	if len(config.ProverEncryptedKeys) != 0 {
		for _, keyPath := range config.ProverEncryptedKeys {
			privKey, err := op_common.LoadPrivateKey(keyPath, config.KeyType)
			op_common.CheckError(err, "unable to load prover encrypted keys")

			config.ProverPrivateKeys = append(config.ProverPrivateKeys, privKey)
			config.ProverAddresses = append(config.ProverAddresses, crypto.PubkeyToAddress(privKey.PublicKey))
		}
	}

	if len(config.OperatorEncryptedKey) != 0 {
		priv, err := op_common.LoadPrivateKey(config.OperatorEncryptedKey, config.KeyType)
		if err != nil {
			log.Fatal("unable to retive operator privateKey")
		}
		config.OperatorAddress = crypto.PubkeyToAddress(priv.PublicKey)
		config.OperatorPrivateKey = priv
	}

	if len(config.OperatorPrivateKeyHex) != 0 {
		priv, err := crypto.HexToECDSA(config.OperatorPrivateKeyHex)
		op_common.CheckError(err, "unable to convert operator privateKey")
		config.OperatorAddress = crypto.PubkeyToAddress(priv.PublicKey)
		config.OperatorPrivateKey = priv
	}

	if config.OperatorAddress.Cmp(common.Address{0}) == 0 {
		panic("operatorAddress is zero")
	}

	SetDefaultValues(&config)

	return &config
}

func SetDefaultValues(config *OperatorConfig) {
	if config.GasLimit == 0 {
		config.GasLimit = op_common.DefaultGasLimit
	}

	if config.TxReceiptTimeout == 0 {
		config.TxReceiptTimeout = op_common.DefaultTxReceiptTimeout
	}

	if config.ExpiryInDays == 0 {
		config.ExpiryInDays = op_common.DefaultExpiration
	}

	if config.KeyType == "" {
		config.KeyType = op_common.KeyTypeW3SecretKey
	}
}
