package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/sonrhq/sonr/pkg/shared"
)

// UseTimeout returns an http.Handler middleware that sets a timeout for the request.
func UseJWT(duration time.Duration) echo.MiddlewareFunc {
	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(shared.JWTHighwayClaims)
		},
		SigningKey: []byte("secret"),
	}
	return echojwt.WithConfig(config)
}
