package home

import (
	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/gateway/templates/components"
	"github.com/sonrhq/sonr/pkg/render"
)

func HomeRoutes() chi.Router {
    r := chi.NewRouter()
    bookHandler := HomeHandler{}
    r.Get("/", bookHandler.RenderView)
    r.Get("/counts", render.TemplComponent(components.Page(9, 1)))
    r.Get("/{id}", bookHandler.GetBooks)
    r.Put("/{id}", bookHandler.UpdateBook)
    r.Delete("/{id}", bookHandler.DeleteBook)
    return r
}
