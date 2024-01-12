package gateway

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/gateway/handlers/console"
	"github.com/sonrhq/sonr/gateway/handlers/explorer"
	"github.com/sonrhq/sonr/gateway/handlers/landing"
	"github.com/sonrhq/sonr/gateway/middleware"
)

func Start() {
	r := chi.NewRouter()
	middleware.UseDefaults(r)
	r.Use(middleware.Session)
    r.Mount("/", landing.Routes())
	r.Mount("/explorer", explorer.Routes())
	r.Mount("/console", console.Routes())
	http.ListenAndServe(":8080", r)
}
