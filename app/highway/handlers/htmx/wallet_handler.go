package htmx

import (
	"net/http"

	"github.com/sonrhq/sonr/pkg/nebula/layouts"
)

// WalletHandler is a handler for the wallet page
type WalletHandler struct{}

// IndexPage renders the wallet page
func (b WalletHandler) IndexPage(w http.ResponseWriter, r *http.Request) {
	err := layouts.ConsoleHome().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
