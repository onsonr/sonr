//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/cmd/motr/fetch"
	"github.com/onsonr/sonr/internal/ctx"
	"github.com/onsonr/sonr/workers/routes"
)

func main() {
	e := echo.New()
	e.Use(ctx.SessionMiddleware)
	routes.RegisterClientAPI(e)
	fetch.Serve(e)
}
