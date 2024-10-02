package commands

import (
	"github.com/spf13/cobra"

	"github.com/onsonr/sonr/cmd/hway/server"
)

func NewStartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Starts the DWN proxy server for the local IPFS node",
		Run: func(cmd *cobra.Command, args []string) {
			s := server.New()
			s.Start()
		},
	}
}
