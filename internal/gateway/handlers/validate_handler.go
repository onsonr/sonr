package handlers

import (
	"github.com/labstack/echo/v4"
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

	return response.TemplEcho(c, input.HandleSuccess(handle))
}

// ValidateProfileHandle finds the chosen handle and verifies it is unique
func ValidateIsHumanSum(c echo.Context) error {
	if ok := context.VerifyIsHumanSum(c); !ok {
		return response.TemplEcho(c, input.HumanSliderError(context.GetCreateProfileData(c)))
	}
	return nil
}
