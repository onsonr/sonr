package handlers

import (
	"github.com/di-dao/core/internal/middleware/htmx"
	ui "github.com/di-dao/core/x/did/components"
	"github.com/labstack/echo/v4"
)

func RenderRegisterModal(c echo.Context) error {
	return htmx.Render(c, ui.AuthModal())
}
