package handlers

import "github.com/labstack/echo/v4"

type PublishRequest struct {
	Name string `json:"name"`
}

func PublishVault(c echo.Context) error {
	return nil
}
