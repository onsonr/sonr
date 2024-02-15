package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/sonrhq/sonr/pkg/highway/components/layouts"
	"github.com/sonrhq/sonr/pkg/highway/components/modals"
	"github.com/sonrhq/sonr/pkg/highway/components/views"
	"github.com/sonrhq/sonr/pkg/highway/middleware"
)

// RegisterPages registers the page routes for HTMX
func RegisterPages(e *echo.Echo) {
	e.GET("/", middleware.ShowTempl(layouts.RegisterPage()))
	e.GET("/console", middleware.ShowTempl(layouts.ConsolePage()))
	e.GET("/explorer", middleware.ShowTempl(layouts.ExplorerPage()))
}

// RegisterModals registers the modal routes
func RegisterModals(e *echo.Echo) {
	e.GET("/login", middleware.ShowTempl(modals.LoginModal()))
	e.GET("/wallet", middleware.ShowTempl(modals.WalletModal()))
}

// RegisterUtilityPages registers the utility page routes
func RegisterUtilityPages(e *echo.Echo) {
	e.GET("/*", middleware.ShowTempl(views.Error404View()))
}
