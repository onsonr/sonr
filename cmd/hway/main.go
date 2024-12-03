package main

import (
	_ "embed"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/onsonr/sonr/pkg/common/clients"
	"github.com/onsonr/sonr/pkg/common/session"
	"github.com/onsonr/sonr/pkg/gateway"
	// TODO: Integrate TigerBeetle
	// _ "github.com/tigerbeetle/tigerbeetle-go"
	// _ "github.com/tigerbeetle/tigerbeetle-go/pkg/types"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.HwayMiddleware())
	e.Use(clients.IPFSMiddleware())
	e.Use(clients.GRPCClientsMiddleware("localhost:9090"))
	gateway.RegisterRoutes(e)

	if err := e.Start(":3000"); err != http.ErrServerClosed {
		log.Fatal(err)
		os.Exit(1)
		return
	}
}
