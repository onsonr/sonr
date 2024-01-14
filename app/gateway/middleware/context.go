package middleware

import (
	"net/http"

	"github.com/segmentio/ksuid"
)

func Context(next http.Handler) http.Handler {
	mw := ControllerMiddleware{
		Next:     next,
		Secure:   true,
		HTTPOnly: true,
	}
	return mw
}

type ContextMiddleware struct {
	Next     http.Handler
	Secure   bool
	HTTPOnly bool
}

func GrpcClientConn(r *http.Request) (id string) {
	cookie, err := r.Cookie("sonrAddress")
	if err != nil {
		return
	}
	return cookie.Value
}

func (mw ContextMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := Address(r)
	if id == "" {
		id = ksuid.New().String()
		http.SetCookie(w, &http.Cookie{Name: "sonrAddress", Value: id, Secure: mw.Secure, HttpOnly: mw.HTTPOnly})
	}
	mw.Next.ServeHTTP(w, r)
}
