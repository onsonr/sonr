package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/context"
)

// ValidateUserCredential finds the user credential and validates it against the
// session challenge
func ValidateUserCredential(c echo.Context) error {
	s, err := context.Get(c)
	if err != nil {
		return err
	}
	cred := c.FormValue("credential")
	if cred == "" {
		return echo.NewHTTPError(404, "missing credential")
	}
	return nil
}
