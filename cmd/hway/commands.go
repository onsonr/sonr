package main

import "github.com/spf13/cobra"

func serveLandingCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "landing",
		Short: "Serve the landing page",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}

func serveGatewayCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "gateway",
		Short: "Serve the Vault Gateway",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
