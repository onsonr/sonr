package session

import (
	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"

	"github.com/onsonr/sonr/internal/headers"
)

// GetSession returns the current Session
func GetSession(c echo.Context) *Session {
	return c.(*Session)
}

// UseSession establishes a Session Cookie.
func UseSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sc := initSession(c)
		headers := new(headers.RequestHeaders)
		err := sc.Bind(headers)
		if err != nil {
			return err
		}
		return next(sc)
	}
}

func initSession(c echo.Context) *Session {
	s := &Session{Context: c}
	if val := ReadCookie(c, "session"); val == "" {
		id := ksuid.New().String()
		WriteCookie(c, "session", id)
	}
	return s
}
