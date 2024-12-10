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
	cmd.Flags().StringVar(&chainId, "chain-id", "sonr-testnet-1", "Chain ID")
	cmd.Flags().StringVar(&ipfsGatewayUrl, "ipfs-gateway-url", "localhost:8080", "IPFS gateway URL")
	cmd.Flags().StringVar(&sonrApiUrl, "sonr-api-url", "localhost:1317", "Sonr API URL")
	cmd.Flags().StringVar(&sonrGrpcUrl, "sonr-grpc-url", "localhost:9090", "Sonr gRPC URL")
	cmd.Flags().StringVar(&sonrRpcUrl, "sonr-rpc-url", "localhost:26657", "Sonr RPC URL")
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
