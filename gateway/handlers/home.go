// Handlers for the router yet to be implemented
package handlers

import (
	"net/http"

	"github.com/sonrhq/sonr/gateway/templates/views"
)

type HomeHandler struct {
}

func (b HomeHandler) ViewPage(w http.ResponseWriter, r *http.Request) {
	err := views.Home("test").Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (b HomeHandler) GetBooks(w http.ResponseWriter, r *http.Request)   {}
func (b HomeHandler) CreateBook(w http.ResponseWriter, r *http.Request) {}
func (b HomeHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {}
func (b HomeHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {}
