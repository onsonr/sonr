package handlers

import (
	"github.com/go-chi/chi/v5"

	app "github.com/sonrhq/sonr/pkg/highway/app"
	ui "github.com/sonrhq/sonr/pkg/highway/components"
	"github.com/sonrhq/sonr/pkg/highway/middleware"
)

// RegisterPages registers the page routes
func RegisterPages(r chi.Router) {
	r.Get("/", middleware.HTMXResponse(app.HomePage()))
	r.Get("/console", middleware.HTMXResponse(app.ConsolePage()))
	r.Get("/explorer", middleware.HTMXResponse(app.ExplorerPage()))
}

// RegisterModals registers the modal routes
func RegisterModals(r chi.Router) {
	r.Get("/login", middleware.HTMXResponse(ui.AuthModal()))
	r.Get("/register", middleware.HTMXResponse(ui.AuthModal()))
}
