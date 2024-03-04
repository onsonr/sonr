package common

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
)

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
