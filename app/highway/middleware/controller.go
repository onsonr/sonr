package middleware

import (
	"net/http"

	"github.com/segmentio/ksuid"
)

// Controller wraps the provided http.Handler with a ControllerMiddleware instance.
// The middleware handles setting a session cookie for tracking user sessions.
func Controller(next http.Handler) http.Handler {
	mw := ControllerMiddleware{
		Next:     next,
		Secure:   true,
		HTTPOnly: true,
	}
	return mw
}

// ControllerMiddleware defines a middleware with context helpers like secure, HTTPOnly flags.
type ControllerMiddleware struct {
	Next     http.Handler
	Secure   bool
	HTTPOnly bool
}

// Address returns the address of the user from the session cookie.
func Address(r *http.Request) (id string) {
	cookie, err := r.Cookie("sonrAddress")
	if err != nil {
		return
	}
	return cookie.Value
}

// ServeHTTP implements the http.Handler interface.
func (mw ControllerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := Address(r)
	if id == "" {
		id = ksuid.New().String()
		http.SetCookie(w, &http.Cookie{Name: "sonrAddress", Value: id, Secure: mw.Secure, HttpOnly: mw.HTTPOnly})
	}
	mw.Next.ServeHTTP(w, r)
}
