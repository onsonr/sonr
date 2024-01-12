package explorer

import (
	"github.com/go-chi/chi/v5"
)

func Routes() chi.Router {
    r := chi.NewRouter()
    homeHandler := Handler{}
    r.Get("/", homeHandler.IndexPage)
    return r
}
