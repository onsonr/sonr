package gateway

import (
	"net/http"
	// "strconv"
	// "time"

	// "github.com/alexandrevicenzi/go-sse"
	"github.com/go-chi/chi/v5"
	chimdw "github.com/go-chi/chi/v5/middleware"
	"github.com/sonrhq/identity"
	"github.com/sonrhq/service"

	"github.com/sonrhq/sonr/gateway/handlers/home"
	"github.com/sonrhq/sonr/gateway/handlers/search"
)

func Start() {
	r := chi.NewRouter()
	r.Use(chimdw.Compress(10))
	r.Use(chimdw.RequestID)
	r.Use(chimdw.RealIP)
	r.Use(chimdw.Logger)
	r.Use(chimdw.Recoverer)
    r.Mount("/", home.Routes())
	r.Mount("/search", search.Routes())
	identity.RegisterGateway(r)
	service.RegisterGateway(r)
	http.ListenAndServe(":8080", r)
}
