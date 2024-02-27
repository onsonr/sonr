package shared

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var tempSigningKey = []byte("secret")

func JWT(c echo.Context) *jwtMdw {
	return &jwtMdw{
		Context: c,
	}
}

// UseJWTController returns a middleware that uses JWT for session management
func UseJWTController() echo.MiddlewareFunc {
	// Configure middleware with the custom claims type
	config := echojwt.Config{
		ContextKey: "controller",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtControllerClaims)
		},
		SigningKey: tempSigningKey,
	}
	return echojwt.WithConfig(config)
}

func (j *jwtMdw) GetControllerClaims() *JwtControllerClaims {
	claims := j.Get("controller").(*jwt.Token).Claims
	return claims.(*JwtControllerClaims)
}

type jwtMdw struct {
	echo.Context
}
