package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/gateway/handlers"
)

func HomeRoutes() chi.Router {
    r := chi.NewRouter()
    bookHandler := handlers.HomeHandler{}
    r.Get("/", bookHandler.ViewPage)
    r.Get("/{id}", bookHandler.GetBooks)
    r.Put("/{id}", bookHandler.UpdateBook)
    r.Delete("/{id}", bookHandler.DeleteBook)
    return r
}

func ExplorerRoutes() chi.Router {
    r := chi.NewRouter()
    bookHandler := handlers.ExplorerHandler{}
    r.Get("/", bookHandler.ViewPage)
    r.Get("/{id}", bookHandler.GetBooks)
    r.Put("/{id}", bookHandler.UpdateBook)
    r.Delete("/{id}", bookHandler.DeleteBook)
    return r
}
