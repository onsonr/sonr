package middleware

import (
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func InitServer(grpcAddr string) *echo.Echo {
	grpcEndpoint = grpcAddr
	e := echo.New()
	// Override default behaviors
	e.IPExtractor = echo.ExtractIPDirect()
	e.HTTPErrorHandler = redirectOnError("http://localhost:3000")

	// Built-in middleware
	e.Use(echoprometheus.NewMiddleware("hway"))
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	return e
}

func redirectOnError(target string) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			// Log the error if needed
			c.Logger().Errorf("Error: %v", he.Message)
		}
		// Redirect to main site
		c.Redirect(http.StatusFound, target)
	}
}
