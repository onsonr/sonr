package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/sonrhq/sonr/pkg/shared"
)

// UseSessionJWT returns a middleware that uses JWT for session management
func UseSessionJWT() echo.MiddlewareFunc {
	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(shared.JwtControllerClaims)
		},
		SigningKey: []byte("secret"),
	}
	return echojwt.WithConfig(config)
}
