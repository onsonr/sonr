package main

import (
	_ "embed"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common/middleware/response"
	"github.com/onsonr/sonr/pkg/common/middleware/session"
	"github.com/onsonr/sonr/web/landing/pages/home"
	"github.com/onsonr/sonr/web/vault/pages/login"
	"github.com/onsonr/sonr/web/vault/pages/register"
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
	e.GET("/", response.Templ(home.Page()))
	e.GET("/register", response.Templ(register.Page()))
	e.GET("/login", response.Templ(login.Page()))

	if err := e.Start(":3000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
