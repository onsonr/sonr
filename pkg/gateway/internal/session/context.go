package session

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common"
)

type contextKey string

// Context keys
const (
	DataContextKey contextKey = "http_session_data"
)

type SessionCtx interface {
	ID() string
	BrowserName() string
	BrowserVersion() string
}

// Get returns the session.Context from the echo context.
func Get(c echo.Context) (SessionCtx, error) {
	ctx, ok := c.(*HTTPContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Session Context not found")
	}
	return ctx, nil
}

// TODO: Returns fixed chain ID for testing.
func GetChainID(c echo.Context) string {
	return "sonr-testnet-1"
}

// SetVaultAddress sets the address of the vault
func SetVaultAddress(c echo.Context, address string) error {
	return common.WriteCookie(c, common.SonrAddress, address)
}

// SetVaultAuthorization sets the UCAN CID of the vault
func SetVaultAuthorization(c echo.Context, ucanCID string) error {
	common.HeaderWrite(c, common.Authorization, formatAuth(ucanCID))
	return nil
}
