package routes

import (
	"github.com/di-dao/sonr/pkg/vault/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterPages(e *echo.Echo) {
	e.GET("/", handlers.HandleHomePage)
}
