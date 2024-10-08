//go:build js && wasm

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/ctx"
	"github.com/onsonr/sonr/pkg/nebula/routes"
	"github.com/syumai/workers"
)

func main() {
	s := echo.New()
	s.Use(ctx.UseSession)

	s.GET("/", routes.Home)
	s.GET("/login", routes.LoginStart)
	s.GET("/register", routes.RegisterStart)
	workers.Serve(s)
}
