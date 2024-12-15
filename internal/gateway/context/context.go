package context

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/medama-io/go-useragent"
	config "github.com/onsonr/sonr/pkg/config/hway"
)

// Middleware creates a new session middleware
func Middleware(env config.Hway) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ua := useragent.NewParser()
			agent := ua.Parse(c.Request().UserAgent())
			cc := &HTTPContext{Context: c, env: env, agent: agent}
			return next(cc)
		}
	}
}

// HTTPContext is the context for HTTP endpoints.
type HTTPContext struct {
	echo.Context
	id    string
	env   config.Hway
	agent useragent.UserAgent
}

// Get returns the HTTPContext from the echo context
func Get(c echo.Context) (*HTTPContext, error) {
	ctx, ok := c.(*HTTPContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Session Context not found")
	}
	return ctx, nil
}

// ForbiddenDevice returns true if the device is unavailable
func ForbiddenDevice(c echo.Context) bool {
	s, err := Get(c)
	if err != nil {
		return true
	}
	return s.agent.IsBot() || s.agent.IsTV()
}
