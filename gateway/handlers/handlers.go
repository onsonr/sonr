package handlers

import (
	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/gateway/handlers/home"
	"github.com/sonrhq/sonr/gateway/handlers/search"
)

func RegisterGateway(r *chi.Mux) {
    r.Mount("/", home.Routes())
	r.Mount("/search", search.Routes())
}
