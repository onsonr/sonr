//go:build js && wasm
// +build js,wasm

package vault

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/web/vault/handlers"
)

// RegisterRoutes registers the Decentralized Web Node API routes.
func RegisterRoutes(e *echo.Echo) {
	e.GET("/register/:subject/start", handlers.RegisterSubjectStart)
	e.POST("/register/:subject/check", handlers.RegisterSubjectCheck)
	e.POST("/register/:subject/finish", handlers.RegisterSubjectFinish)

	e.GET("/login/:subject/start", handlers.LoginSubjectStart)
	e.POST("/login/:subject/check", handlers.LoginSubjectCheck)
	e.POST("/login/:subject/finish", handlers.LoginSubjectFinish)

	e.GET("/:origin/grant/jwks", handlers.GetJWKS)
	e.GET("/:origin/grant/token", handlers.GetToken)
	e.POST("/:origin/grant/:subject", handlers.GrantAuthorization)
}
