package operator_commands

import (
	"log"

	operator_config "github.com/witnesschain-com/dcl-operator-cli/config"

	"github.com/urfave/cli/v2"
)

func DeRegisterWatchtowerCmd() *cli.Command {
	flags := append([]cli.Flag{}, CommonFlags...)
	var deRegisterWatchtowerCmd = &cli.Command{
		Name:  "deRegisterWatchtower",
		Usage: "DeRegister a watchtower as challenger and prover",
		Flags: flags,
		Action: func(cCtx *cli.Context) error {
			if !(cCtx.Bool("testnet") || cCtx.Bool("mainnet")){
				log.Fatalf("either --testnet or --mainnet is required")
			}
			
			if (cCtx.String("watchtower-private-key") == ""){
				log.Fatalf("--watchtower-private-key required");
			}

			var config *operator_config.OperatorConfig

			if (cCtx.Bool("testnet") || cCtx.Bool("mainnet") ) {
				config = operator_config.GetConfigFromContext(cCtx)
			} 

			DeRegisterChallenger(config)
			DeRegisterProver(config)
			return nil
		},
	}
	return deRegisterWatchtowerCmd
}

