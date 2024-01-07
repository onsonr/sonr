package gateway

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/sonrhq/sonr/gateway/routes"
)

func Start() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Mount("/", routes.HomeRoutes())
    http.ListenAndServe(":8080", r)
}
