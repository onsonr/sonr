package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/mpc"
)

// UCANConfig defines the configuration for UCAN middleware
type UCANConfig struct {
	// Skipper defines a function to skip middleware
	Skipper func(c echo.Context) bool

	// KeySource provides the source for validating UCANs
	KeySource mpc.KeyshareSource

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

// DefaultUCANConfig is the default UCAN middleware config
var DefaultUCANConfig = UCANConfig{
	Skipper:     nil,
	TokenLookup: "header:Authorization",
	AuthScheme:  "Bearer",
}
