package mdw

import (
	"net/http"
	"time"

	"github.com/donseba/go-htmx"
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/db"
	"github.com/segmentio/ksuid"
)

// UseSession establishes a Session Cookie.
func UseSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sc := newSession(c)
		bindRequestHeaders(sc)
		return next(sc)
	}
}

// GetSession returns the current Session
func GetSession(c echo.Context) *Session {
	return c.(*Session)
}

type Session struct {
	echo.Context
	htmx *htmx.HTMX
	dB   *db.DB
}

func (c *Session) ID() string {
	return readCookie(c, "session")
}

func (c *Session) Htmx() *htmx.HTMX {
	return c.htmx
}

func (c *Session) DB() *db.DB {
	return c.dB
}

func newSession(c echo.Context) *Session {
	s := &Session{Context: c}
	if val := readCookie(c, "session"); val == "" {
		id := ksuid.New().String()
		writeCookie(c, "session", id)
	}
	return s
}

func readCookie(c echo.Context, key string) string {
	cookie, err := c.Cookie(key)
	if err != nil {
		return ""
	}
	if cookie == nil {
		return ""
	}
	return cookie.Value
}

func writeCookie(c echo.Context, key string, value string) {
	cookie := new(http.Cookie)
	cookie.Name = key
	cookie.Value = value
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
}
