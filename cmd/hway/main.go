package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common/middleware/render"
	"github.com/onsonr/sonr/pkg/common/middleware/session"
	"github.com/onsonr/sonr/pkg/webapp/pages"
)

func main() {
	// Setup
	e := echo.New()
	e.Use(session.HwayMiddleware())
	e.Use(staticCSS())

	// Add Gateway Specific Routes
	e.GET("/", render.Templ(pages.HomePage()))
	e.GET("/register", render.Templ(pages.AuthPage()))
	e.GET("/login", render.Templ(pages.AuthPage()))

	if err := e.Start(":3000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
