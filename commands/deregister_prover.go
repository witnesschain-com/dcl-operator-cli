package operator_commands

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	dcl_common "github.com/witnesschain-com/dcl-operator-cli/common"
	"github.com/witnesschain-com/dcl-operator-cli/common/bindings/ProverRegistry"
	operator_config "github.com/witnesschain-com/dcl-operator-cli/config"
	"github.com/witnesschain-com/diligencewatchtower-client/keystore"
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
			if cCtx.Value("config-file") == "" {
				cCtx.Set("config-file", dcl_common.DefaultOpProverConfig)
			}
			config := operator_config.GetProverConfigFromContext(cCtx)
			DeRegisterProver(config)
			return nil
		},
	}
	return deregisterProverCmd
}

func DeRegisterProver(config *operator_config.OperatorConfig) {

	var client *ethclient.Client
	client, config.ChainID = op_common.ConnectToUrl(config.EthRPCUrl)

	proverRegistry, err := ProverRegistry.NewProverRegistry(dcl_common.NetworkConfig[config.ChainID.String()].ProverRegistryAddress, client)
	op_common.CheckError(err, "Instantiating ProverRegistry contract failed")

	vc := &keystore.VaultConfig{Address: config.OperatorAddress, ChainID: config.ChainID, PrivateKey: config.OperatorPrivateKey, Endpoint: config.Endpoint}
	operatorVault, err := keystore.SetupVault(vc)
	if err != nil {
		op_common.CheckError(err, "unable to setup vault")
	}

	transactOpts := operatorVault.NewTransactOpts(config.ChainID)
	
	if (dcl_common.NetworkConfig[config.ChainID.String()].GasPrice == -1) {
		transactOpts.GasPrice = big.NewInt(0)
	}

	for _, proverAddress := range config.ProverAddresses {

		fmt.Println("proverAddress: " + proverAddress.Hex())

		if !dcl_common.IsProverRegistered(proverAddress, config.OperatorAddress, proverRegistry) {
			fmt.Printf("prover %s is not registered\n", proverAddress.Hex())
			continue
		}

		deRegTx, err := proverRegistry.DeRegisterProver(transactOpts, proverAddress)
		op_common.CheckError(err, "Registering prover-operator failed")
		fmt.Printf("Tx sent: %s/tx/%s\n", dcl_common.NetworkConfig[config.ChainID.String()].BlockExplorer, deRegTx.Hash().Hex())
		op_common.WaitForTransactionReceipt(client, deRegTx, config.TxReceiptTimeout)
	}
}
