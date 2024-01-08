package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/gateway/handlers"
	"github.com/sonrhq/sonr/gateway/templates/components"
	"github.com/sonrhq/sonr/gateway/templates/views"
	"github.com/sonrhq/sonr/pkg/render"
)

func HomeRoutes() chi.Router {
    r := chi.NewRouter()
    bookHandler := handlers.HomeHandler{}
    r.Get("/", bookHandler.ViewHome)
    r.Get("/counts", render.TemplComponent(components.Page(9, 1)))
    r.Get("/{id}", bookHandler.GetBooks)
    r.Put("/{id}", bookHandler.UpdateBook)
    r.Delete("/{id}", bookHandler.DeleteBook)
    return r
}

func ExplorerRoutes() chi.Router {
    r := chi.NewRouter()
    bookHandler := handlers.HomeHandler{}
    r.Get("/", render.TemplComponent(views.Home("test")))
    r.Get("/counts", render.TemplComponent(components.Page(9, 1)))
    r.Get("/{id}", bookHandler.GetBooks)
    r.Put("/{id}", bookHandler.UpdateBook)
    r.Delete("/{id}", bookHandler.DeleteBook)
    return r
}
