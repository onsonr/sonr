package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/app/gateway/handlers/htmx"
)

func HomeEndpoints() chi.Router {
	r := chi.NewRouter()
	homeHandler := htmx.LandingHandler{}
	r.Get("/", homeHandler.IndexPage)
	return r
}

func ConsoleEndpoints() chi.Router {
	r := chi.NewRouter()
	consoleHandler := htmx.ConsoleHandler{}
	r.Get("/", consoleHandler.IndexPage)
	return r
}

func DashboardEndpoints() chi.Router {
	r := chi.NewRouter()
	dashHandler := htmx.DashboardHandler{}
	r.Get("/", dashHandler.IndexPage)
	return r
}

func ModuleEndpoints() chi.Router {
	r := chi.NewRouter()
	// moduleHandler := htmx.ModuleHandler{}
	// r.Get("/", moduleHandler.IndexPage)
	return r
}

func SSEEndpoints() chi.Router {
	r := chi.NewRouter()
	// moduleHandler := htmx.ModuleHandler{}
	// r.Get("/", moduleHandler.IndexPage)
	return r
}
