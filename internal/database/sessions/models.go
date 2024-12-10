package sessions

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	ErrInvalidCredentials = echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	ErrInvalidSubject     = echo.NewHTTPError(http.StatusBadRequest, "Invalid subject")
	ErrInvalidUser        = echo.NewHTTPError(http.StatusBadRequest, "Invalid user")

	ErrUserAlreadyExists = echo.NewHTTPError(http.StatusConflict, "User already exists")
	ErrUserNotFound      = echo.NewHTTPError(http.StatusNotFound, "User not found")
)

// Define the credential structure matching our frontend data
type Credential struct {
	ID                      string                 `json:"id"`
	RawID                   string                 `json:"rawId"`
	Type                    string                 `json:"type"`
	AuthenticatorAttachment string                 `json:"authenticatorAttachment"`
	Transports              []string               `json:"transports"`
	ClientExtensionResults  map[string]interface{} `json:"clientExtensionResults"`
	Response                struct {
		AttestationObject string `json:"attestationObject"`
		ClientDataJSON    string `json:"clientDataJSON"`
	} `json:"response"`
}

type User struct {
	gorm.Model
	Address     string        `json:"address"`
	Handle      string        `json:"handle"`
	Name        string        `json:"name"`
	CID         string        `json:"cid"`
	Credentials []*Credential `json:"credentials"`
}

type Session struct {
	gorm.Model
	ID               string `json:"id" gorm:"primaryKey"`
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
	HumanSum         int    `json:"humanSum"`
	Challenge        string `json:"challenge"`
}
