package gateway

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common/middleware/session"
	"github.com/onsonr/sonr/pkg/webapp"
)

func New() *echo.Echo {
	e := echo.New()
	e.Use(session.HwayMiddleware())

	// Add WASM-specific routes
	webapp.RegisterLandingFrontend(e)
	return e
}
