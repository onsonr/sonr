package main

import (
	_ "embed"
	"fmt"
	"os"

	config "github.com/onsonr/sonr/pkg/config/hway"
	"github.com/onsonr/sonr/pkg/ipfsapi"
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

func initDeps(env config.Hway) (ipfsapi.Client, error) {
	ipc, err := ipfsapi.NewClient()
	if err != nil {
		return nil, err
	}

	return ipc, nil
}

func loadEnvImplFromArgs(args []string) (config.Hway, error) {
	cmd := rootCmd()
	if err := cmd.ParseFlags(args); err != nil {
		return nil, err
	}

	env := &config.HwayImpl{
		ServePort:      servePort,
		SqliteFile:     sqliteFile,
		ChainId:        chainID,
		IpfsGatewayUrl: ipfsGatewayURL,
		SonrApiUrl:     sonrAPIURL,
		SonrGrpcUrl:    sonrGrpcURL,
		SonrRpcUrl:     sonrRPCURL,
		PsqlDSN:        formatPsqlDSN(),
	}
	return env, nil
}
