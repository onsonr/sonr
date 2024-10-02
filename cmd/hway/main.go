package main

import (
	"github.com/onsonr/sonr/cmd/hway/commands"
)

func main() {
	rootCmd := commands.NewRootCmd()
	rootCmd.AddCommand(commands.NewStartCmd())
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
