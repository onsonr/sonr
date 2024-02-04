package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chimdw "github.com/go-chi/chi/v5/middleware"
	"github.com/segmentio/ksuid"
)

// UseDefaults adds chi provided middleware libraries to the router.
func UseDefaults(r *chi.Mux) {
	r.Use(chimdw.Compress(10))
	r.Use(chimdw.RequestID)
	r.Use(chimdw.RealIP)
	r.Use(chimdw.Logger)
	r.Use(chimdw.Recoverer)
	r.Use(Session)
}

// Session wraps the provided http.Handler with a SessionMiddleware instance.
func Session(next http.Handler) http.Handler {
	mw := SessionMiddleware{
		Next:     next,
		Secure:   true,
		HTTPOnly: true,
	}
	return mw
}

// SessionMiddleware defines a middleware with context helpers like secure, HTTPOnly flags.
type SessionMiddleware struct {
	Next     http.Handler
	Secure   bool
	HTTPOnly bool
}

// SessionID returns the address of the user from the session cookie.
func SessionID(r *http.Request) (id string) {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		return
	}
	return cookie.Value
}

// ServeHTTP implements the http.Handler interface.
func (mw SessionMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// s := sse.NewServer(nil)
	// defer s.Shutdown()
	// r.Mount("/events/", s)

	//     go func() {
	//         for {
	//             s.SendMessage("/events/channel-1", sse.SimpleMessage(time.Now().Format("2006/02/01/ 15:04:05")))
	//             time.Sleep(1 * time.Second)
	//         }
	//     }()

	//     go func() {
	//         i := 0
	//         for {
	//             i++
	//             s.SendMessage("/events/channel-2", sse.SimpleMessage(strconv.Itoa(i)))
	//             time.Sleep(1 * time.Second)
	//         }
	//     }()

	id := SessionID(r)
	if id == "" {
		id = ksuid.New().String()
		http.SetCookie(w, &http.Cookie{Name: "sessionID", Value: id, Secure: mw.Secure, HttpOnly: mw.HTTPOnly})
	}
	mw.Next.ServeHTTP(w, r)
}
