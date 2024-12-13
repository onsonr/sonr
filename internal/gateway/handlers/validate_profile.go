package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/internal/gateway/context"
	"github.com/onsonr/sonr/internal/nebula/input"
	"github.com/onsonr/sonr/pkg/common/response"
)

// ValidateProfileHandle finds the chosen handle and verifies it is unique
func ValidateProfileHandle(c echo.Context) error {
	handle := c.FormValue("handle")
	ok, err := context.HandleExists(c, handle)
	if err != nil {
		return response.TemplEcho(c, input.HandleError(handle))
	}
	if ok {
		return response.TemplEcho(c, input.HandleError(handle))
	}
	ks, err := mpc.GenEnclave()
	if err != nil {
		return response.TemplEcho(c, input.HandleError(handle))
	}
	err = context.InsertProfile(c, ks.Address(), handle, fmt.Sprintf("%s %s", c.FormValue("first_name"), c.FormValue("last_name")))
	if err != nil {
		return err
	}
	return nil
}

// ValidateProfileHandle finds the chosen handle and verifies it is unique
func ValidateIsHumanSum(c echo.Context) error {
	d, err := context.GetCreateProfileData(c)
	if err != nil {
		return err
	}
	if ok := context.VerifyIsHumanSum(c); !ok {
		return response.TemplEcho(c, input.HumanSliderError(d.FirstNumber, d.LastNumber))
	}
	return nil
}
