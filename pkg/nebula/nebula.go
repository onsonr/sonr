package nebula

import (
	"embed"
	"net/http"
)

// TODO: Add CDN package for Shoelace as go-embed
//	labels: HTMX/Frontend,Plane,Github
//	milestone: 24

//go:embed assets
var assets embed.FS

// ServeAssets serves the assets from the embed.FS including stylesheets, images, and javascript files.
func ServeAssets() (pattern string, handler http.Handler) {
	return "/*", http.FileServer(http.FS(assets))
}
