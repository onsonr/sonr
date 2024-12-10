//go:build js && wasm
// +build js,wasm

// Package vault provides the routes for the Decentralized Web Node (...or Sonr Motr).
package vault

import (
	"github.com/labstack/echo/v4"

	session "github.com/onsonr/sonr/internal/vault/session"
	"github.com/onsonr/sonr/internal/vault/types"
)

// RegisterRoutes registers the Decentralized Web Node API routes.
func RegisterRoutes(e *echo.Echo, config *types.Config) {
	e.Use(session.Middleware(config))
}
