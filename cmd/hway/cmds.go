package main

import (
	"github.com/onsonr/sonr/internal/gateway/config"
	"github.com/spf13/cobra"
)

// Command line flags
var (
	servePort      int
	configDir      string
	sqliteFile     string
	chainId        string
	ipfsGatewayUrl string
	sonrApiUrl     string
	sonrGrpcUrl    string
	sonrRpcUrl     string
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hway",
		Short: "Sonr DID gateway",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.Flags().IntVar(&servePort, "serve-port", 8080, "Port to serve the gateway on")
	cmd.Flags().StringVar(&configDir, "config-dir", "", "Directory to store config files")
	cmd.Flags().StringVar(&sqliteFile, "sqlite-file", "", "File to store sqlite database")
	cmd.Flags().StringVar(&chainId, "chain-id", "", "Chain ID")
	cmd.Flags().StringVar(&ipfsGatewayUrl, "ipfs-gateway-url", "", "IPFS gateway URL")
	cmd.Flags().StringVar(&sonrApiUrl, "sonr-api-url", "", "Sonr API URL")
	cmd.Flags().StringVar(&sonrGrpcUrl, "sonr-grpc-url", "", "Sonr gRPC URL")
	cmd.Flags().StringVar(&sonrRpcUrl, "sonr-rpc-url", "", "Sonr RPC URL")
	return cmd
}

func loadEnvImplFromArgs(args []string) (config.Env, error) {
	cmd := NewRootCmd()
	if err := cmd.ParseFlags(args); err != nil {
		return nil, err
	}

	env := &config.EnvImpl{
		ServePort:      servePort,
		ConfigDir:      configDir,
		SqliteFile:     sqliteFile,
		ChainId:        chainId,
		IpfsGatewayUrl: ipfsGatewayUrl,
		SonrApiUrl:     sonrApiUrl,
		SonrGrpcUrl:    sonrGrpcUrl,
		SonrRpcUrl:     sonrRpcUrl,
	}
	return env, nil
}
