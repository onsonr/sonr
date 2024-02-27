package shared

import (
	"fmt"

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
func UseJWTController(address, handle, origin string) echo.MiddlewareFunc {
	// Configure middleware with the custom claims type
	config := echojwt.Config{
		ContextKey: "controller",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &JwtControllerClaims{
				Handle:  Cookies(c).GetHandle(),
				Address: Cookies(c).GetAddress(),
				Origin:  Cookies(c).GetOrigin(),
			}
		},
		SigningKey: tempSigningKey,
	}
	return echojwt.WithConfig(config)
}

func (j *jwtMdw) GetControllerClaims() (*JwtControllerClaims, error) {
	claims := j.Get("controller").(*jwt.Token).Claims
	if c, ok := claims.(*JwtControllerClaims); ok {
		return c, nil
	}
	return nil, fmt.Errorf("invalid claims")
}

type jwtMdw struct {
	echo.Context
}
