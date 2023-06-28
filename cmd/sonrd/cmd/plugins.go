package cmd

import (
	"github.com/sonrhq/core/pkg/highway"
	"github.com/spf13/cobra"

)

// HighwayCmd starts the Highway Gateway server
func HighwayCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "highway",
		Short: "Starts the Highway Gateway server",
		Long: `The Gateway server is the main entry point for the Highway protocol. This network allows peers and services to
		interact with the Sonr blockchain in an authenticated and secure manner. The Gateway server is responsible for Registration, Authentication,
		Transaction Signing, Transaction Broadcasting, and Mailbox management.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return highway.Start()
		},
	}
	return cmd
}
