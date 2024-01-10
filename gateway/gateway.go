package gateway

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/gateway/handlers/home"
	"github.com/sonrhq/sonr/gateway/handlers/search"
	"github.com/sonrhq/sonr/gateway/middleware"
)

func Start() {
	r := chi.NewRouter()
	middleware.UseDefaults(r)
	r.Use(middleware.Session)
    r.Mount("/", home.Routes())
	r.Mount("/search", search.Routes())
	http.ListenAndServe(":8080", r)
}
