//go:build js && wasm
// +build js,wasm

package vault

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/vault/handlers"
)

// RegisterAPI registers the Decentralized Web Node API routes.
func RegisterAPI(e *echo.Echo) {
	e.GET("/register/:subject/start", handlers.RegisterSubjectStart)
	e.POST("/register/:subject/finish", handlers.RegisterSubjectFinish)

	e.GET("/login/:subject/start", handlers.LoginSubjectStart)
	e.POST("/login/:subject/finish", handlers.LoginSubjectFinish)

	e.GET("/authz/jwks", handlers.GetJWKS)
	e.GET("/authz/token", handlers.GetToken)
	e.POST("/:origin/grant/:subject", handlers.GrantAuthorization)
}
