// Package gateway provides the default routes for the Sonr hway.
package gateway

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/app/gateway/config"
	"github.com/onsonr/sonr/app/gateway/handlers"
	"github.com/onsonr/sonr/app/gateway/internal/database"
	"github.com/onsonr/sonr/app/gateway/internal/session"
	"github.com/onsonr/sonr/pkg/common/response"
)

func RegisterRoutes(e *echo.Echo, env config.Env) error {
	// Custom error handler for gateway
	e.HTTPErrorHandler = response.RedirectOnError("http://localhost:3000")

	// Initialize database
	db, err := database.InitDB(env)
	if err != nil {
		return err
	}

	// Inject session middleware with database connection
	e.Use(session.Middleware(db, env))

	// Register routes
	e.GET("/", handlers.HandleIndex)
	e.GET("/register", handlers.HandleRegisterView)
	e.POST("/register/start", handlers.HandleRegisterStart)
	e.POST("/register/finish", handlers.HandleRegisterFinish)
	return nil
}
