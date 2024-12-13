package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/mpc/spec"
)

// ControllerConfig defines the configuration for UCAN middleware
type ControllerConfig struct {
	// Skipper defines a function to skip middleware
	Skipper func(c echo.Context) bool

	// KeySource provides the source for validating UCANs
	KeySource spec.KeyshareSource

	// TokenLookup is a string in the form of "<source>:<name>" that is used
	// to extract token from the request.
	// Optional. Default value "header:Authorization".
	// Possible values:
	// - "header:<name>"
	// - "query:<name>"
	// - "param:<name>"
	// - "cookie:<name>"
	TokenLookup string

	// AuthScheme to be used in the Authorization header.
	// Optional. Default value "Bearer".
	AuthScheme string
}

// DefaultControllerConfig is the default UCAN middleware config
var DefaultControllerConfig = ControllerConfig{
	Skipper:     nil,
	TokenLookup: "header:Authorization",
	AuthScheme:  "Bearer",
}

type Option func(c *ControllerConfig)

func WithSkipper(skipper func(c echo.Context) bool) Option {
	return func(c *ControllerConfig) {
		c.Skipper = skipper
	}
}

func WithAuthScheme(scheme string) Option {
	return func(c *ControllerConfig) {
		c.AuthScheme = scheme
	}
}

// WithTokenLookup sets the token lookup strategy
func WithTokenLookup(lookup string) Option {
	return func(c *ControllerConfig) {
		c.TokenLookup = lookup
	}
}
