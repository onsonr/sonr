package storage

import (
	"github.com/labstack/echo/v4"
)

// SetLocal sets a value in localStorage
func SetLocal(c echo.Context, key string, value any) error {
	return SetLocalStorageItem(key, value).Render(c.Request().Context(), c.Response().Writer)
}

// SetSession sets a value in sessionStorage
func SetSession(c echo.Context, key string, value any) error {
	return SetSessionStorageItem(key, value).Render(c.Request().Context(), c.Response().Writer)
}
