package main

import (
	_ "embed"
	"fmt"
	"net/http"

	"github.com/ipfs/kubo/client/rpc"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/onsonr/sonr/pkg/common/session"
	"github.com/onsonr/sonr/pkg/gateway"
	"github.com/onsonr/sonr/pkg/webui/landing"

	gatewaymiddleware "github.com/onsonr/sonr/pkg/gateway/middleware"
	// TODO: Integrate TigerBeetle
	// _ "github.com/tigerbeetle/tigerbeetle-go"
	// _ "github.com/tigerbeetle/tigerbeetle-go/pkg/types"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func main() {
	// Setup Echo
	hosts := map[string]*Host{}
	api, err := rpc.NewLocalApi()
	if err != nil {
		panic(err)
	}

	//---------
	// Website
	//---------
	site := echo.New()
	site.Use(middleware.Logger())
	site.Use(middleware.Recover())
	site.Use(session.HwayMiddleware())
	landing.RegisterRoutes(site)
	hosts["localhost:3000"] = &Host{Echo: site}

	//---------
	// Gateway
	//---------
	highway := echo.New()
	highway.Use(middleware.Logger())
	highway.Use(middleware.Recover())
	highway.Use(gatewaymiddleware.IPFSMiddleware(api))
	gateway.RegisterRoutes(highway)
	hosts["auth.localhost:3000"] = &Host{Echo: highway}

	// Server
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()

		host := hosts[req.Host]
		if host != nil {
			host.Echo.ServeHTTP(res, req)
			return nil
		}

		// Default to site for unmatched hosts
		site.ServeHTTP(res, req)
		return nil
	})

	// Log startup information using Echo's logger
	fmt.Println("\n----------------------------------")
	fmt.Println("Server Configuration:")
	fmt.Println("\nAvailable endpoints:")
	fmt.Println("➜ http://localhost:3000 (main site)")
	fmt.Println("➜ http://to.localhost:3000/QmHash/... (IPFS content)")
	fmt.Println("----------------------------------")

	e.Logger.Fatal(e.Start(":3000"))
}
