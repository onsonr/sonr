package htmx

import (
	"net/http"

	"github.com/sonrhq/sonr/app/gateway/ui/views"
)

type LandingHandler struct{}

func (b LandingHandler) IndexPage(w http.ResponseWriter, r *http.Request) {
	err := views.LandingHome().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
