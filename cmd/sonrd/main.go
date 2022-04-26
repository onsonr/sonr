package main

import (
	"fmt"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/sonr-io/blockchain/app"
	cmd "github.com/tendermint/spm/cosmoscmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.New,
		cmd.WithEnvPrefix("SONR_CHAIN"),
		// this line is used by starport scaffolding # root/arguments

	)
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Execute the root command.
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
