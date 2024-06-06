package operator_commands

import (
	"fmt"

	wc_common "github.com/witnesschain-com/dcl-operator-cli/common"
	"github.com/witnesschain-com/dcl-operator-cli/common/bindings/ProverRegistry"
	operator_config "github.com/witnesschain-com/dcl-operator-cli/dcl-operator/config"

	"github.com/urfave/cli/v2"
)

func DeRegisterProverCmd() *cli.Command {
	wc_common.ConfigPathFlag.Value = wc_common.DefaultOpConfig
	var deregisterProverCmd = &cli.Command{
		Name:  "deRegisterProver",
		Usage: "De-register the prover",
		Flags: []cli.Flag{
			&wc_common.ConfigPathFlag,
		},
		Action: func(cCtx *cli.Context) error {
			config := operator_config.GetConfigFromContext(cCtx)
			DeRegisterProver(config)
			return nil
		},
	}
	return deregisterProverCmd
}

func DeRegisterProver(config *operator_config.OperatorConfig) {
	client := wc_common.ConnectToUrl(config.EthRPCUrl)

	proverRegistry, err := ProverRegistry.NewProverRegistry(config.ProverRegistryAddress, client)
	wc_common.CheckError(err, "Instantiating ProverRegistry contract failed")

	operatorPrivateKey, operatorAddress := wc_common.GetECDSAPrivateAndPublicKey(wc_common.GetPrivateKey(config.OperatorPrivateKey))

	if !wc_common.IsOperatorWhitelisted(operatorAddress, proverRegistry) {
		fmt.Printf("Operator %s is not whitelisted\n", operatorAddress.Hex())
		return
	}

	deRegTransactOpts := wc_common.PrepareTransactionOptions(client, config.ChainId, config.GasLimit, operatorPrivateKey)

	for _, proverPkName := range config.ProverPrivateKeys {
		_, proverAddress := wc_common.GetECDSAPrivateAndPublicKey(wc_common.GetPrivateKey(proverPkName))

		if !wc_common.IsProverRegistered(proverAddress, operatorAddress, proverRegistry) {
			fmt.Printf("prover %s is not registered\n", proverAddress.Hex())
			continue
		}

		deRegTransactOpts.Nonce = wc_common.GetLatestNonce(client, operatorPrivateKey)
		deRegTx, err := proverRegistry.DeRegisterProver(deRegTransactOpts, proverAddress)
		wc_common.CheckError(err, "Deregister failed")
		fmt.Printf("Tx sent: %s\n", deRegTx.Hash().Hex())
		wc_common.WaitForTransactionReceipt(client, deRegTx, config.TxReceiptTimeout)
	}
}
