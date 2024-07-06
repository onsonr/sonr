package main

import (
	"github.com/onsonr/hway/pkg/vault"
	"github.com/spf13/cobra"
)

func main() {
	if err := serveCmd().Execute(); err != nil {
		panic(err)
	}
}

func serveCmd() *cobra.Command {
	return &cobra.Command{
		Use:                "vltd",
		Aliases:            []string{"vault"},
		Short:              "run the vault rest api and htmx frontend",
		DisableFlagParsing: false,
		Run: func(cmd *cobra.Command, args []string) {
			vault.Serve(cmd.Context())
		},
	}
}
