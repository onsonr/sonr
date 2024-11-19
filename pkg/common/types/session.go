package commonv1

import (
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/motr/config"
)

var (
	ErrInvalidCredentials = echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	ErrInvalidSubject     = echo.NewHTTPError(http.StatusBadRequest, "Invalid subject")
	ErrInvalidUser        = echo.NewHTTPError(http.StatusBadRequest, "Invalid user")

	ErrUserAlreadyExists = echo.NewHTTPError(http.StatusConflict, "User already exists")
	ErrUserNotFound      = echo.NewHTTPError(http.StatusNotFound, "User not found")
)

type (
	CredDescriptor  = protocol.CredentialDescriptor
	LoginOptions    = protocol.PublicKeyCredentialRequestOptions
	RegisterOptions = protocol.PublicKeyCredentialCreationOptions
	VaultSchema     = config.Schema
)

type BrowserName string

const (
	BrowserNameUnknown  BrowserName = " Not A;Brand"
	BrowserNameChromium BrowserName = "Chromium"
)

func (n BrowserName) String() string {
	return string(n)
}

type PeerRole string

const (
	RoleUnknown PeerRole = "none"
	RoleHway    PeerRole = "hway"
	RoleMotr    PeerRole = "motr"
)

func (r PeerRole) Is(role PeerRole) bool {
	return r == role
}

func (r PeerRole) String() string {
	return string(r)
}

type ClientConfig struct {
	ChainID    string `json:"chainID"`
	IPFSHost   string `json:"ipfsHost"`
	SonrAPIURL string `json:"sonrAPIURL"`
	SonrRPCURL string `json:"sonrRPCURL"`
	SonrWSURL  string `json:"sonrWSURL"`
}

type PeerInfo struct {
	ID        string                    `json:"id"`
	Challenge protocol.URLEncodedBase64 `json:"challenge"`
}

type BrowserInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type UserAgent struct {
	Architecture    string       `json:"architecture"`
	Bitness         string       `json:"bitness"`
	Browser         *BrowserInfo `json:"browser"`
	Model           string       `json:"model"`
	PlatformName    string       `json:"platformName"`
	PlatformVersion string       `json:"platformVersion"`
	IsMobile        bool         `json:"isMobile"`
}

type VaultDetails struct {
	Schema  *VaultSchema `json:"schema"`
	Address string       `json:"address"`
}
