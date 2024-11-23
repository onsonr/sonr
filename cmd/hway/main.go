package main

import (
	_ "embed"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common/middleware/response"
	"github.com/onsonr/sonr/pkg/common/middleware/session"
	"github.com/onsonr/sonr/pkg/webapp/pages"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func main() {
	// Setup
	e := echo.New()
	e.Use(session.HwayMiddleware())

	// Add Gateway Specific Routes
	e.GET("/", response.Templ(pages.HomePage()))
	e.GET("/register", response.Templ(pages.AuthPage()))
	e.GET("/login", response.Templ(pages.AuthPage()))

	if err := e.Start(":3000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
