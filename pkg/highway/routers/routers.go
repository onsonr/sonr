package routers

import (
	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/pkg/highway/routers/api"
)

// MountAPI mounts all routes to the router
func MountAPI(gr chi.Router) {
	r := chi.NewRouter()
	api.BankHandler.RegisterRoutes(r)
	api.GovHandler.RegisterRoutes(r)
	api.NodeHandler.RegisterRoutes(r)
	api.StakingHandler.RegisterRoutes(r)
	gr.Mount("/api", r)
}

// MountSSE mounts all routes to the router
func MountSSE(r chi.Router) {
	r.Mount(sseEndpoints())
}

func sseEndpoints() (string, chi.Router) {
	r := chi.NewRouter()
	// moduleHandler := htmx.ModuleHandler{}
	// r.Get("/", moduleHandler.IndexPage)
	return "/events", r
}
