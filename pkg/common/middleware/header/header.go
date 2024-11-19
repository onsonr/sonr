package header

import "github.com/labstack/echo/v4"

func Equals(c echo.Context, key Key, value string) bool {
	return c.Response().Header().Get(key.String()) == value
}

// Exists returns true if the request has the header Key.
func Exists(c echo.Context, key Key) bool {
	return c.Response().Header().Get(key.String()) != ""
}

// Read returns the header value for the Key.
func Read(c echo.Context, key Key) string {
	return c.Response().Header().Get(key.String())
}

// Write sets the header value for the Key.
func Write(c echo.Context, key Key, value string) {
	c.Response().Header().Set(key.String(), value)
}
