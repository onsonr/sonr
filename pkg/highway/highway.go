package highway

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pterm/pterm"

	"github.com/sonrhq/sonr/pkg/highway/middleware"
	"github.com/sonrhq/sonr/pkg/highway/routes"
)

// Start starts the highway server
func Start() {
	pterm.DefaultHeader.Printf(PersistentHeader)
	r := chi.NewRouter()
	middleware.UseDefaults(r)
	r.Use(middleware.Session)
	routes.Mount(r)
	http.ListenAndServe(":8000", r)
}

// PersistentHeader is the header that is printed on start
const PersistentHeader = `
Sonr Highway
· Gateway: http://localhost:8000
· Node RPC: http://localhost:26657
`
