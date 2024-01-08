package explorer

import (
	"github.com/go-chi/chi/v5"
)

func Routes() chi.Router {
    r := chi.NewRouter()
    bookHandler := Handler{}
    r.Get("/", bookHandler.IndexPage)
    return r
}
