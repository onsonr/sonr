package session

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/app/gateway/config"
	"github.com/onsonr/sonr/app/gateway/internal/database"
	"gorm.io/gorm"
)

// Middleware creates a new session middleware
func Middleware(db *gorm.DB, env config.Env) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := NewHTTPContext(c, db)
			if err := cc.InitSession(); err != nil {
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
	sess *database.Session
	env  config.Env
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
func NewHTTPContext(c echo.Context, db *gorm.DB) *HTTPContext {
	return &HTTPContext{
		Context: c,
		db:      db,
	}
}

// Session returns the current session
func (s *HTTPContext) Session() *database.Session {
	return s.sess
}
