package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/onsonr/sonr/crypto/ucan"
	"github.com/onsonr/sonr/internal/gateway"
	"github.com/onsonr/sonr/internal/gateway/config"
	"github.com/onsonr/sonr/pkg/common/ipfs"
	"github.com/onsonr/sonr/pkg/didauth/producer"
)

//go:embed config.pkl
var configBz []byte

func loadConfig() (config.Env, error) {
	return config.LoadFromBytes(configBz)
}

// setupServer sets up the server
func setupServer(env config.Env) (*echo.Echo, error) {
	ipc, err := ipfs.NewClient()
	if err != nil {
		return nil, err
	}
	e := echo.New()
	e.IPExtractor = echo.ExtractIPDirect()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(producer.Middleware(ipc, ucan.ServicePermissions))
	gateway.RegisterRoutes(e, env)
	return e, nil
}

// main is the entry point for the application
func main() {
	env, err := loadConfig()
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
}
