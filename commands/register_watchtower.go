package operator_commands

import (
	"log"

	operator_config "github.com/witnesschain-com/dcl-operator-cli/config"

	"github.com/urfave/cli/v2"
)

func RegisterWatchtowerCmd() *cli.Command {
	flags := append([]cli.Flag{}, CommonFlags...)
	var registerChallengerCmd = &cli.Command{
		Name:  "registerWatchtower",
		Usage: "Register a watchtower as challenger and prover",
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

			RegisterChallenger(config)
			RegisterProver(config)
			return nil
		},
	}
	return registerChallengerCmd
}

