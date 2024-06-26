package operator_commands

import (
	"fmt"

	dcl_common "github.com/witnesschain-com/dcl-operator-cli/common"
	"github.com/witnesschain-com/dcl-operator-cli/common/bindings/ProverRegistry"
	operator_config "github.com/witnesschain-com/dcl-operator-cli/dcl-operator/config"
	op_common "github.com/witnesschain-com/operator-cli/common"

	"github.com/urfave/cli/v2"
)

func DeRegisterProverCmd() *cli.Command {
	var deregisterProverCmd = &cli.Command{
		Name:  "deRegisterProver",
		Usage: "De-register the prover",
		Flags: []cli.Flag{
			&op_common.ConfigPathFlag,
		},
		Action: func(cCtx *cli.Context) error {
			cCtx.Set("config-file", dcl_common.DefaultOpProverConfig)
			config := operator_config.GetConfigFromContext(cCtx)
			DeRegisterProver(config)
			return nil
		},
	}
	return deregisterProverCmd
}

func DeRegisterProver(config *operator_config.OperatorConfig) {
	client := op_common.ConnectToUrl(config.EthRPCUrl)

	proverRegistry, err := ProverRegistry.NewProverRegistry(config.ProverRegistryAddress, client)
	op_common.CheckError(err, "Instantiating ProverRegistry contract failed")

	operatorPrivateKey, operatorAddress := op_common.GetECDSAPrivateAndPublicKey(op_common.GetPrivateKey(config.OperatorPrivateKey))

	if !dcl_common.IsOperatorWhitelisted(operatorAddress, proverRegistry) {
		fmt.Printf("Operator %s is not whitelisted\n", operatorAddress.Hex())
		return
	}

	deRegTransactOpts := op_common.PrepareTransactionOptions(client, config.ChainId, config.GasLimit, operatorPrivateKey)

	for _, proverPkName := range config.ProverPrivateKeys {
		_, proverAddress := op_common.GetECDSAPrivateAndPublicKey(op_common.GetPrivateKey(proverPkName))

		if !dcl_common.IsProverRegistered(proverAddress, operatorAddress, proverRegistry) {
			fmt.Printf("prover %s is not registered\n", proverAddress.Hex())
			continue
		}

		deRegTransactOpts.Nonce = op_common.GetLatestNonce(client, operatorPrivateKey)
		deRegTx, err := proverRegistry.DeRegisterProver(deRegTransactOpts, proverAddress)
		op_common.CheckError(err, "Deregister failed")
		fmt.Printf("Tx sent: %s\n", deRegTx.Hash().Hex())
		op_common.WaitForTransactionReceipt(client, deRegTx, config.TxReceiptTimeout)
	}
}
