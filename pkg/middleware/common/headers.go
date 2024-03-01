package common

import (
	"context"

	"github.com/donseba/go-htmx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// HeaderKey is the key for the htmx request header
type HeaderKey string

// HTMXHeaderKey is the key for the htmx request header
const HTMXHeaderKey HeaderKey = "htmx-request-header"

// UseDefaults adds chi provided middleware libraries to the router.
func UseDefaults(e *echo.Echo) {
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
}

// UseHTMX sets the htmx request header as context value
func UseHTMX(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		hxh := htmx.HxRequestHeaderFromRequest(c.Request())
		ctx = context.WithValue(ctx, HTMXHeaderKey, hxh)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

// Use HyperView sets the htmx request header as context value
func UseHyperView(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		hxh := htmx.HxRequestHeaderFromRequest(c.Request())
		ctx = context.WithValue(ctx, HTMXHeaderKey, hxh)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
