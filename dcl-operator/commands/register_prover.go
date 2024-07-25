package operator_commands

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	dcl_common "github.com/witnesschain-com/dcl-operator-cli/common"
	"github.com/witnesschain-com/dcl-operator-cli/common/bindings/ProverRegistry"
	operator_config "github.com/witnesschain-com/dcl-operator-cli/dcl-operator/config"
	"github.com/witnesschain-com/diligencewatchtower-client/keystore"
	op_common "github.com/witnesschain-com/operator-cli/common"

	"github.com/urfave/cli/v2"
)

func RegisterProverCmd() *cli.Command {
	var registerProverCmd = &cli.Command{
		Name:  "registerProver",
		Usage: "Register a prover",
		Flags: []cli.Flag{
			&op_common.ConfigPathFlag,
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.Value("config-file") == "" {
				cCtx.Set("config-file", dcl_common.DefaultOpProverConfig)
			}
			config := operator_config.GetProverConfigFromContext(cCtx)
			RegisterProver(config)
			return nil
		},
	}
	return registerProverCmd
}

func RegisterProver(config *operator_config.OperatorConfig) {
	var client *ethclient.Client
	client, config.ChainID = op_common.ConnectToUrl(config.EthRPCUrl)

	proverRegistry, err := ProverRegistry.NewProverRegistry(dcl_common.NetworkConfig[config.ChainID.String()].ProverRegistryAddress, client)
	op_common.CheckError(err, "Instantiating ProverRegistry contract failed")

	vc := &keystore.VaultConfig{Address: config.OperatorAddress, ChainID: config.ChainID, PrivateKey: config.OperatorPrivateKey, Endpoint: config.Endpoint}
	operatorVault, err := keystore.SetupVault(vc)
	if err != nil {
		op_common.CheckError(err, "unable to setup vault")
	}

	if !dcl_common.IsOperatorWhitelisted(config.OperatorAddress, proverRegistry) {
		fmt.Printf("Operator %s is not allow listed\n", config.OperatorAddress.Hex())
		return
	}

	transactOpts := operatorVault.NewTransactOpts(config.ChainID)

	expiry := op_common.CalculateExpiry(client, config.ExpiryInDays)

	for i, proverAddress := range config.ProverAddresses {

		fmt.Println("proverAddress: " + proverAddress.Hex())

		var proverPrivateKey *ecdsa.PrivateKey
		if len(config.ProverPrivateKeys) != 0 {
			proverPrivateKey = config.ProverPrivateKeys[i]
		}

		vc := &keystore.VaultConfig{Address: proverAddress, ChainID: config.ChainID, PrivateKey: proverPrivateKey, Endpoint: config.Endpoint}
		proverVault, err := keystore.SetupVault(vc)
		op_common.CheckError(err, "unable to setup prover vault")

		if dcl_common.IsProverRegistered(proverAddress, config.OperatorAddress, proverRegistry) {
			fmt.Printf("prover %s is already registered\n", proverAddress.Hex())
			continue
		}

		salt := op_common.GenerateSalt()

		proverSignature := GetProverSignature(client, proverRegistry, proverAddress, proverVault, config.OperatorAddress, salt, expiry)

		regTx, err := proverRegistry.RegisterProver(transactOpts, proverAddress, salt, expiry, proverSignature)
		op_common.CheckError(err, "Registering prover-operator failed")
		fmt.Printf("Tx sent: %s\n", regTx.Hash().Hex())
		op_common.WaitForTransactionReceipt(client, regTx, config.TxReceiptTimeout)
	}
}
