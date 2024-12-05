package common

import (
	"encoding/base64"
	"net/http"

	"github.com/ipfs/kubo/client/rpc"
	"github.com/labstack/echo/v4"
)

var (
	ErrInvalidCredentials = echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	ErrInvalidSubject     = echo.NewHTTPError(http.StatusBadRequest, "Invalid subject")
	ErrInvalidUser        = echo.NewHTTPError(http.StatusBadRequest, "Invalid user")

	ErrUserAlreadyExists = echo.NewHTTPError(http.StatusConflict, "User already exists")
	ErrUserNotFound      = echo.NewHTTPError(http.StatusNotFound, "User not found")
)

type SessionCtx interface {
	ID() string
	Challenge() string
	BrowserName() string
	BrowserVersion() string
}

type LargeBlob struct {
	Support string `json:"support"`
	Write   string `json:"write"`
}

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

func Base64Encode(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

func Base64Decode(data string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(data)
}

// IPFSClient is an alias for the IPFS HTTP API
type IPFSClient = *rpc.HttpApi
