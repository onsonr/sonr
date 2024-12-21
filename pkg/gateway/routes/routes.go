package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/gateway/handlers"
)

func Register(e *echo.Echo) error {
	// Register View Handlers
	e.GET("/", handlers.HandleIndex)
	handlers.HandleRegistration(e.Group("/register"))
	return nil
}
