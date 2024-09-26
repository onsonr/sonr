package middleware

import (
	"net/http"
	"time"

	"github.com/donseba/go-htmx"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
)

type Session struct {
	echo.Context
	htmx *htmx.HTMX
}

func (c *Session) Htmx() *htmx.HTMX {
	return c.htmx
}

func (c *Session) ID() string {
	return ReadCookie(c, "session")
}

func initSession(c echo.Context) *Session {
	s := &Session{Context: c}
	if val := ReadCookie(c, "session"); val == "" {
		id := ksuid.New().String()
		WriteCookie(c, "session", id)
	}
	return s
}

func ReadCookie(c echo.Context, key string) string {
	cookie, err := c.Cookie(key)
	if err != nil {
		return ""
	}
	if cookie == nil {
		return ""
	}
	return cookie.Value
}

func WriteCookie(c echo.Context, key string, value string) {
	cookie := new(http.Cookie)
	cookie.Name = key
	cookie.Value = value
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
}
