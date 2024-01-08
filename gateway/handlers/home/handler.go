package home

import (
	"net/http"

	home_views "github.com/sonrhq/sonr/gateway/handlers/home/views"
)

type HomeHandler struct {
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                   Operations                                   ||
// ! ||--------------------------------------------------------------------------------||

func (b HomeHandler) GetBooks(w http.ResponseWriter, r *http.Request)   {}
func (b HomeHandler) CreateBook(w http.ResponseWriter, r *http.Request) {}
func (b HomeHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {}
func (b HomeHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {}


// ! ||--------------------------------------------------------------------------------||
// ! ||                                    View HTMX                                   ||
// ! ||--------------------------------------------------------------------------------||
func (b HomeHandler) RenderView(w http.ResponseWriter, r *http.Request) {
	err := home_views.Page("test").Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
