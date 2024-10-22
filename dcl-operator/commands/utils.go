package operator_commands

import (
	"github.com/urfave/cli/v2"
)

var CommonFlags = []cli.Flag{
	&cli.BoolFlag{
		Name: "testnet",
		Usage: "BlueOrangutan testnet",
	},
	&cli.BoolFlag{
		Name: "mainnet",
		Usage: "Witness Chain Mainnet",
	},
	&cli.StringFlag{
		Name: "watchtower-private-key",
		Usage: "Private key of watchtower. e.g. 01234567890abcde01234567890abcde01234567890abcde01234567890abcde",
	},
}
