package gateway

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/sonrhq/sonr/gateway/handlers"
	"github.com/sonrhq/sonr/gateway/middleware/render"
	"github.com/sonrhq/sonr/gateway/templates/views"
)

func Start() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Get("/", render.TemplComponent(views.Home("test")))
    r.Mount("/books", BookRoutes())

    http.ListenAndServe(":3000", r)
}

func BookRoutes() chi.Router {
    r := chi.NewRouter()
    bookHandler := handlers.HomeHandler{}
    r.Get("/", bookHandler.ListBooks)
    r.Post("/", bookHandler.CreateBook)
    r.Get("/{id}", bookHandler.GetBooks)
    r.Put("/{id}", bookHandler.UpdateBook)
    r.Delete("/{id}", bookHandler.DeleteBook)
    return r
}
