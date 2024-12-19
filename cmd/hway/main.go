package main

import (
	"context"
	_ "embed"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	config "github.com/onsonr/sonr/internal/config/hway"
	hwayorm "github.com/onsonr/sonr/pkg/gateway/orm"
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

func setupPostgresDB() (*hwayorm.Queries, error) {
	pgdsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", psqlHost, psqlUser, psqlPass, psqlDB)
	conn, err := pgx.Connect(context.Background(), pgdsn)
	if err != nil {
		return nil, err
	}
	return hwayorm.New(conn), nil
}
