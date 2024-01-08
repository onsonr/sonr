// Handlers for the router yet to be implemented
package handlers

import (
	"net/http"

	"github.com/sonrhq/sonr/gateway/templates/components"
)

type ExplorerHandler struct {
}

func (b ExplorerHandler) ViewPage(w http.ResponseWriter, r *http.Request) {
	err := components.Page(0,1).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func (b ExplorerHandler) ListBooks(w http.ResponseWriter, r *http.Request)  {}
func (b ExplorerHandler) GetBooks(w http.ResponseWriter, r *http.Request)   {}
func (b ExplorerHandler) CreateBook(w http.ResponseWriter, r *http.Request) {}
func (b ExplorerHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {}
func (b ExplorerHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {}
