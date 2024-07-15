package operator_commands

import (
	"fmt"

	dcl_common "github.com/witnesschain-com/dcl-operator-cli/common"
	"github.com/witnesschain-com/dcl-operator-cli/common/bindings/ProverRegistry"
	operator_config "github.com/witnesschain-com/dcl-operator-cli/dcl-operator/config"
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
			cCtx.Set("config-file", dcl_common.DefaultOpProverConfig)
			config := operator_config.GetConfigFromContext(cCtx)
			RegisterProver(config)
			return nil
		},
	}
	return registerProverCmd
}

func RegisterProver(config *operator_config.OperatorConfig) {
	client := op_common.ConnectToUrl(config.EthRPCUrl)

	proverRegistry, err := ProverRegistry.NewProverRegistry(config.ProverRegistryAddress, client)
	op_common.CheckError(err, "Instantiating ProverRegistry contract failed")

	operatorPrivateKey, operatorAddress := op_common.GetECDSAPrivateAndPublicKey(op_common.GetPrivateKey(config.OperatorPrivateKey))

	if !dcl_common.IsOperatorWhitelisted(operatorAddress, proverRegistry) {
		fmt.Printf("Operator %s is not whitelisted\n", operatorAddress.Hex())
		return
	}

	regTransactOpts := op_common.PrepareTransactionOptions(client, config.ChainId, config.GasLimit, operatorPrivateKey)
	expiry := op_common.CalculateExpiry(client, config.ExpiryInDays)

	for _, proverPkName := range config.ProverPrivateKeys {
		proverPrivateKey, proverAddress := op_common.GetECDSAPrivateAndPublicKey(op_common.GetPrivateKey(proverPkName))

		if dcl_common.IsProverRegistered(proverAddress, operatorAddress, proverRegistry) {
			fmt.Printf("prover %s is already registered\n", proverAddress.Hex())
			continue
		}

		salt := op_common.GenerateSalt()

		proverSignature := GetProverSignature(client, proverRegistry, proverAddress, proverPrivateKey, operatorAddress, salt, expiry)

		regTransactOpts.Nonce = op_common.GetLatestNonce(client, operatorPrivateKey)

		regTx, err := proverRegistry.RegisterProver(regTransactOpts, proverAddress, salt, expiry, proverSignature)
		op_common.CheckError(err, "Registering prover-operator failed")
		fmt.Printf("Tx sent: %s\n", regTx.Hash().Hex())
		op_common.WaitForTransactionReceipt(client, regTx, config.TxReceiptTimeout)
	}
}
