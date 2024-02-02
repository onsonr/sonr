package nebula

import (
	"embed"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//go:embed assets
var assets embed.FS

// ServeAssets serves the assets from the embed.FS including stylesheets, images, and javascript files.
func ServeAssets(r chi.Router) {
	r.Handle("/*", http.FileServer(http.FS(assets)))
}
