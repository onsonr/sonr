package handlers

import (
	"strconv"

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
	data := context.GetCreateProfileData(c)
	value := c.FormValue("is_human")
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return response.TemplEcho(c, input.HumanSliderError(data.FirstNumber, data.LastNumber))
	}
	if intValue != data.Sum() {
		return response.TemplEcho(c, input.HumanSliderError(data.FirstNumber, data.LastNumber))
	}
	return response.TemplEcho(c, input.HumanSliderSuccess())
}
