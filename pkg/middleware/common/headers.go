package common

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
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

// Partial renders a templ.Component
func Partial(cmp templ.Component) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Writer.WriteHeader(http.StatusOK)
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
		return cmp.Render(c.Request().Context(), c.Response())
	}
}

// Render renders a templ.Component
func Render(c echo.Context, cmp templ.Component) error {
	c.Response().Writer.WriteHeader(http.StatusOK)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(c.Request().Context(), c.Response())
}
