package middleware

import (
	"net/http"

	"github.com/segmentio/ksuid"
)

func Provider(next http.Handler) http.Handler {
	mw := ControllerMiddleware{
		Next:     next,
		Secure:   true,
		HTTPOnly: true,
	}
	return mw
}

type ProviderMiddleware struct {
	Next     http.Handler
	Secure   bool
	HTTPOnly bool
}

func Origin(r *http.Request) (id string) {
	cookie, err := r.Cookie("origin")
	r.Header.Add("Access-Control-Allow-Origin", "*")

	if err != nil {
		return
	}
	return cookie.Value
}

func (mw ProviderMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := Origin(r)
	if id == "" {
		id = ksuid.New().String()
		http.SetCookie(w, &http.Cookie{Name: "sonrAddress", Value: id, Secure: mw.Secure, HttpOnly: mw.HTTPOnly})
	}
	mw.Next.ServeHTTP(w, r)
}
