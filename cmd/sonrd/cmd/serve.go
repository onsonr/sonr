package cmd

import (
	"github.com/spf13/cobra"

	"github.com/sonrhq/sonr/gateway"
)

func ServeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "serve",
		Short:                      "Serves the Sonr Highway Gateway",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		Run:                        serveGateway,
	}
	return cmd
}

func serveGateway(cmd *cobra.Command, args []string) {
	// 1. Read config from file

	// 2. Check reachable to enabled services

	// 3. Start Gateway router as system service
	gateway.Start()
}
