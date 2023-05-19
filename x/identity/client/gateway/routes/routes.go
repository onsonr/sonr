package routes

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/helmet/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/x/identity/client/gateway/config"
	"github.com/sonrhq/core/x/identity/client/gateway/handler"
)

func SetupRoutes(c *config.ProtocolConfig) {
	// Middleware
	c.Use(cors.New())
	c.Use(helmet.New())

	// -- RESTRICTED METHODS --
	c.Use(jwtware.New(jwtware.Config{
		SigningKey: local.Context().SigningKey(),
	}))

	// Account Methods
	c.Get("/accounts", timeout.New(handler.ListAccounts, time.Second*5))
	c.Get("/accounts/:address", timeout.New(handler.GetAccount, time.Second*5))
	c.Get("/accounts/create/:coin_type/:name", timeout.New(handler.CreateAccount, time.Second*5))
	c.Post("/accounts/:address/sign", timeout.New(handler.SignWithAccount, time.Second*5))
	c.Post("/accounts/:address/verify", timeout.New(handler.VerifyWithAccount, time.Second*5))

	// Staking Methods
	c.Get("/staking/validators", timeout.New(handler.ListValidators, time.Second*5))
	c.Get("/staking/validators/:address", timeout.New(handler.ListDelegators, time.Second*5))
	c.Post("/staking/validators/:address/delegate", timeout.New(handler.SubmitDelegate, time.Second*5))
	c.Post("/staking/validators/:address/undelegate", timeout.New(handler.SubmitUndelegate, time.Second*5))
	c.Post("/staking/validators/:address/cancel", timeout.New(handler.SubmitCancelUnbondingDelegation, time.Second*5))

	// Mailbox Methods
	c.Get("/mailbox/:address/read", timeout.New(handler.ReadInboxMessages, time.Second*5))
	c.Post("/mailbox/:address/send/:to", timeout.New(handler.SendInboxMessage, time.Second*5))
}
