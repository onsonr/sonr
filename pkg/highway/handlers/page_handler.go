package handlers

import (
	"github.com/go-chi/chi/v5"

	ui "github.com/sonrhq/sonr/pkg/highway/components"
	"github.com/sonrhq/sonr/pkg/highway/middleware"
	"github.com/sonrhq/sonr/pkg/highway/pages"
)

// RegisterPages registers the page routes
func RegisterPages(r chi.Router) {
	r.Get("/", middleware.HTMXResponse(pages.HomePage()))
	r.Get("/console", middleware.HTMXResponse(pages.ConsolePage()))
	r.Get("/explorer", middleware.HTMXResponse(pages.ExplorerPage()))
}

// RegisterModals registers the modal routes
func RegisterModals(r chi.Router) {
	r.Get("/login", middleware.HTMXResponse(ui.AuthModal()))
	r.Get("/register", middleware.HTMXResponse(ui.AuthModal()))
}
