//go:build js && wasm
// +build js,wasm

package controller

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/ucan/spec"
)

// Middleware returns middleware to validate Middleware tokens
func Middleware(source spec.KeyshareSource, opts ...Option) echo.MiddlewareFunc {
	c := DefaultControllerConfig
	for _, opt := range opts {
		opt(&c)
	}
	c.KeySource = source
	return initWithConfig(c)
}

// initWithConfig returns UCAN middleware with custom config
func initWithConfig(config ControllerConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultControllerConfig.Skipper
	}
	if config.TokenLookup == "" {
		config.TokenLookup = DefaultControllerConfig.TokenLookup
	}
	if config.AuthScheme == "" {
		config.AuthScheme = DefaultControllerConfig.AuthScheme
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
			token, err := parser.ParseAndVerify(c.Request().Context(), auth)
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
