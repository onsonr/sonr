package session

import "github.com/labstack/echo/v4"

func IsUniqueHandle(c echo.Context, handle string) bool {
	return true
}

func IsValidFirstName(c echo.Context, firstName string) bool {
	return true
}

func IsValidLastInitial(c echo.Context, lastInitial string) bool {
	return true
}

func IsHuman(c echo.Context, sum int) bool {
	return true
}
