//go:build js && wasm

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/ctx"
	"github.com/onsonr/sonr/nebula/components/auth"
	"github.com/onsonr/sonr/nebula/components/home"
	"github.com/syumai/workers"
)

func main() {
	s := echo.New()
	s.Use(ctx.UseSession)
	s.GET("/", home.Route)
	s.GET("/login", auth.LoginRoute)
	s.GET("/register", auth.RegisterRoute)
	workers.Serve(s)
}
