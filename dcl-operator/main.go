package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	operator_commands "github.com/witnesschain-com/dcl-operator-cli/dcl-operator/commands"
	op_common "github.com/witnesschain-com/operator-cli/common"
)

var VERSION string = "UNKNOWN"

func main() {
	cli.AppHelpTemplate = fmt.Sprintf(`
	__        ___ _                        ____ _           _       
	\ \      / (_) |_ _ __   ___  ___ ___ / ___| |__   __ _(_)_ __  
	 \ \ /\ / /| | __|  _ \ / _ \/ __/ __| |   |  _ \ / _  | |  _ \ 
	  \ V  V / | | |_| | | |  __/\__ \__ \ |___| | | | (_| | | | | |
	   \_/\_/  |_|\__|_| |_|\___||___/___/\____|_| |_|\__,_|_|_| |_|

	%s`, cli.AppHelpTemplate)
	app := cli.NewApp()

	app.Name = "WitnessChain"
	app.Usage = "WitnessChain DCL Operator CLI"
	app.Version = VERSION
	app.Copyright = "(c) 2024 WitnessChain"

	app.Commands = []*cli.Command{
		op_common.KeysCmd(),
		operator_commands.RegisterProverCmd(),
		operator_commands.DeRegisterProverCmd(),
		operator_commands.RegisterChallengerCmd(),
		operator_commands.DeRegisterChallengerCmd(),
	}

	if err := app.Run(os.Args); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		op_common.Unmount()
		os.Exit(1)
	}

	op_common.Unmount()
}
