package gateway

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/helmet/v2"
	"github.com/kataras/go-sessions/v3"
	"github.com/sonrhq/core/internal/local"
	"github.com/valyala/fasthttp"
)

// Authenticator represents the interface for session-based authentication.
type Authenticator interface {
	// Router returns the underlying fiber app.
	Router() *fiber.App

	// StartSession starts a new session for a user. It returns the session ID.
	StartSession(ctx *fasthttp.RequestCtx, values ...SessionValue) (string, error)

	// EndSession ends the specified session. It returns an error if the session does not exist.
	EndSession(ctx *fasthttp.RequestCtx, sessionID string) error

	// IsValidSessionID checks if the specified session ID is valid (i.e., it corresponds to an active session).
	IsValidSessionID(ctx *fasthttp.RequestCtx, sessionID string) bool

	// GetSession retrieves the session information for a session ID. It returns an error if the session does not exist.
	GetSession(ctx *fasthttp.RequestCtx, sessionID string) (*Session, error)

	// Serve serves the fiber app.
	Serve()
}

// NewAuthenticator returns a new Authenticator instance. It also initializes the underlying fiber app.
// this is used to have authenticated routes.
func NewAuthenticator() Authenticator {
	auth := &authenticator{
		app: fiber.New(fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		}),
	}
	auth.app.Use(cors.New())
	auth.app.Use(helmet.New())
	return auth
}

// authenticator implements the Authenticator interface.
type authenticator struct {
	app *fiber.App
}

// StartSession starts a new session for a user. It returns the session ID.
func (a *authenticator) StartSession(ctx *fasthttp.RequestCtx, values ...SessionValue) (string, error) {
	sess := sessions.StartFasthttp(ctx)
	session := defaultSession()
	for _, value := range values {
		value(session)
	}
	session.Save(sess)
	sessionID := sess.ID()
	return sessionID, nil
}

// EndSession ends the specified session. It returns an error if the session does not exist.
func (a *authenticator) EndSession(ctx *fasthttp.RequestCtx, sessionID string) error {
	sess := sessions.StartFasthttp(ctx)
	sess.Destroy()
	return nil
}

// IsValidSessionID checks if the specified session ID is valid (i.e., it corresponds to an active session).
func (a *authenticator) IsValidSessionID(ctx *fasthttp.RequestCtx, sessionID string) bool {
	sess := sessions.StartFasthttp(ctx)
	return sess.ID() == sessionID
}

// GetSession retrieves the session information for a session ID. It returns an error if the session does not exist.
func (a *authenticator) GetSession(ctx *fasthttp.RequestCtx, sessionID string) (*Session, error) {
	sess := sessions.StartFasthttp(ctx)
	return LoadSession(sess), nil
}

// Router returns the underlying fiber app.
func (a *authenticator) Router() *fiber.App {
	return a.app
}

// Serve serves the fiber app.
func (a *authenticator) Serve() {
	go serveFiber(a.app)
}

// helper function to serve the fiber app.
func serveFiber(app *fiber.App) {
	if local.Context().HasTlsCert() {
		app.ListenTLS(
			local.Context().FiberListenAddress(),
			local.Context().TlsCertPath,
			local.Context().TlsKeyPath,
		)
	} else {
		app.Listen(local.Context().FiberListenAddress())
	}
}
