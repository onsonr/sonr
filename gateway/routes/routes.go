package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/gateway/handlers"
	"github.com/sonrhq/sonr/gateway/templates/views"
	"github.com/sonrhq/sonr/pkg/render"
)

func HomeRoutes() chi.Router {
    r := chi.NewRouter()
    bookHandler := handlers.HomeHandler{}
    r.Get("/", bookHandler.ViewHome)
    r.Post("/", bookHandler.CreateBook)
    r.Get("/{id}", bookHandler.GetBooks)
    r.Put("/{id}", bookHandler.UpdateBook)
    r.Delete("/{id}", bookHandler.DeleteBook)
    return r
}

func ExplorerRoutes() chi.Router {
    r := chi.NewRouter()
    bookHandler := handlers.HomeHandler{}
    r.Get("/", render.TemplComponent(views.Home("test")))
    r.Post("/", bookHandler.CreateBook)
    r.Get("/{id}", bookHandler.GetBooks)
    r.Put("/{id}", bookHandler.UpdateBook)
    r.Delete("/{id}", bookHandler.DeleteBook)
    return r
}
