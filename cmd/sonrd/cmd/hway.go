package cmd

import (
	"github.com/spf13/cobra"

	// _ "github.com/sonrhq/sonr/config"
	"github.com/sonrhq/sonr/pkg/highway"
)

// ServeCommand returns the serve command
func ServeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "hway",
		Short:                      "Serves the Sonr Highway",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		Run:                        serveAction,
	}
	return cmd
}

func serveAction(_ *cobra.Command, _ []string) {
	// 1. Read config from file
	// 2. Check reachable to enabled services
	// 3. Start Gateway router as system service
	highway.Start()
}
