package handlers

import (
	pages "github.com/di-dao/sonr/pkg/vault/components"
	"github.com/di-dao/sonr/pkg/vault/middleware"
	"github.com/labstack/echo/v4"
)

func HandleLoginPage(e echo.Context) error {
	return middleware.Render(e, pages.Home())
}
