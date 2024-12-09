//go:build js && wasm
// +build js,wasm

// Package vault provides the routes for the Decentralized Web Node (...or Sonr Motr).
package vault

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/app/vault/handlers"
	session "github.com/onsonr/sonr/app/vault/internal"
	"github.com/onsonr/sonr/app/vault/types"
)

// RegisterRoutes registers the Decentralized Web Node API routes.
func RegisterRoutes(e *echo.Echo, config *types.Config) {
	e.Use(session.Middleware(config))

	e.GET("/register/:subject/start", handlers.RegisterSubjectStart)
	e.POST("/register/:subject/finish", handlers.RegisterSubjectFinish)

	e.GET("/login/:subject/start", handlers.LoginSubjectStart)
	e.POST("/login/:subject/finish", handlers.LoginSubjectFinish)

	e.GET("/authz/jwks", handlers.GetJWKS)
	e.GET("/authz/token", handlers.GetToken)
	e.POST("/:origin/grant/:subject", handlers.GrantAuthorization)
}
