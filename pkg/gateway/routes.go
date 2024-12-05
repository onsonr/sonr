package gateway

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common/session"
	"github.com/onsonr/sonr/pkg/gateway/config"
	"github.com/onsonr/sonr/pkg/gateway/handlers"
	"github.com/onsonr/sonr/pkg/gateway/internal/database"
)

func RegisterRoutes(e *echo.Echo, env config.Env) {
	// Inject session middleware
	e.Use(session.GatewayMiddleware(env))
	e.Use(database.Middleware(env))

	// Custom error handler for gateway
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			// Log the error if needed
			c.Logger().Errorf("Gateway error: %v", he.Message)
		}
		// Redirect to main site
		c.Redirect(http.StatusFound, "http://localhost:3000")
	}

	// Register routes
	e.GET("/", handlers.HandleIndex)
	e.GET("/register", handlers.HandleRegisterView(env))
	e.POST("/register/start", handlers.HandleRegisterStart)
	e.POST("/register/finish", handlers.HandleRegisterFinish)
}
