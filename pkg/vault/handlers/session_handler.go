package handlers

import (
	"github.com/di-dao/sonr/pkg/vault/components"
	"github.com/di-dao/sonr/pkg/vault/middleware"
	"github.com/labstack/echo/v4"
)

func HandleHomePage(e echo.Context) error {
	return middleware.Render(e, components.Home())
}
