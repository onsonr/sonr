package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/di-dao/sonr/internal/middleware/htmx"
	ui "github.com/di-dao/sonr/x/did/internal/components"
)

func RenderRegisterModal(c echo.Context) error {
	return htmx.Render(c, ui.AuthModal())
}
