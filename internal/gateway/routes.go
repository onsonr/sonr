// Package gateway provides the default routes for the Sonr hway.
package gateway

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/config"
	"github.com/onsonr/sonr/internal/gateway/context"
	"github.com/onsonr/sonr/internal/gateway/handlers/index"
	"github.com/onsonr/sonr/internal/gateway/handlers/register"
	"github.com/onsonr/sonr/pkg/common/response"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, env config.Env, db *gorm.DB) error {
	// Custom error handler for gateway
	e.HTTPErrorHandler = response.RedirectOnError("http://localhost:3000")

	// Inject session middleware with database connection
	e.Use(context.Middleware(db, env))

	// Register routes
	e.GET("/", index.Handler)
	e.GET("/register", register.HandleCreateProfile)
	e.POST("/register/start", register.HandlePasskeyStart)
	e.POST("/register/finish", register.HandlePasskeyFinish)
	return nil
}
