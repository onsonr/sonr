package main

import "github.com/spf13/cobra"

func main() {
	rootCmd := &cobra.Command{
		Use:   "motr",
		Short: "Manage a local DWN instance for the Sonr blockchain",
	}
	rootCmd.AddCommand(NewProxyCmd())
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
