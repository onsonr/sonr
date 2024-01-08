package gateway

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sonrhq/identity"
	"github.com/sonrhq/service"
	"github.com/sonrhq/sonr/gateway/handlers"
)

func Start() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    handlers.RegisterGateway(r)
    identity.RegisterGateway(r)
    service.RegisterGateway(r)
    http.ListenAndServe(":8080", r)
}
