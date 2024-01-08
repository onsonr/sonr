package gateway

import (
	"net/http"
	// "strconv"
	// "time"

	// "github.com/alexandrevicenzi/go-sse"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sonrhq/identity"
	"github.com/sonrhq/service"

	"github.com/sonrhq/sonr/gateway/handlers"
)

func Start() {
    r := chi.NewRouter()
	r.Use(middleware.Compress(10))
    r.Use(middleware.Logger)

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
    handlers.RegisterGateway(r)
    identity.RegisterGateway(r)
    service.RegisterGateway(r)
    http.ListenAndServe(":8080", r)
}
