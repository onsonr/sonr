// Package gateway provides the default routes for the Sonr hway.
package gateway

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/onsonr/sonr/gateway/context"
	"github.com/onsonr/sonr/gateway/handlers"
	"github.com/onsonr/sonr/internal/common"
	config "github.com/onsonr/sonr/internal/config/hway"
	hwayorm "github.com/onsonr/sonr/internal/database/hwayorm"
)

type Gateway = *echo.Echo

// New returns a new Gateway instance
func New(env config.Hway, ipc common.IPFS, dbq *hwayorm.Queries) (Gateway, error) {
	e := echo.New()

	// Built-in middleware
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.IPExtractor = echo.ExtractIPDirect()
	e.Use(echoprometheus.NewMiddleware("hway"))
	e.Use(context.UseGateway(env, ipc, dbq))

	// Register View Handlers
	e.HTTPErrorHandler = handlers.ErrorHandler
	e.GET("/", handlers.IndexHandler)
	handlers.RegisterHandler(e.Group("/register"))
	return e, nil
}
