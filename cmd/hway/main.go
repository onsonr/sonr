//go:build js && wasm

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/ctx"
	"github.com/onsonr/sonr/workers/proxy"
	"github.com/syumai/workers"
)

func main() {
	s := echo.New()
	s.Use(ctx.UseSession)
	proxy.RegisterViews(s)
	workers.Serve(s)
}
