// Handlers for the router yet to be implemented
package handlers

import (
	"net/http"
)

type ExplorerHandler struct {
}

func (b ExplorerHandler) ListBooks(w http.ResponseWriter, r *http.Request)  {}
func (b ExplorerHandler) GetBooks(w http.ResponseWriter, r *http.Request)   {}
func (b ExplorerHandler) CreateBook(w http.ResponseWriter, r *http.Request) {}
func (b ExplorerHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {}
func (b ExplorerHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {}
