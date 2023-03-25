package protocol

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/helmet/v2"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var hway *Protocol

type Protocol struct {
	ctx client.Context
}

func RegisterHighway(ctx client.Context) {
	if isFiber() {
		setupFiber(ctx)

	} else {
		setupFiber(ctx)
	}
}



func setupFiber(ctx client.Context) {
	hway = &Protocol{ctx: ctx}
	app := fiber.New(fiber.Config{
        JSONEncoder: json.Marshal,
        JSONDecoder: json.Unmarshal,
    })
	app.Use(cors.New())

	app.Use(helmet.New())
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK. Highway version v0.6.0. Running on HTTP/TLS")
	})
	app.Post("/highway/auth/keygen", timeout.New(Keygen, time.Second*10))
	app.Post("/highway/auth/login", timeout.New(Login, time.Second*10))
	app.Get("/highway/query/service/:origin", QueryService)
	app.Get("/highway/query/document/:did", QueryDocument)
	app.Post("/highway/vault/add", timeout.New(AddShare, time.Second*5))
	app.Post("/highway/vault/sync", timeout.New(SyncShare, time.Second*5))
	go hway.serveFiber(app)
}

func (p *Protocol) serveConnect(mux *http.ServeMux) {
	if hasTLSCert() {
		http.ListenAndServeTLS(
			fmt.Sprintf(":%s", getServerPort()),
			getTLSCert(),
			getTLSKey(),
			mux,
		)
	} else {
		http.ListenAndServe(
			fmt.Sprintf(":%s", getServerPort()),
			h2c.NewHandler(mux, &http2.Server{}),
		)
	}
}

func (p *Protocol) serveFiber(app *fiber.App) {
	if hasTLSCert() {
		app.ListenTLS(
			fmt.Sprintf(":%s", getServerPort()),
			getTLSCert(),
			getTLSKey(),
		)
	} else {
		app.Listen(
			fmt.Sprintf(":%s", getServerPort()),
		)
	}
}

func getServerPort() string {
	if port := os.Getenv("CONNECT_SERVER_PORT"); port != "" {
		log.Printf("using CONNECT_SERVER_PORT: %s", port)
		return port
	}
	return "8080"
}

func getTLSCert() string {
	if cert := os.Getenv("CONNECT_SERVER_TLS_CERT"); cert != "" {
		log.Printf("using CONNECT_SERVER_TLS_CERT: %s", cert)
		return cert
	}
	return ""
}

func getTLSKey() string {
	if key := os.Getenv("CONNECT_SERVER_TLS_KEY"); key != "" {
		log.Printf("using CONNECT_SERVER_TLS_KEY: %s", key)
		return key
	}
	return ""
}

func hasTLSCert() bool {
	return getTLSCert() != "" && getTLSKey() != "" && !isDev()
}

func isDev() bool {
	return os.Getenv("ENVIRONMENT") == "dev"
}

func isFiber() bool {
	return os.Getenv("HIGHWAY_MODE") != "connect"
}
