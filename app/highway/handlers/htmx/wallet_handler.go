package htmx

import (
	"net/http"

	"github.com/sonrhq/sonr/app/highway/ui/views"
)

type WalletHandler struct{}

func (b WalletHandler) IndexPage(w http.ResponseWriter, r *http.Request) {
	err := views.AccountHome().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
