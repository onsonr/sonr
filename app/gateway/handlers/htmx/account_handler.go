package htmx

import (
	"net/http"

	"github.com/sonrhq/sonr/app/gateway/ui/views"
)

type DashboardHandler struct{}

func (b DashboardHandler) IndexPage(w http.ResponseWriter, r *http.Request) {
	err := views.AccountHome().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
