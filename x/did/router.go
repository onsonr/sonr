package module

import (
	"github.com/labstack/echo/v4"

	"github.com/di-dao/sonr/x/did/internal/handlers"
)

func SetRouterProxy(e *echo.Echo) {
	e.GET("/register", handlers.RenderRegisterModal)
}

func SetRouterLocal(e *echo.Echo) {
	e.GET("/", handlers.RenderRegisterModal)
}
