package cmd

import (
	"github.com/spf13/cobra"
)

func ServeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "serve",
		Short:                      "Gateway subcommands",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
	}
	return cmd
}
