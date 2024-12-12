// Package gateway provides the default routes for the Sonr hway.
package gateway

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/context"
	"github.com/onsonr/sonr/internal/gateway/handlers"
	"github.com/onsonr/sonr/pkg/common/response"
	config "github.com/onsonr/sonr/pkg/config/hway"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, env config.Hway, db *gorm.DB) error {
	// Custom error handler for gateway
	e.HTTPErrorHandler = response.RedirectOnError("http://localhost:3000")

	// Inject session middleware with database connection
	e.Use(context.Middleware(db, env))

	// Register View Handlers
	e.GET("/", handlers.RenderIndex)
	e.GET("/register", handlers.RenderProfileRegister)
	e.POST("/register/passkey", handlers.RenderPasskeyStart)
	e.POST("/register/finish", handlers.RenderPasskeyFinish)
	return nil
}
