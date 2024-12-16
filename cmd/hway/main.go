package main

import (
	_ "embed"
	"fmt"
	"os"

	config "github.com/onsonr/sonr/internal/config/hway"
)

// main is the entry point for the application
func main() {
	cmd := rootCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func loadEnvImplFromArgs(args []string) (config.Hway, error) {
	cmd := rootCmd()
	if err := cmd.ParseFlags(args); err != nil {
		return nil, err
	}

	env := &config.HwayImpl{
		ServePort:      servePort,
		ChainId:        chainID,
		IpfsGatewayUrl: ipfsGatewayURL,
		SonrApiUrl:     sonrAPIURL,
		SonrGrpcUrl:    sonrGrpcURL,
		SonrRpcUrl:     sonrRPCURL,
		PsqlDSN:        formatPsqlDSN(),
	}
	return env, nil
}
