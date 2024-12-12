package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// Command line flags
var (
	servePort      int    // Gateway http entry point (default 3000)
	chainID        string // Current chain ID (default sonr-testnet-1)
	ipfsGatewayURL string // IPFS gateway URL (default localhost:8080)
	sonrAPIURL     string // Sonr API URL (default localhost:1317)
	sonrGrpcURL    string // Sonr gRPC URL (default localhost:9090)
	sonrRPCURL     string // Sonr RPC URL (default localhost:26657)

	sqliteFile string // SQLite database file (default hway.db)
	psqlHost   string // PostgresSQL Host Flag
	psqlUser   string // PostgresSQL User Flag
	psqlPass   string // PostgresSQL Password Flag
	psqlDB     string // PostgresSQL Database Flag
)

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hway",
		Short: "Sonr DID gateway",
		Run: func(cmd *cobra.Command, args []string) {
			env, err := loadEnvImplFromArgs(args)
			if err != nil {
				panic(err)
			}
			e, err := setupServer(env)
			if err != nil {
				panic(err)
			}
			if err := e.Start(fmt.Sprintf(":%d", env.GetServePort())); err != http.ErrServerClosed {
				log.Fatal(err)
				os.Exit(1)
				return
			}
		},
	}
	cmd.Flags().IntVar(&servePort, "serve-port", 3000, "Port to serve the gateway on")
	cmd.Flags().StringVar(&chainID, "chain-id", "sonr-testnet-1", "Chain ID")
	cmd.Flags().StringVar(&ipfsGatewayURL, "ipfs-gateway-url", "localhost:8080", "IPFS gateway URL")
	cmd.Flags().StringVar(&sonrAPIURL, "sonr-api-url", "localhost:1317", "Sonr API URL")
	cmd.Flags().StringVar(&sonrGrpcURL, "sonr-grpc-url", "localhost:9090", "Sonr gRPC URL")
	cmd.Flags().StringVar(&sonrRPCURL, "sonr-rpc-url", "localhost:26657", "Sonr RPC URL")
	cmd.Flags().StringVar(&sqliteFile, "sqlite-file", "hway.db", "File to store sqlite database")
	cmd.Flags().StringVar(&psqlHost, "psql-host", "localhost", "PostgresSQL Host")
	cmd.Flags().StringVar(&psqlUser, "psql-user", "postgres", "PostgresSQL User")
	cmd.Flags().StringVar(&psqlPass, "psql-pass", "postgres", "PostgresSQL Password")
	cmd.Flags().StringVar(&psqlDB, "psql-db", "hway", "PostgresSQL Database")
	return cmd
}

func formatPsqlDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", psqlHost, psqlUser, psqlPass, psqlDB)
}
