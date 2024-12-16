package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/gateway/middleware"
	"github.com/onsonr/sonr/internal/nebula/input"
)

// ValidateProfileHandle finds the chosen handle and verifies it is unique
func ValidateProfileHandle(c echo.Context) error {
	handle := c.FormValue("handle")
	//
	// if ok {
	// 	return middleware.Render(c, input.HandleError(handle))
	// }
	//
	return middleware.Render(c, input.HandleSuccess(handle))
}

// ValidateProfileHandle finds the chosen handle and verifies it is unique
func ValidateIsHumanSum(c echo.Context) error {
	// data := context.GetCreateProfileData(c)
	// value := c.FormValue("is_human")
	// intValue, err := strconv.Atoi(value)
	// if err != nil {
	// 	return middleware.Render(c, input.HumanSliderError(data.FirstNumber, data.LastNumber))
	// }
	// if intValue != data.Sum() {
	// 	return middleware.Render(c, input.HumanSliderError(data.FirstNumber, data.LastNumber))
	// }
	return middleware.Render(c, input.HumanSliderSuccess())
}
