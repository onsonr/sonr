package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"

	// _ "github.com/sonrhq/sonr/config"
	"github.com/sonrhq/sonr/web"
)

func ServeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "serve",
		Short:                      "Serves the Sonr Highway",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		Run:                        serveGateway,
	}
	return cmd
}

func serveGateway(_ *cobra.Command, _ []string) {
	// 1. Read config from file

	// 2. Check reachable to enabled services

	// 3. Start Gateway router as system service

	pterm.DefaultHeader.Printf(persistentHeader)
	gateway.Start()
}

const persistentHeader = `
Sonr Highway
· Gateway: http://localhost:8080
· Node RPC: http://localhost:26657
`
