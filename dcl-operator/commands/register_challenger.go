package operator_commands

import (
	"fmt"

	dcl_common "github.com/witnesschain-com/dcl-operator-cli/common"
	"github.com/witnesschain-com/dcl-operator-cli/common/bindings/ChallengerRegistry"
	operator_config "github.com/witnesschain-com/dcl-operator-cli/dcl-operator/config"
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
			cCtx.Set("config-file", dcl_common.DefaultOpChallengerConfig)
			config := operator_config.GetConfigFromContext(cCtx)
			RegisterChallenger(config)
			return nil
		},
	}
	return registerChallengerCmd
}

func RegisterChallenger(config *operator_config.OperatorConfig) {
	client := op_common.ConnectToUrl(config.EthRPCUrl)

	challengerRegistry, err := ChallengerRegistry.NewChallengerRegistry(config.ChallengerRegistryAddress, client)
	op_common.CheckError(err, "Instantiating ChallengerRegistry contract failed")

	operatorPrivateKey, operatorAddress := op_common.GetECDSAPrivateAndPublicKey(op_common.GetPrivateKey(config.OperatorPrivateKey))

	if !dcl_common.IsOperatorWhitelisted(operatorAddress, challengerRegistry) {
		fmt.Printf("Operator %s is not allow listed\n", operatorAddress.Hex())
		return
	}

	regTransactOpts := op_common.PrepareTransactionOptions(client, config.ChainId, config.GasLimit, operatorPrivateKey)
	expiry := op_common.CalculateExpiry(client, config.ExpiryInDays)

	for _, challengerPkName := range config.ChallengerPrivateKeys {
		challengerPrivateKey, challengerAddress := op_common.GetECDSAPrivateAndPublicKey(op_common.GetPrivateKey(challengerPkName))

		if dcl_common.IsValidChallenger(challengerAddress, challengerRegistry) {
			fmt.Printf("challenger %s is already registered\n", challengerAddress.Hex())
			continue
		}

		salt := op_common.GenerateSalt()

		challengerSignature := GetChallengerSignature(client, challengerRegistry, challengerAddress, challengerPrivateKey, operatorAddress, salt, expiry)

		regTransactOpts.Nonce = op_common.GetLatestNonce(client, operatorPrivateKey)

		regTx, err := challengerRegistry.RegisterChallenger(regTransactOpts, challengerAddress, salt, expiry, challengerSignature)
		op_common.CheckError(err, "Registering challenger failed")
		fmt.Printf("Tx sent: %s\n", regTx.Hash().Hex())
		op_common.WaitForTransactionReceipt(client, regTx, config.TxReceiptTimeout)
	}
}
