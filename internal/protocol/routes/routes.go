package routes

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/helmet/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/internal/protocol/config"
	"github.com/sonrhq/core/internal/protocol/handler"
)

func SetupRoutes(c *config.ProtocolConfig) {
	// Middleware
	c.Use(cors.New())
	c.Use(helmet.New())

	// Query Methods
	c.Get("/id", timeout.New(handler.ListDIDs, time.Second*5))
	c.Get("/id/:did", timeout.New(handler.GetDID, time.Second*5))
	c.Get("/id/alias/:alias", timeout.New(handler.GetDIDByAlias, time.Second*5))
	c.Get("/id/owner/:owner", timeout.New(handler.GetDIDByOwner, time.Second*5))

	// Service Methods
	c.Get("/service", timeout.New(handler.ListServices, time.Second*10))
	c.Get("/service/:origin", timeout.New(handler.GetService, time.Second*10))
	c.Get("/service/:origin/register/start", timeout.New(handler.GetServiceAttestion, time.Second*5))
	c.Post("/service/:origin/register/finish", timeout.New(handler.VerifyServiceAttestion, time.Second*10))
	c.Get("/service/:origin/login/start", timeout.New(handler.GetServiceAssertion, time.Second*5))
	c.Post("/service/:origin/login/finish", timeout.New(handler.VerifyServiceAssertion, time.Second*10))

	// -- RESTRICTED METHODS --
	c.Use(jwtware.New(jwtware.Config{
		SigningKey: local.Context().SigningKey(),
	}))

	// MPC Methods
	c.Get("/accounts", timeout.New(handler.ListAccounts, time.Second*5))
	c.Get("/accounts/:address", timeout.New(handler.GetAccount, time.Second*5))
	c.Post("/accounts/create/:coin_type/:name", timeout.New(handler.CreateAccount, time.Second*5))
	c.Post("/accounts/:address/sign", timeout.New(handler.SignWithAccount, time.Second*5))
	c.Post("/accounts/:address/verify", timeout.New(handler.VerifyWithAccount, time.Second*5))

	// Inbox Methods
	c.Get("/inbox/:address/read", timeout.New(handler.ReadInboxMessages, time.Second*5))
	c.Post("/inbox/:address/send/:to", timeout.New(handler.SendInboxMessage, time.Second*5))
}
