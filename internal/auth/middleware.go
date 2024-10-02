package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/macaroon.v2"
)

func MacaroonMiddleware(secretKeyStr string, location string) echo.MiddlewareFunc {
	secretKey := []byte(secretKeyStr)
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
				for _, c := range MacroonCaveats {
					if c.String() == caveat {
						return nil
					}
				}
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
