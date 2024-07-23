package operator_commands

import (
	"fmt"

	dcl_common "github.com/witnesschain-com/dcl-operator-cli/common"
	"github.com/witnesschain-com/dcl-operator-cli/common/bindings/ChallengerRegistry"
	operator_config "github.com/witnesschain-com/dcl-operator-cli/dcl-operator/config"
	op_common "github.com/witnesschain-com/operator-cli/common"

	"github.com/urfave/cli/v2"
)

func DeRegisterChallengerCmd() *cli.Command {
	var deregisterChallengerCmd = &cli.Command{
		Name:  "deRegisterChallenger",
		Usage: "De-register the challenger",
		Flags: []cli.Flag{
			&op_common.ConfigPathFlag,
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.Value("config-file") == "" {
				cCtx.Set("config-file", dcl_common.DefaultOpChallengerConfig)
			}
			config := operator_config.GetConfigFromContext(cCtx)
			DeRegisterChallenger(config)
			return nil
		},
	}
	return deregisterChallengerCmd
}

func DeRegisterChallenger(config *operator_config.OperatorConfig) {
	client := op_common.ConnectToUrl(config.EthRPCUrl)

	challengerRegistry, err := ChallengerRegistry.NewChallengerRegistry(config.ChallengerRegistryAddress, client)
	op_common.CheckError(err, "Instantiating ChallengerRegistry contract failed")

	operatorPrivateKey, operatorAddress := op_common.GetECDSAPrivateAndPublicKey(op_common.GetPrivateKey(config.OperatorPrivateKey))

	if !dcl_common.IsOperatorWhitelisted(operatorAddress, challengerRegistry) {
		fmt.Printf("Operator %s is not allow listed\n", operatorAddress.Hex())
		return
	}

	deRegTransactOpts := op_common.PrepareTransactionOptions(client, config.ChainId, config.GasLimit, operatorPrivateKey)

	for _, challengerPkName := range config.ChallengerPrivateKeys {
		_, challengerAddress := op_common.GetECDSAPrivateAndPublicKey(op_common.GetPrivateKey(challengerPkName))

		if !dcl_common.IsValidChallenger(challengerAddress, challengerRegistry) {
			fmt.Printf("challenger %s is not registered\n", challengerAddress.Hex())
			continue
		}

		deRegTransactOpts.Nonce = op_common.GetLatestNonce(client, operatorPrivateKey)
		deRegTx, err := challengerRegistry.DeRegisterChallenger(deRegTransactOpts, challengerAddress)
		op_common.CheckError(err, "Deregister failed")
		fmt.Printf("Tx sent: %s\n", deRegTx.Hash().Hex())
		op_common.WaitForTransactionReceipt(client, deRegTx, config.TxReceiptTimeout)
	}
}
