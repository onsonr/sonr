package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/onsonr/sonr/crypto/ucan"
	"github.com/onsonr/sonr/internal/gateway"
	"github.com/onsonr/sonr/pkg/common/ipfs"
	config "github.com/onsonr/sonr/pkg/config/hway"
	"github.com/onsonr/sonr/pkg/database/sessions"
	"github.com/onsonr/sonr/pkg/didauth/producer"
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
func setupServer(env config.Hway) (*echo.Echo, error) {
	ipc, err := ipfs.NewClient()
	if err != nil {
		return nil, err
	}
	db, err := sessions.NewGormDB(env)
	if err != nil {
		return nil, err
	}
	e := echo.New()
	e.Use(echoprometheus.NewMiddleware("hway"))
	e.IPExtractor = echo.ExtractIPDirect()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(producer.Middleware(ipc, ucan.ServicePermissions))
	gateway.RegisterRoutes(e, env, db)
	return e, nil
}
