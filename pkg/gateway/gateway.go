package gateway

import "github.com/go-chi/chi/v5"

type Gateway interface {
    RegisterRoutes(mux *chi.Mux) error
}
