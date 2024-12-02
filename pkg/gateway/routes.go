package gateway

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/gateway/handlers"
)

func RegisterRoutes(e *echo.Echo) {
	// Custom error handler for gateway
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			// Log the error if needed
			c.Logger().Errorf("Gateway error: %v", he.Message)
		}
		// Redirect to main site
		c.Redirect(http.StatusFound, "http://localhost:3000")
	}
	e.POST("/spawn", handlers.SpawnVault)
	e.POST("/pin", handlers.PinVault)
	e.POST("/publish", handlers.PublishVault)
}
