package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/onsonr/sonr/crypto/ucan"
	"github.com/onsonr/sonr/internal/database/sessions"
	"github.com/onsonr/sonr/internal/gateway"
	"github.com/onsonr/sonr/internal/gateway/config"
	"github.com/onsonr/sonr/pkg/common/ipfs"
	"github.com/onsonr/sonr/pkg/didauth/producer"
	"github.com/spf13/cobra"
)

// Command line flags
var (
	servePort      int    // Gateway http entry point (default 3000)
	configDir      string // Hway config directory (default hway)
	sqliteFile     string // SQLite database file (default hway.db)
	chainID        string // Current chain ID (default sonr-testnet-1)
	ipfsGatewayURL string // IPFS gateway URL (default localhost:8080)
	sonrAPIURL     string // Sonr API URL (default localhost:1317)
	sonrGrpcURL    string // Sonr gRPC URL (default localhost:9090)
	sonrRPCURL     string // Sonr RPC URL (default localhost:26657)
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
	cmd.Flags().StringVar(&configDir, "config-dir", "hway", "Directory to store config files")
	cmd.Flags().StringVar(&sqliteFile, "sqlite-file", "hway.db", "File to store sqlite database")
	cmd.Flags().StringVar(&chainID, "chain-id", "sonr-testnet-1", "Chain ID")
	cmd.Flags().StringVar(&ipfsGatewayURL, "ipfs-gateway-url", "localhost:8080", "IPFS gateway URL")
	cmd.Flags().StringVar(&sonrAPIURL, "sonr-api-url", "localhost:1317", "Sonr API URL")
	cmd.Flags().StringVar(&sonrGrpcURL, "sonr-grpc-url", "localhost:9090", "Sonr gRPC URL")
	cmd.Flags().StringVar(&sonrRPCURL, "sonr-rpc-url", "localhost:26657", "Sonr RPC URL")
	return cmd
}

func loadEnvImplFromArgs(args []string) (config.Env, error) {
	cmd := rootCmd()
	if err := cmd.ParseFlags(args); err != nil {
		return nil, err
	}

	env := &config.EnvImpl{
		ServePort:      servePort,
		ConfigDir:      configDir,
		SqliteFile:     sqliteFile,
		ChainId:        chainID,
		IpfsGatewayUrl: ipfsGatewayURL,
		SonrApiUrl:     sonrAPIURL,
		SonrGrpcUrl:    sonrGrpcURL,
		SonrRpcUrl:     sonrRPCURL,
	}
	return env, nil
}

// setupServer sets up the server
func setupServer(env config.Env) (*echo.Echo, error) {
	ipc, err := ipfs.NewClient()
	if err != nil {
		return nil, err
	}
	db, err := sessions.NewGormDB(env)
	if err != nil {
		return nil, err
	}
	e := echo.New()
	e.IPExtractor = echo.ExtractIPDirect()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(producer.Middleware(ipc, ucan.ServicePermissions))
	gateway.RegisterRoutes(e, env, db)
	return e, nil
}
