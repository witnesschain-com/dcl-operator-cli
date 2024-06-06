package dcl_operator_config

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	op_common "github.com/witnesschain-com/operator-cli/common"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

type OperatorConfig struct {
	ProverPrivateKeys     []string       `json:"prover_private_keys"`
	OperatorPrivateKey    string         `json:"operator_private_key"`
	ProverRegistryAddress common.Address `json:"prover_registry_address"`
	EthRPCUrl             string         `json:"eth_rpc_url"`
	ChainId               big.Int        `json:"chain_id"`
	GasLimit              uint64         `json:"gas_limit"`
	TxReceiptTimeout      int64          `json:"tx_receipt_timeout"`
	ExpiryInDays          int64          `json:"expiry_in_days"`
	UseEncryptedKeys      bool           `json:"use_encrypted_keys"`
}

func GetConfigFromContext(cCtx *cli.Context) *OperatorConfig {
	configFilePath := cCtx.String("config-file")
	fmt.Printf("Using config file path : %s\n", configFilePath)

	data, err := os.ReadFile(configFilePath)
	op_common.CheckError(err, "Error reading json file")

	// Parse the json data into a struct
	var config OperatorConfig
	err = json.Unmarshal(data, &config)
	op_common.CheckError(err, "Error unmarshaling json data")

	SetDefaultConfigValues(&config)

	if config.UseEncryptedKeys {
		op_common.UseEncryptedKeys()
	}

	return &config
}

func SetDefaultConfigValues(config *OperatorConfig) {
	if config.ExpiryInDays == 0 {
		config.ExpiryInDays = 1 // 1 day
	}

	if config.TxReceiptTimeout == 0 {
		config.TxReceiptTimeout = 5 * 60 // 5 minutes
	}

	if config.GasLimit == 0 {
		config.GasLimit = 300000
	}
}
