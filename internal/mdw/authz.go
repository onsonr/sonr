package mdw

import (
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gopkg.in/macaroon.v2"
)

type Authz struct {
	echo.Context
	echojwt.Config

	signKey []byte
}

func newAuthz(c echo.Context, signKey []byte) *Authz {
	return &Authz{Context: c, signKey: signKey}
}

func (a *Authz) Accessible(route string, handler echo.HandlerFunc) echo.HandlerFunc {
	// Verify the macaroon
	//	verified := a.Verify(a.signKey, func(caveat string) error {
	// Implement your caveat verification logic here
	// For example, you might check if the caveat is still valid (e.g., not expired)
	// return nil // Return nil if the caveat is valid
	//	}, nil)
	// if !verified {
	// return func(c echo.Context) error {
	// return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid macaroon"})
	// }
	// }
	a.SetPath(route)
	return handler
}

func ValidateMacaroonMiddleware(secretKey []byte, location string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract the macaroon from the Authorization header
			auth := c.Request().Header.Get("Authorization")
			if auth == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing Authorization header"})
			}

			// Decode the macaroon
			mac, err := macaroon.Base64Decode([]byte(auth))
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid macaroon encoding"})
			}

			token, err := macaroon.New(secretKey, mac, location, macaroon.LatestVersion)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid macaroon"})
			}

			// Verify the macaroon
			err = token.Verify(secretKey, func(caveat string) error {
				// Implement your caveat verification logic here
				// For example, you might check if the caveat is still valid (e.g., not expired)
				return nil // Return nil if the caveat is valid
			}, nil)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid macaroon"})
			}

			// Macaroon is valid, proceed to the next handler
			return next(c)
		}
	}
}
