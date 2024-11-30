package jwt

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common"
	"github.com/onsonr/sonr/pkg/crypto/mpc"
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

// UCAN returns middleware to validate UCAN tokens
func UCAN(source mpc.KeyshareSource) echo.MiddlewareFunc {
	c := DefaultUCANConfig
	c.KeySource = source
	return UCANWithConfig(c)
}

// UCANWithConfig returns UCAN middleware with custom config
func UCANWithConfig(config UCANConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultUCANConfig.Skipper
	}
	if config.TokenLookup == "" {
		config.TokenLookup = DefaultUCANConfig.TokenLookup
	}
	if config.AuthScheme == "" {
		config.AuthScheme = DefaultUCANConfig.AuthScheme
	}

	// Initialize
	parts := strings.Split(config.TokenLookup, ":")
	extractor := tokenFromHeader(parts[1], config.AuthScheme)
	switch parts[0] {
	case "query":
		extractor = tokenFromQuery(parts[1])
	case "param":
		extractor = tokenFromParam(parts[1])
	case "cookie":
		extractor = tokenFromCookie(parts[1])
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper != nil && config.Skipper(c) {
				return next(c)
			}

			auth, err := extractor(c)
			if err != nil {
				return echo.NewHTTPError(401, err.Error())
			}

			parser := config.KeySource.UCANParser()
			token, err := parser.Parse(auth)
			if err != nil {
				return echo.NewHTTPError(401, "invalid UCAN token")
			}

			// Store token in context
			c.Set("ucan", token)
			return next(c)
		}
	}
}

// tokenFromHeader extracts token from header
func tokenFromHeader(header string, authScheme string) func(echo.Context) (string, error) {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		if auth == "" {
			return "", fmt.Errorf("missing auth token")
		}
		if authScheme == "" {
			return auth, nil
		}
		l := len(authScheme)
		if len(auth) > l+1 && auth[:l] == authScheme {
			return auth[l+1:], nil
		}
		return "", fmt.Errorf("invalid auth scheme")
	}
}

// tokenFromQuery extracts token from query string
func tokenFromQuery(param string) func(echo.Context) (string, error) {
	return func(c echo.Context) (string, error) {
		token := c.QueryParam(param)
		if token == "" {
			return "", fmt.Errorf("missing auth token")
		}
		return token, nil
	}
}

// tokenFromParam extracts token from url param
func tokenFromParam(param string) func(echo.Context) (string, error) {
	return func(c echo.Context) (string, error) {
		token := c.Param(param)
		if token == "" {
			return "", fmt.Errorf("missing auth token")
		}
		return token, nil
	}
}

// tokenFromCookie extracts token from cookie
func tokenFromCookie(name string) func(echo.Context) (string, error) {
	return func(c echo.Context) (string, error) {
		cookie, err := c.Cookie(name)
		if err != nil {
			return "", fmt.Errorf("missing auth token")
		}
		return cookie.Value, nil
	}
}
