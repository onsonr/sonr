//go:build js && wasm

package vault

import (
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/onsonr/sonr/internal/config/motr"
	"github.com/onsonr/sonr/internal/database/motrorm"
	"github.com/onsonr/sonr/vault/context"
)

type Vault = *echo.Echo

// New returns a new Vault instance
func New(config *motr.Config, dbq *motrorm.Queries) (Vault, error) {
	e := echo.New()
	// Override default behaviors
	e.IPExtractor = echo.ExtractIPDirect()
	e.HTTPErrorHandler = handleError()

	// Built-in middleware
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(context.WASMMiddleware)
	registerRoutes(e)
	return e, nil
}

func handleError() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			// Log the error if needed
			c.Logger().Errorf("Error: %v", he.Message)
		}
	}
}

// RegisterRoutes registers the Decentralized Web Node API routes.
func registerRoutes(e *echo.Echo) {
}
