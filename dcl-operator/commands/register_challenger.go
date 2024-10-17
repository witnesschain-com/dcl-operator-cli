package operator_commands

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	dcl_common "github.com/witnesschain-com/dcl-operator-cli/common"
	"github.com/witnesschain-com/dcl-operator-cli/common/bindings/ChallengerRegistry"
	operator_config "github.com/witnesschain-com/dcl-operator-cli/dcl-operator/config"
	"github.com/witnesschain-com/diligencewatchtower-client/keystore"
	op_common "github.com/witnesschain-com/operator-cli/common"

	"github.com/urfave/cli/v2"
)

func RegisterChallengerCmd() *cli.Command {
	var registerChallengerCmd = &cli.Command{
		Name:  "registerChallenger",
		Usage: "Register a challenger",
		Flags: []cli.Flag{
			&op_common.ConfigPathFlag,
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.Value("config-file") == "" {
				cCtx.Set("config-file", dcl_common.DefaultOpChallengerConfig)
			}
			config := operator_config.GetChallengerConfigFromContext(cCtx)
			RegisterChallenger(config)
			return nil
		},
	}
	return registerChallengerCmd
}

func RegisterChallenger(config *operator_config.OperatorConfig) {
	var client *ethclient.Client
	client, config.ChainID = op_common.ConnectToUrl(config.EthRPCUrl)

	challengerRegistry, err := ChallengerRegistry.NewChallengerRegistry(dcl_common.NetworkConfig[config.ChainID.String()].ChallengerRegistryAddress, client)
	op_common.CheckError(err, "Instantiating ChallengerRegistry contract failed")

	vc := &keystore.VaultConfig{Address: config.OperatorAddress, ChainID: config.ChainID, PrivateKey: config.OperatorPrivateKey, Endpoint: config.Endpoint}
	operatorVault, err := keystore.SetupVault(vc)
	if err != nil {
		op_common.CheckError(err, "unable to setup vault")
	}

	if !dcl_common.IsOperatorWhitelisted(config.OperatorAddress, challengerRegistry) {
		fmt.Printf("Operator %s is not allow listed\n", config.OperatorAddress.Hex())
		return
	}

	transactOpts := operatorVault.NewTransactOpts(config.ChainID)

	if (dcl_common.NetworkConfig[config.ChainID.String()].GasPrice == -1) {
		transactOpts.GasPrice = big.NewInt(0)
	}

	expiry := op_common.CalculateExpiry(client, config.ExpiryInDays)

	for i, challengerAddress := range config.ChallengerAddresses {

		fmt.Println("challengerAddress: " + challengerAddress.Hex())

		var challengerPrivateKey *ecdsa.PrivateKey
		if len(config.ChallengerPrivateKeys) != 0 {
			challengerPrivateKey = config.ChallengerPrivateKeys[i]
		}

		vc := &keystore.VaultConfig{Address: challengerAddress, ChainID: config.ChainID, PrivateKey: challengerPrivateKey, Endpoint: config.Endpoint}
		challengerVault, err := keystore.SetupVault(vc)
		op_common.CheckError(err, "unable to setup challenger vault")

		if dcl_common.IsValidChallenger(challengerAddress, challengerRegistry) {
			fmt.Printf("challenger %s is already registered\n", challengerAddress.Hex())
			continue
		}

		salt := op_common.GenerateSalt()

		challengerSignature := GetChallengerSignature(client, challengerRegistry, challengerAddress, challengerVault, config.OperatorAddress, salt, expiry)

		regTx, err := challengerRegistry.RegisterChallenger(transactOpts, challengerAddress, salt, expiry, challengerSignature)
		op_common.CheckError(err, "Registering challenger failed")
		fmt.Printf("Tx sent: %s/tx/%s\n", dcl_common.NetworkConfig[config.ChainID.String()].BlockExplorer ,regTx.Hash().Hex())
		op_common.WaitForTransactionReceipt(client, regTx, config.TxReceiptTimeout)
	}
}
