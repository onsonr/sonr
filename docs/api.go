package docs

import (
	"embed"
	httptemplate "html/template"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	apiFile   = "/static/sonr.swagger.yaml"
	indexFile = "template/index.tpl"
)

//go:embed static
var static embed.FS

//go:embed template
var template embed.FS

// RegisterOpenAPIService registers an OpenAPI console service at /docs.
func RegisterOpenAPIService(appName string, rtr *mux.Router) {
	rtr.Handle(apiFile, http.FileServer(http.FS(static)))
	rtr.HandleFunc("/", handler(appName))
}

// handler returns an http handler that servers OpenAPI console for an OpenAPI spec at specURL.
func handler(title string) http.HandlerFunc {
	t, _ := httptemplate.ParseFS(template, indexFile)

	return func(w http.ResponseWriter, _ *http.Request) {
		t.Execute(w, struct {
			Title string
			URL   string
		}{
			title,
			apiFile,
		})
	}
}
