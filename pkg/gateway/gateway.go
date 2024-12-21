// Package gateway provides the default routes for the Sonr hway.
package gateway

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	config "github.com/onsonr/sonr/internal/config/hway"
	hwayorm "github.com/onsonr/sonr/internal/database/hwayorm"
	"github.com/onsonr/sonr/pkg/common"
	"github.com/onsonr/sonr/pkg/gateway/context"
	"github.com/onsonr/sonr/pkg/gateway/routes"
)

type Gateway = *echo.Echo

// New returns a new Gateway instance
func New(env config.Hway, ipc common.IPFS, dbq *hwayorm.Queries) (Gateway, error) {
	e := echo.New()
	// Override default behaviors
	e.IPExtractor = echo.ExtractIPDirect()
	e.HTTPErrorHandler = handleError()

	// Built-in middleware
	e.Use(echoprometheus.NewMiddleware("hway"))
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(context.UseGateway(env, ipc, dbq))
	routes.Register(e)
	return e, nil
}

func handleError() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			// Log the error if needed
			c.Logger().Errorf("Error: %v", he.Message)
			context.RenderError(c, he)
		}
	}
}
