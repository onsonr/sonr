package htmx

import (
	"net/http"

	"github.com/sonrhq/sonr/pkg/nebula/layouts"
)

// ConsoleHandler is a handler for the console page
type ConsoleHandler struct{}

// IndexPage renders the console page
func (b ConsoleHandler) IndexPage(w http.ResponseWriter, r *http.Request) {
	err := layouts.ConsoleHome().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
