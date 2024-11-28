package main

import (
	_ "embed"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/onsonr/sonr/pkg/common/middleware/session"
	"github.com/onsonr/sonr/web/gateway"
	"github.com/onsonr/sonr/web/landing"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func main() {
	// Setup Echo
	hosts := map[string]*Host{}

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
	highway.Use(session.MotrMiddleware(nil))

	// Custom error handler for gateway
	highway.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			// Log the error if needed
			c.Logger().Errorf("Gateway error: %v", he.Message)
		}
		// Redirect to main site
		c.Redirect(http.StatusFound, "http://localhost:3000")
	}

	if err := gateway.RegisterRoutes(highway); err != nil {
		panic(err)
	}

	hosts["to.localhost:3000"] = &Host{Echo: highway}

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
