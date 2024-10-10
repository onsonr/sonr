//go:build js && wasm

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/ctx"
	"github.com/onsonr/sonr/pkg/workers/routes"
	"github.com/syumai/workers"
)

func main() {
	s := echo.New()
	s.Use(ctx.SessionMiddleware)
	routes.RegisterProxyViews(s)
	routes.RegisterProxyAPI(s)
	workers.Serve(s)
}
