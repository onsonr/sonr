//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/dwn/front"
	"github.com/onsonr/sonr/internal/dwn/handlers"
	"github.com/onsonr/sonr/internal/dwn/middleware"
	"github.com/onsonr/sonr/internal/vfs"
)

func main() {
	e := echo.New()
	e.Use(middleware.UseSession)
	front.RegisterViews(e)
	handlers.RegisterState(e)
	vfs.Serve(e)
}
