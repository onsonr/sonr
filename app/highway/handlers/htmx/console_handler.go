package htmx

import (
	"net/http"

	"github.com/sonrhq/sonr/pkg/nebula/layouts"
)

type ConsoleHandler struct{}

func (b ConsoleHandler) IndexPage(w http.ResponseWriter, r *http.Request) {
	err := layouts.ConsoleHome().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
