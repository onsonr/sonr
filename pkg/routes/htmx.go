package routes

import (
	"github.com/labstack/echo/v4"

	handlers "github.com/sonrhq/sonr/pkg/handlers/ui"
	"github.com/sonrhq/sonr/pkg/middleware/common"
)

// RegisterPages registers the page routes for HTMX
func RegisterPages(e *echo.Echo, assetsDir string) {
	e.Static("/*", assetsDir)
	e.GET("/", handlers.Pages.Index, common.UseHTMX)
	e.GET("/404", handlers.Pages.Error, common.UseHTMX)
}

// RegisterModals registers the modal routes for HTMX
func RegisterModals(e *echo.Echo) {
	e.GET("/swap", handlers.Modals.Swap, common.UseHTMX)
	e.GET("/deposit", handlers.Modals.Deposit, common.UseHTMX)
	e.GET("/settings", handlers.Modals.Settings, common.UseHTMX)
	e.GET("/share", handlers.Modals.Share, common.UseHTMX)
}
