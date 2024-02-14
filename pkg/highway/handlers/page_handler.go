package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/sonrhq/sonr/pkg/highway/components/pages"
	 "github.com/sonrhq/sonr/pkg/highway/components/modals"
	"github.com/sonrhq/sonr/pkg/highway/middleware"
)

// MountHTMX mounts the HTMX routes
func MountHTMX(or *echo.Echo) {
	registerPages(or)
	registerModals(or)
}

// RegisterPages registers the page routes
func registerPages(e *echo.Echo) {
	e.GET("/", middleware.ShowTempl(pages.RegisterPage()))
	e.GET("/console", middleware.ShowTempl(pages.ConsolePage()))
	e.GET("/explorer", middleware.ShowTempl(pages.ExplorerPage()))
}

// RegisterModals registers the modal routes
func registerModals(e *echo.Echo) {
	e.GET("/login", middleware.ShowTempl(modals.LoginModal()))
	e.GET("/wallet", middleware.ShowTempl(modals.WalletModal()))
}
