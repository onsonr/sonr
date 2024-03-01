package routes

import (
	"github.com/labstack/echo/v4"

	handlers "github.com/sonrhq/sonr/pkg/handlers/ui"
	"github.com/sonrhq/sonr/pkg/middleware/common"
)

// RegisterHTMXPages registers the page routes for HTMX
func RegisterHTMXPages(e *echo.Echo, assetsDir string) {
	e.Static("/*", assetsDir)
	e.GET("/", handlers.Pages.Index, common.UseHTMX)
	e.GET("/_panels/home", handlers.Pages.Home, common.UseHTMX)
	e.GET("/_panels/chat", handlers.Pages.Chat, common.UseHTMX)
	e.GET("/_panels/wallet", handlers.Pages.Wallet, common.UseHTMX)
	e.GET("/_panels/status", handlers.Pages.Status, common.UseHTMX)
	e.GET("/_panels/governance", handlers.Pages.Governance, common.UseHTMX)
	e.GET("/_panels/console", handlers.Pages.Console, common.UseHTMX)
	e.GET("/error-404", handlers.Pages.Error, common.UseHTMX)
}

// RegisterHTMXModals registers the modal routes for HTMX
func RegisterHTMXModals(e *echo.Echo) {
	e.GET("/swap", handlers.Modals.Swap, common.UseHTMX)
	e.GET("/deposit", handlers.Modals.Deposit, common.UseHTMX)
	e.GET("/settings", handlers.Modals.Settings, common.UseHTMX)
	e.GET("/share", handlers.Modals.Share, common.UseHTMX)
}
