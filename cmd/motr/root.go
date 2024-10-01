package main

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "motr",
		Short: "Manage a local DWN instance for the Sonr blockchain",
	}
}
