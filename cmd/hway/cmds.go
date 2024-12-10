package main

import (
	"github.com/onsonr/sonr/internal/gateway/config"
	"github.com/spf13/cobra"
)

var (
	FlagServePort      int
	FlagConfigDir      string
	FlagSqliteFile     string
	FlagChainId        string
	FlagIpfsGatewayUrl string
	FlagSonrApiUrl     string
	FlagSonrGrpcUrl    string
	FlagSonrRpcUrl     string
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hway",
		Short: "Sonr DID gateway",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.Flags().IntVar(&servePort, "serve-port", FlagServePort, "Port to serve the gateway on")
	cmd.Flags().StringVar(&configDir, "config-dir", FlagConfigDir, "Directory to store config files")
	cmd.Flags().StringVar(&sqliteFile, "sqlite-file", FlagSqliteFile, "File to store sqlite database")
	cmd.Flags().StringVar(&chainId, "chain-id", FlagChainId, "Chain ID")
	cmd.Flags().StringVar(&ipfsGatewayUrl, "ipfs-gateway-url", FlagIpfsGatewayUrl, "IPFS gateway URL")
	cmd.Flags().StringVar(&sonrApiUrl, "sonr-api-url", FlagSonrApiUrl, "Sonr API URL")
	cmd.Flags().StringVar(&sonrGrpcUrl, "sonr-grpc-url", FlagSonrGrpcUrl, "Sonr gRPC URL")
	return cmd
}

func loadEnvImplFromArgs(args []string) (config.Env, error) {
	var servePort int
	var configDir string
	var sqliteFile string
	var chainId string
	var ipfsGatewayUrl string
	var sonrApiUrl string
	var sonrGrpcUrl string
	var sonrRpcUrl string

	cmd := NewRootCmd()

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
