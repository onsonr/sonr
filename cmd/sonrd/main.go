package main

import (
	"os"

	"cosmossdk.io/log"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	_ "github.com/joho/godotenv/autoload"

	"github.com/onsonr/sonr/app"
	"github.com/onsonr/sonr/internal/tui"
)

func main() {
	rootCmd := NewRootCmd()
	rootCmd.AddCommand(tui.NewBuildProtoMsgCmd())
	// rootCmd.AddCommand(tui.NewExplorerCmd())

	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		log.NewLogger(rootCmd.OutOrStderr()).Error("failure when running app", "err", err)
		os.Exit(1)
	}
}
