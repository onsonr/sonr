package context

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/medama-io/go-useragent"
	config "github.com/onsonr/sonr/pkg/config/hway"
	"github.com/onsonr/sonr/pkg/database/sessions"
	"gorm.io/gorm"
)

// Middleware creates a new session middleware
func Middleware(db *gorm.DB, env config.Hway) echo.MiddlewareFunc {
	ua := useragent.NewParser()
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			agent := ua.Parse(c.Request().UserAgent())
			cc := NewHTTPContext(c, db, agent)
			if err := cc.initSession(); err != nil {
				return err
			}
			return next(cc)
		}
	}
}

// HTTPContext is the context for HTTP endpoints.
type HTTPContext struct {
	echo.Context
	db   *gorm.DB
	sess *sessions.Session
	user *sessions.User
	env  config.Hway
	useragent.UserAgent
}

// Get returns the HTTPContext from the echo context
func Get(c echo.Context) (*HTTPContext, error) {
	ctx, ok := c.(*HTTPContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Session Context not found")
	}
	return ctx, nil
}

// NewHTTPContext creates a new session context
func NewHTTPContext(c echo.Context, db *gorm.DB, a useragent.UserAgent) *HTTPContext {
	return &HTTPContext{
		Context: c,
		db:      db,
	}
}

// Session returns the current session
func (s *HTTPContext) Session() *sessions.Session {
	return s.sess
}
