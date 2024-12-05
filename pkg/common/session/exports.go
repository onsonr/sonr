package session

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common"
)

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
