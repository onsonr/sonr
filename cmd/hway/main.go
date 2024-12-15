package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	"os"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/onsonr/sonr/internal/gateway"
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

func initDeps(env config.Hway) (*sql.DB, ipfsapi.Client, error) {
	db, err := gateway.NewDB(env)
	if err != nil {
		return nil, nil, err
	}

	ipc, err := ipfsapi.NewClient()
	if err != nil {
		return nil, nil, err
	}

	return db, ipc, nil
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

// setupServer sets up the server
func setupServer(env config.Hway, db *sql.DB, ipc ipfsapi.Client) (*echo.Echo, error) {
	e := echo.New()
	e.Use(echoprometheus.NewMiddleware("hway"))
	e.IPExtractor = echo.ExtractIPDirect()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	gateway.RegisterRoutes(e, env, db, ipc)
	return e, nil
}
