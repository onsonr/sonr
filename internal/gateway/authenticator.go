package gateway

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type RouteHandler func(w http.ResponseWriter, r *http.Request)

// Authenticator represents the interface for session-based authentication.
type Authenticator interface {
	// StartSession starts a new session for a user. It returns the session ID.
	StartSession(w http.ResponseWriter, r *http.Request, values ...SessionValue) (string, error)
	EndSession(w http.ResponseWriter, r *http.Request, sessionID string) error
	IsValidSessionID(r *http.Request, sessionID string) bool
	GetSession(r *http.Request, sessionID string) (*Session, error)

	// Serve serves the fiber app.
	Serve(router *mux.Router)
}

// NewAuthenticator returns a new Authenticator instance. It also initializes the underlying fiber app.
// this is used to have authenticated routes.
func NewAuthenticator() Authenticator {
	return &authenticator{
		getRoutes:  make(map[string]RouteHandler),
		postRoutes: make(map[string]RouteHandler),
	}
}

// authenticator implements the Authenticator interface.
type authenticator struct {
	getRoutes  map[string]RouteHandler
	postRoutes map[string]RouteHandler
}

// Change parameter type
func (a *authenticator) StartSession(w http.ResponseWriter, r *http.Request, values ...SessionValue) (string, error) {
	// Use gorilla/sessions package
	return "", fmt.Errorf("not implemented")
}

func (a *authenticator) EndSession(w http.ResponseWriter, r *http.Request, sessionID string) error {
	// Use gorilla/sessions package
	return fmt.Errorf("not implemented")
}

func (a *authenticator) IsValidSessionID(r *http.Request, sessionID string) bool {
	// Use gorilla/sessions package
	return false
}

func (a *authenticator) GetSession(r *http.Request, sessionID string) (*Session, error) {
	// Use gorilla/sessions package
	return nil, fmt.Errorf("not implemented")
}

// Serve serves the fiber app.
func (a *authenticator) Serve(rtr *mux.Router) {
	for path, handler := range a.getRoutes {
		rtr.HandleFunc(path, handler).Methods("GET")
	}
	for path, handler := range a.postRoutes {
		rtr.HandleFunc(path, handler).Methods("POST")
	}
}

func (a *authenticator) GET(path string, handler RouteHandler) {
	if a.getRoutes == nil {
		a.getRoutes = make(map[string]RouteHandler)
	}
	a.getRoutes[path] = handler
}

func (a *authenticator) POST(path string, handler RouteHandler) {
	if a.postRoutes == nil {
		a.postRoutes = make(map[string]RouteHandler)
	}
	a.postRoutes[path] = handler
}
