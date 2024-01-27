package highway

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/app/highway/middleware"
	"github.com/sonrhq/sonr/app/highway/routes"
)

// Start starts the highway server
func Start() {
	r := chi.NewRouter()
	middleware.UseDefaults(r)
	r.Use(middleware.Session)
	routes.Mount(r)
	http.ListenAndServe(":8080", r)
}
