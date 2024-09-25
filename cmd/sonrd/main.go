package main

import (
	"os"

	"cosmossdk.io/log"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	_ "github.com/joho/godotenv/autoload"

	"github.com/onsonr/sonr/app"
	"github.com/onsonr/sonr/app/cli"
	"github.com/onsonr/sonr/app/proxy"
)

func main() {
	rootCmd := NewRootCmd()
	rootCmd.AddCommand(cli.NewBuildTxnTUICmd())
	rootCmd.AddCommand(cli.NewExplorerTUICmd())
	rootCmd.AddCommand(proxy.NewProxyCmd())

	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		log.NewLogger(rootCmd.OutOrStderr()).Error("failure when running app", "err", err)
		os.Exit(1)
	}
}
