package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/donseba/go-htmx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/segmentio/ksuid"
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

// UseSessionID sets the session id cookie as middleware function
func UseSessionID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			cookie, err := ctx.Cookie("session-id")
			if err != nil || cookie.Value == "" {
				ksid, _ := ksuid.NewRandomWithTime(time.Now())
				cookie := new(http.Cookie)
				cookie.Name = "session-id"
				cookie.Value = ksid.String()
				cookie.Expires = time.Now().Add(1 * time.Hour)
				ctx.SetCookie(cookie)
			}
			return next(ctx)
		}
	}
}

func UseHTMX(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		hxh := htmx.HxRequestHeaderFromRequest(c.Request())
		ctx = context.WithValue(ctx, HTMXHeaderKey, hxh)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
