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
	s.Use(ctx.HighwaySessionMiddleware)
	routes.RegisterGatewayViews(s)
	routes.RegisterGatewayAPI(s)
	workers.Serve(s)
}
