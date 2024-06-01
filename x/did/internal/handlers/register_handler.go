package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/di-dao/sonr/internal/middleware"
	ui "github.com/di-dao/sonr/x/did/internal/components"
)

func RenderRegisterModal(c echo.Context) error {
	return middleware.Render(c, ui.AuthModal())
}
