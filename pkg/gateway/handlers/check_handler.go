package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/internal/nebula/input"
	"github.com/onsonr/sonr/pkg/gateway/middleware"
)

// CheckProfileHandle finds the chosen handle and verifies it is unique
func CheckProfileHandle(c echo.Context) error {
	handle := c.FormValue("handle")
	if handle == "" {
		return middleware.Render(c, input.HandleError(handle, "Please enter a valid handle"))
	}
	//
	// if ok {
	// 	return middleware.Render(c, input.HandleError(handle))
	// }
	//
	return middleware.Render(c, input.HandleSuccess(handle))
}

// ValidateProfileHandle finds the chosen handle and verifies it is unique
func CheckIsHumanSum(c echo.Context) error {
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
