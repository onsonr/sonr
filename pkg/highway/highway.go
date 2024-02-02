package highway

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pterm/pterm"

	"github.com/sonrhq/sonr/pkg/highway/middleware"
	"github.com/sonrhq/sonr/pkg/highway/routers"
	"github.com/sonrhq/sonr/pkg/nebula"
)

// Start starts the highway server
func Start() {
	pterm.DefaultHeader.Printf(PersistentHeader)
	r := chi.NewRouter()
	middleware.UseDefaults(r)
	r.Use(middleware.Session)
	routers.MountAPI(r)
	routers.MountSSE(r)
	nebula.ServeAssets(r)
	http.ListenAndServe(":8000", r)
}

// PersistentHeader is the header that is printed on start
const PersistentHeader = `
Sonr Highway
· Gateway: http://localhost:8000
· Node RPC: http://localhost:26657
`
