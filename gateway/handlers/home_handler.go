// Handlers for the router yet to be implemented
package handlers

import "net/http"

type HomeHandler struct {
}

func (b HomeHandler) ListBooks(w http.ResponseWriter, r *http.Request)  {}
func (b HomeHandler) GetBooks(w http.ResponseWriter, r *http.Request)   {}
func (b HomeHandler) CreateBook(w http.ResponseWriter, r *http.Request) {}
func (b HomeHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {}
func (b HomeHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {}
