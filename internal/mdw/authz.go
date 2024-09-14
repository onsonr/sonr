package mdw

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gopkg.in/macaroon.v2"
)

const (
	OriginMacroonCaveat  MacroonCaveat = "origin"
	ScopesMacroonCaveat  MacroonCaveat = "scopes"
	SubjectMacroonCaveat MacroonCaveat = "subject"
	ExpMacroonCaveat     MacroonCaveat = "exp"
	TokenMacroonCaveat   MacroonCaveat = "token"
)

type MacroonCaveat string

func (c MacroonCaveat) Equal(other string) bool {
	return string(c) == other
}

func (c MacroonCaveat) String() string {
	return string(c)
}

func (c MacroonCaveat) Verify(value string) error {
	switch c {
	case OriginMacroonCaveat:
		return nil
	case ScopesMacroonCaveat:
		return nil
	case SubjectMacroonCaveat:
		return nil
	case ExpMacroonCaveat:
		// Check if the expiration time is still valid
		exp, err := time.Parse(time.RFC3339, value)
		if err != nil {
			return err
		}
		if time.Now().After(exp) {
			return fmt.Errorf("expired")
		}
		return nil
	case TokenMacroonCaveat:
		return nil
	default:
		return fmt.Errorf("unknown caveat: %s", c)
	}
}

var MacroonCaveats = []MacroonCaveat{OriginMacroonCaveat, ScopesMacroonCaveat, SubjectMacroonCaveat, ExpMacroonCaveat, TokenMacroonCaveat}

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
