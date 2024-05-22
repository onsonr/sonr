package module

import (
	"github.com/di-dao/core/x/did/handlers"
	"github.com/labstack/echo/v4"
)

func SetRouterProxy(e *echo.Echo) {
	e.GET("/register", handlers.RenderRegisterModal)
}

func SetRouterLocal(e *echo.Echo) {
	e.GET("/", handlers.RenderRegisterModal)
}
