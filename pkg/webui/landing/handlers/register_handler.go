package handlers

import "github.com/labstack/echo/v4"

func HandleRegister(c echo.Context) error {
	return nil
}

func HandleRegisterSubmit(c echo.Context) error {
	return nil
}

// ╭────────────────────────────────────────────────────────╮
// │                  	Utility Functions 	                │
// ╰────────────────────────────────────────────────────────╯

func hasProfile(c echo.Context) bool {
	return false
}

func hasPasscode(c echo.Context) bool {
	return false
}

func hasCredentials(c echo.Context) bool {
	return false
}
