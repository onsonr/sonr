package config

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type ProtocolConfig struct {
	*fiber.App
	SessionStore  *session.Store
	ClientContext client.Context
}

func NewConfig(ctx client.Context) *ProtocolConfig {
	return &ProtocolConfig{
		App: fiber.New(fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		}),
		ClientContext: ctx,
		SessionStore:  session.New(),
	}
}
