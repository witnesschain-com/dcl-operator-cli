package operator_commands

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	dcl_common "github.com/witnesschain-com/dcl-operator-cli/common"
	"github.com/witnesschain-com/dcl-operator-cli/common/bindings/ChallengerRegistry"
	operator_config "github.com/witnesschain-com/dcl-operator-cli/dcl-operator/config"
	"github.com/witnesschain-com/diligencewatchtower-client/keystore"
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
			config := operator_config.GetChallengerConfigFromContext(cCtx)
			DeRegisterChallenger(config)
			return nil
		},
	}
	return deregisterChallengerCmd
}

func DeRegisterChallenger(config *operator_config.OperatorConfig) {
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

	for _, challengerAddress := range config.ChallengerAddresses {

		fmt.Println("challengerAddress: " + challengerAddress.Hex())

		if !dcl_common.IsValidChallenger(challengerAddress, challengerRegistry) {
			fmt.Printf("challenger %s is not registered\n", challengerAddress.Hex())
			continue
		}

		deRegTx, err := challengerRegistry.DeRegisterChallenger(transactOpts, challengerAddress)
		op_common.CheckError(err, "Deregistering challenger failed")
		fmt.Printf("Tx sent: %s\n", deRegTx.Hash().Hex())
		op_common.WaitForTransactionReceipt(client, deRegTx, config.TxReceiptTimeout)
	}
}
