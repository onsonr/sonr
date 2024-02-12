package handlers

import (
	"github.com/labstack/echo/v4"

	ui "github.com/sonrhq/sonr/pkg/highway/components"
	"github.com/sonrhq/sonr/pkg/highway/middleware"
	"github.com/sonrhq/sonr/pkg/highway/pages"
)

// MountHTMX mounts the HTMX routes
func MountHTMX(or *echo.Echo) {
	registerPages(or)
	registerModals(or)
}

// RegisterPages registers the page routes
func registerPages(e *echo.Echo) {
	e.GET("/", middleware.ShowTempl(pages.HomePage()))
	e.GET("/console", middleware.ShowTempl(pages.ConsolePage()))
	e.GET("/explorer", middleware.ShowTempl(pages.ExplorerPage()))
}

// RegisterModals registers the modal routes
func registerModals(e *echo.Echo) {
	e.GET("/login", middleware.ShowTempl(ui.AuthModal()))
	e.GET("/register", middleware.ShowTempl(ui.AuthModal()))
}
