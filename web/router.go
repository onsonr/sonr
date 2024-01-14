package gateway

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/web/htmx/console"
	landing "github.com/sonrhq/sonr/web/htmx/home"
	"github.com/sonrhq/sonr/web/middleware"
)

func Start() {
	r := chi.NewRouter()
	middleware.UseDefaults(r)
	r.Use(middleware.Session)
	r.Mount("/", landing.Routes())
	r.Mount("/console", console.Routes())
	http.ListenAndServe(":8080", r)
}
