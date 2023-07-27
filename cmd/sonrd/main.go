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
	highlight.SetProjectID("zg0wxve9")
	highlight.Start()
	defer highlight.Stop()
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
