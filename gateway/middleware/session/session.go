package session

import (
	"net/http"

	"github.com/segmentio/ksuid"
)

type MiddlewareOpts func(*Middleware)

func SetSessionMiddleWare(next http.Handler, opts ...MiddlewareOpts) http.Handler {
	mw := Middleware{
		Next:     next,
		Secure:   true,
		HTTPOnly: true,
	}
	for _, opt := range opts {
		opt(&mw)
	}
	return mw
}

func WithSecure(secure bool) MiddlewareOpts {
	return func(m *Middleware) {
		m.Secure = secure
	}
}

func WithHTTPOnly(httpOnly bool) MiddlewareOpts {
	return func(m *Middleware) {
		m.HTTPOnly = httpOnly
	}
}

type Middleware struct {
	Next     http.Handler
	Secure   bool
	HTTPOnly bool
}

func ID(r *http.Request) (id string) {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		return
	}
	return cookie.Value
}

func (mw Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

	id := ID(r)
	if id == "" {
		id = ksuid.New().String()
		http.SetCookie(w, &http.Cookie{Name: "sessionID", Value: id, Secure: mw.Secure, HttpOnly: mw.HTTPOnly})
	}
	mw.Next.ServeHTTP(w, r)
}
