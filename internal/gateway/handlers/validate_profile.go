package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/internal/gateway/context"
)

// ValidateProfileHandle finds the chosen handle and verifies it is unique
func ValidateProfileSubmit(c echo.Context) error {
	if err := context.VerifyIsHumanSum(c); err != nil {
		return err
	}
	handle := c.FormValue("handle")
	ok, err := context.HandleExists(c, handle)
	if err != nil {
		return err
	}
	if ok {
		return echo.NewHTTPError(400, "handle already exists")
	}
	ks, err := mpc.GenEnclave()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	err = context.RefreshChallenge(c)
	if err != nil {
		return err
	}
	err = context.InsertProfile(c, ks.Address(), handle, fmt.Sprintf("%s %s", c.FormValue("first_name"), c.FormValue("last_name")))
	if err != nil {
		return err
	}
	return nil
}
