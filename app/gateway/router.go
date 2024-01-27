package gateway

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/app/gateway/middleware"
	"github.com/sonrhq/sonr/app/gateway/routes"
)

func Start() {
	r := chi.NewRouter()
	middleware.UseDefaults(r)
	r.Use(middleware.Session)
	routes.Mount(r)
	http.ListenAndServe(":8080", r)
}
