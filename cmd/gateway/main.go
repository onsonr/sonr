//go:build js && wasm

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/syumai/workers"

	"github.com/onsonr/sonr/pkg/common/ctx"
	"github.com/onsonr/sonr/pkg/hway/routes"
)

func main() {
	s := echo.New()
	s.Use(ctx.HighwaySessionMiddleware)
	routes.RegisterGatewayViews(s)
	routes.RegisterGatewayViews(s)
	workers.Serve(s)
}
