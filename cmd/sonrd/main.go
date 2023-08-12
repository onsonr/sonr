package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/highlight/highlight/sdk/highlight-go"

	"github.com/sonrhq/core/app"
	"github.com/sonrhq/core/cmd/sonrd/cmd"
)

func main() {

	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "SONR", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			highlight.Stop()
			os.Exit(e.Code)

		default:
			highlight.Stop()
			os.Exit(1)
		}
	}
}
