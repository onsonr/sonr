package handlers

import "github.com/labstack/echo/v4"

type PinRequest struct {
	Name string `json:"name"`
}

func PinVault(c echo.Context) error {
	return nil
}
