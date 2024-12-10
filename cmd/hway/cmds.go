package main

import (
	"github.com/onsonr/sonr/internal/gateway/config"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "hway",
		Short: "Sonr DID gateway",

		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
}

func getEnvImpl(cmd *cobra.Command) (config.Env, error) {
	servePort, _ := cmd.Flags().GetInt("serve-port")
	configDir, _ := cmd.Flags().GetString("config-dir")
	sqliteFile, _ := cmd.Flags().GetString("sqlite-file")
	chainId, _ := cmd.Flags().GetString("chain-id")
	ipfsGatewayUrl, _ := cmd.Flags().GetString("ipfs-gateway-url")
	sonrApiUrl, _ := cmd.Flags().GetString("sonr-api-url")
	sonrGrpcUrl, _ := cmd.Flags().GetString("sonr-grpc-url")
	sonrRpcUrl, _ := cmd.Flags().GetString("sonr-rpc-url")
	// Load from flags
	env := config.EnvImpl{
		ServePort:      servePort,
		ConfigDir:      configDir,
		SqliteFile:     sqliteFile,
		ChainId:        chainId,
		IpfsGatewayUrl: ipfsGatewayUrl,
		SonrApiUrl:     sonrApiUrl,
		SonrGrpcUrl:    sonrGrpcUrl,
		SonrRpcUrl:     sonrRpcUrl,
	}
	return &env, nil
}
