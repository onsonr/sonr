package cmd

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"

	// _ "github.com/sonrhq/sonr/config"
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

func serveGateway(_ *cobra.Command, _ []string) {
	pterm.DefaultBigText.WithLetters(putils.LettersFromStringWithStyle("SONR", pterm.FgCyan.ToStyle())).Render()
	// 1. Read config from file

	// 2. Check reachable to enabled services

	// 3. Start Gateway router as system service
	gateway.Start()
}
