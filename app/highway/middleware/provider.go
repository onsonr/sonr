package middleware

import (
	"net/http"

	"github.com/segmentio/ksuid"
)

// Provider wraps the provided http.Handler with a ProviderMiddleware instance.
func Provider(next http.Handler) http.Handler {
	mw := ControllerMiddleware{
		Next:     next,
		Secure:   true,
		HTTPOnly: true,
	}
	return mw
}

// ProviderMiddleware defines a middleware with context helpers like secure, HTTPOnly flags.
type ProviderMiddleware struct {
	Next     http.Handler
	Secure   bool
	HTTPOnly bool
}

// Origin returns the address of the user from the session cookie.
func Origin(r *http.Request) (id string) {
	cookie, err := r.Cookie("origin")
	r.Header.Add("Access-Control-Allow-Origin", "*")

	if err != nil {
		return
	}
	return cookie.Value
}

// ServeHTTP implements the http.Handler interface.
func (mw ProviderMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := Origin(r)
	if id == "" {
		id = ksuid.New().String()
		http.SetCookie(w, &http.Cookie{Name: "sonrAddress", Value: id, Secure: mw.Secure, HttpOnly: mw.HTTPOnly})
	}
	mw.Next.ServeHTTP(w, r)
}
