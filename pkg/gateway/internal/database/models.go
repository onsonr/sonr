package database

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrInvalidCredentials = echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	ErrInvalidSubject     = echo.NewHTTPError(http.StatusBadRequest, "Invalid subject")
	ErrInvalidUser        = echo.NewHTTPError(http.StatusBadRequest, "Invalid user")

	ErrUserAlreadyExists = echo.NewHTTPError(http.StatusConflict, "User already exists")
	ErrUserNotFound      = echo.NewHTTPError(http.StatusNotFound, "User not found")
)

type Session struct {
	ID               string `json:"id"`
	BrowserName      string `json:"browserName"`
	BrowserVersion   string `json:"browserVersion"`
	UserArchitecture string `json:"userArchitecture"`
	Platform         string `json:"platform"`
	PlatformVersion  string `json:"platformVersion"`
	DeviceModel      string `json:"deviceModel"`
	UserHandle       string `json:"userHandle"`
	FirstName        string `json:"firstName"`
	LastInitial      string `json:"lastInitial"`
	VaultAddress     string `json:"vaultAddress"`
}

type User struct {
	Address   string `json:"address"`
	Handle    string `json:"handle"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	VaultCID  string `json:"vaultCID"`
}
