// Package gateway provides the default routes for the Sonr hway.
package gateway

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	config "github.com/onsonr/sonr/internal/config/hway"
	"github.com/onsonr/sonr/internal/database"
	"github.com/onsonr/sonr/pkg/common"
	"github.com/onsonr/sonr/pkg/gateway/middleware"
	"github.com/onsonr/sonr/pkg/gateway/routes"
)

type Gateway = *echo.Echo

// New returns a new Gateway instance
func New(env config.Hway, ipc common.IPFS) (Gateway, error) {
	db, err := database.NewDB(env)
	if err != nil {
		return nil, err
	}
	e := echo.New()
	// Override default behaviors
	e.IPExtractor = echo.ExtractIPDirect()
	e.HTTPErrorHandler = redirectOnError("http://localhost:3000")

	// Built-in middleware
	e.Use(echoprometheus.NewMiddleware("hway"))
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(middleware.UseGateway(env, ipc, db))
	routes.Register(e)
	return e, nil
}

func redirectOnError(target string) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			// Log the error if needed
			c.Logger().Errorf("Error: %v", he.Message)
			middleware.RenderError(c, he)
		}
	}
}
