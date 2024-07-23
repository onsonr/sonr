package main

import (
	"os"

	"cosmossdk.io/log"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	_ "github.com/joho/godotenv/autoload"
	"github.com/onsonr/hway/app"
)

func main() {
	rootCmd := NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		log.NewLogger(rootCmd.OutOrStderr()).Error("failure when running app", "err", err)
		os.Exit(1)
	}
}
