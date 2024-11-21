package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common/middleware/session"
	"github.com/onsonr/sonr/pkg/webapp"
)

func main() {
	// Setup
	e := echo.New()
	e.Use(session.HwayMiddleware())

	// Add WASM-specific routes
	webapp.RegisterLandingFrontend(e)

	if err := e.Start(":3000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
