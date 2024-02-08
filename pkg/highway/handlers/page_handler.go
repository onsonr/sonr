package handlers

import (
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"

	ui "github.com/sonrhq/sonr/pkg/highway/components"
	"github.com/sonrhq/sonr/pkg/highway/pages"
)

// MountHTMX mounts the HTMX routes
func MountHTMX(or chi.Router) {
	r := chi.NewRouter()
	registerPages(r)
	registerModals(r)
	or.Mount("/", r)
}

// RegisterPages registers the page routes
func registerPages(r chi.Router) {
	r.Handle("/", templ.Handler(pages.HomePage()))
	r.Handle("/console", templ.Handler(pages.ConsolePage()))
	r.Handle("/explorer", templ.Handler(pages.ExplorerPage()))
}

// RegisterModals registers the modal routes
func registerModals(r chi.Router) {
	r.Handle("/login", templ.Handler(ui.AuthModal()))
	r.Handle("/register", templ.Handler(ui.AuthModal()))
}
