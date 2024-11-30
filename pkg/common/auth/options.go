package auth

import (
	"github.com/labstack/echo/v4"
)

type Option func(c *UCANConfig)

func WithSkipper(skipper func(c echo.Context) bool) Option {
	return func(c *UCANConfig) {
		c.Skipper = skipper
	}
}

func WithAuthScheme(scheme string) Option {
	return func(c *UCANConfig) {
		c.AuthScheme = scheme
	}
}

// WithTokenLookup sets the token lookup strategy
func WithTokenLookup(lookup string) Option {
	return func(c *UCANConfig) {
		c.TokenLookup = lookup
	}
}
