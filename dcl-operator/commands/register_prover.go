package operator_commands

import (
	"fmt"

	wc_common "github.com/witnesschain-com/dcl-operator-cli/common"
	"github.com/witnesschain-com/dcl-operator-cli/common/bindings/ProverRegistry"
	operator_config "github.com/witnesschain-com/dcl-operator-cli/dcl-operator/config"

	"github.com/urfave/cli/v2"
)

func RegisterProverCmd() *cli.Command {
	wc_common.ConfigPathFlag.Value = wc_common.DefaultOpConfig
	var registerProverCmd = &cli.Command{
		Name:  "registerProver",
		Usage: "Register a prover",
		Flags: []cli.Flag{
			&wc_common.ConfigPathFlag,
		},
		Action: func(cCtx *cli.Context) error {
			config := operator_config.GetConfigFromContext(cCtx)
			RegisterProver(config)
			return nil
		},
	}
	return registerProverCmd
}

func RegisterProver(config *operator_config.OperatorConfig) {
	client := wc_common.ConnectToUrl(config.EthRPCUrl)

	proverRegistry, err := ProverRegistry.NewProverRegistry(config.ProverRegistryAddress, client)
	wc_common.CheckError(err, "Instantiating ProverRegistry contract failed")

	operatorPrivateKey, operatorAddress := wc_common.GetECDSAPrivateAndPublicKey(wc_common.GetPrivateKey(config.OperatorPrivateKey))

	if !wc_common.IsOperatorWhitelisted(operatorAddress, proverRegistry) {
		fmt.Printf("Operator %s is not whitelisted\n", operatorAddress.Hex())
		return
	}

	regTransactOpts := wc_common.PrepareTransactionOptions(client, config.ChainId, config.GasLimit, operatorPrivateKey)
	expiry := wc_common.CalculateExpiry(client, config.ExpiryInDays)

	for _, proverPkName := range config.ProverPrivateKeys {
		proverPrivateKey, proverAddress := wc_common.GetECDSAPrivateAndPublicKey(wc_common.GetPrivateKey(proverPkName))

		if wc_common.IsProverRegistered(proverAddress, operatorAddress, proverRegistry) {
			fmt.Printf("prover %s is already registered\n", proverAddress.Hex())
			continue
		}

		salt := wc_common.GenerateSalt()

		proverSignature := GetProverSignature(client, proverRegistry, proverAddress, proverPrivateKey, operatorAddress, salt, expiry)

		regTransactOpts.Nonce = wc_common.GetLatestNonce(client, operatorPrivateKey)

		regTx, err := proverRegistry.RegisterProver(regTransactOpts, proverAddress, salt, expiry, proverSignature)
		wc_common.CheckError(err, "Registering prover-operator failed")
		fmt.Printf("Tx sent: %s\n", regTx.Hash().Hex())
		wc_common.WaitForTransactionReceipt(client, regTx, config.TxReceiptTimeout)
	}
}
