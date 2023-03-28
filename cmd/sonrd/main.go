package main

import (
	"log"
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/getsentry/sentry-go"
	"github.com/sonrhq/core/app"
	"github.com/sonrhq/core/cmd/sonrd/cmd"
	"github.com/sonrhq/core/internal/local"
)

func main() {
	snrctx := local.NewContext()
	if snrctx.IsProd() {
		err := sentry.Init(sentry.ClientOptions{
			Dsn:              "https://df1e6b7c6e944de0acbab4f7a862a3d9@o4504358155911168.ingest.sentry.io/4504358160039936",
			TracesSampleRate: 1.0,
		})
		if err != nil {
			log.Fatalf("sentry.Init: %s", err)
		}
	}
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
