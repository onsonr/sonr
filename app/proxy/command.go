package proxy

import (
	"github.com/spf13/cobra"
)

func NewProxyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "dwn-proxy",
		Short: "Starts the DWN proxy server for the local IPFS node",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Load config
			// c, err := LoadConfig(".")
			// if err != nil {
			// 	return err
			// }
			// log.Printf("Config: %+v", c)
			startServer()
			return nil
		},
	}
}
