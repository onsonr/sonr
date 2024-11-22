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

//go:embed styles.css
var cssData string

func staticCSS() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Content-Type", "text/css")
			return c.String(200, cssData)
		}
	}
}
