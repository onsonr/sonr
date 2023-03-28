package rest

import (
	"context"
	"errors"
	"regexp"

	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sonrhq/core/internal/protocol/packages/resolver"
	v1 "github.com/sonrhq/core/types/highway/v1"
	"github.com/sonrhq/core/x/identity/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/helmet/v2"
)

type HttpTransport struct {
	*fiber.App
	SessionStore *session.Store
	ClientContext   client.Context
}

func NewHttpTransport(ctx client.Context) *HttpTransport {
	// HttPTransport
	rest := &HttpTransport{
		App: fiber.New(fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		}),
		ClientContext: ctx,
		SessionStore: session.New(),
	}

	// Middleware
	rest.Use(cors.New())
	rest.Use(helmet.New())

	// Status Methods
	rest.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK. Highway version v0.6.0. Running on HTTP/TLS")
	})

	// Query Methods
	rest.Get("/highway/query/service/:origin", timeout.New(rest.QueryService, time.Second*5))
	rest.Get("/highway/query/document/:did", timeout.New(rest.QueryDocument, time.Second*5))

	// Auth Methods
	rest.Post("/highway/auth/challenge", timeout.New(rest.Challenge, time.Second*10))
	rest.Post("/highway/auth/keygen", timeout.New(rest.Keygen, time.Second*10))
	rest.Post("/highway/auth/login", timeout.New(rest.Login, time.Second*10))
	rest.Post("/highway/vault/add", timeout.New(rest.AddShare, time.Second*5))
	rest.Post("/highway/vault/sync", timeout.New(rest.SyncShare, time.Second*5))
	return rest
}

func (htt *HttpTransport) Challenge(c *fiber.Ctx) error {
	store, err := htt.SessionStore.Get(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	req := new(v1.ChallengeRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	params := types.DefaultParams()
	// Get the origin from the request.
	origin := regexp.MustCompile(`[^a-zA-Z]+`).ReplaceAllString(req.Origin, "")
	uuid := req.Uuid

	service, _ := resolver.GetService(context.Background(), origin)
	if service == nil {
		service, _ = resolver.GetService(context.Background(), "localhost")
	}
	// Check if service is still nil - return internal server error
	if service == nil {
		sentry.CaptureException(errors.New("Service not found"))
		return c.Status(500).SendString("Service not found")
	}

	chal, err := service.IssueChallenge()
	if err != nil {
		sentry.CaptureException(err)
		return c.Status(500).SendString(err.Error())
	}

	// Set Store origin/uuid = challenge
	store.Set(challengeUuidStoreKey(origin, uuid), chal)
	ops, err := params.NewWebauthnCreationOptions(service, req.Uuid, chal)
	if err != nil {
		sentry.CaptureException(err)
		return c.Status(500).SendString(err.Error())
	}
	bz, err := json.MarshalIndent(ops, "", "  ")
	if err != nil {
		sentry.CaptureException(err)
		return c.Status(500).SendString(err.Error())
	}
	res := &v1.ChallengeResponse{
		RpId:              service.Name,
		RpName:            service.Name,
		CredentialOptions: string(bz),
	}
	return c.JSON(res)
}

