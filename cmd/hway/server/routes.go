//go:build js && wasm
// +build js,wasm

package server

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/cmd/hway/server/handlers"
)

func RegisterGatewayAPI(e *echo.Echo) {
}

func RegisterFrontendViews(e *echo.Echo) {
	e.GET("/", handlers.RenderHome)
	e.GET("/login", handlers.RenderLogin)
	e.GET("/register", handlers.RenderRegister)
}
