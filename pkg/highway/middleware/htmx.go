package middleware

import (
	"context"

	"github.com/donseba/go-htmx"
	"github.com/labstack/echo/v4"
)

// HeaderKey is the key for the htmx request header
type HeaderKey string

// HTMXHeaderKey is the key for the htmx request header
const HTMXHeaderKey HeaderKey = "htmx-request-header"

func HTMX(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		hxh := htmx.HxRequestHeaderFromRequest(c.Request())

		ctx = context.WithValue(ctx, HTMXHeaderKey, hxh)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
