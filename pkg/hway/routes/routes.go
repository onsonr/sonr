package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/hway/handlers"
)

func RegisterGatewayAPI(e *echo.Echo) {
}

func RegisterGatewayViews(e *echo.Echo) {
	e.GET("/", handlers.RenderHome)
	e.GET("/login", handlers.RenderLogin)
	e.GET("/register", handlers.RenderRegister)
}
