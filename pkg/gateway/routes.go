// Package gateway provides the default routes for the Sonr hway.
package gateway

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common/response"
	"github.com/onsonr/sonr/pkg/gateway/config"
	"github.com/onsonr/sonr/pkg/gateway/handlers"
	"github.com/onsonr/sonr/pkg/gateway/internal/database"
	"github.com/onsonr/sonr/pkg/gateway/internal/session"
)

func RegisterRoutes(e *echo.Echo, env config.Env) {
	// Custom error handler for gateway
	e.HTTPErrorHandler = response.RedirectOnError("http://localhost:3000")

	// Inject session middleware
	e.Use(session.Middleware(env))
	e.Use(database.Middleware(env))
	// Register routes
	e.GET("/", handlers.HandleIndex)
	e.GET("/register", handlers.HandleRegisterView(env))
	e.POST("/register/start", handlers.HandleRegisterStart)
	e.POST("/register/finish", handlers.HandleRegisterFinish)
}
