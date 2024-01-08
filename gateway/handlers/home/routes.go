package home

import (
	"github.com/go-chi/chi/v5"
)

func Routes() chi.Router {
    r := chi.NewRouter()
    homeHandler := Handler{}
    r.Get("/", homeHandler.IndexPage)
    r.Get("/app", homeHandler.AppPage)
    r.Get("/explorer", homeHandler.AppPage)
    return r
}
