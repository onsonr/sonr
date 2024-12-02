package session

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common/cookie"
	"github.com/onsonr/sonr/pkg/common/header"
	"github.com/onsonr/sonr/pkg/vault/types"
)

// TODO: Returns fixed chain ID for testing.
func GetChainID(c echo.Context) string {
	return "sonr-testnet-1"
}

// GetVaultConfig returns the default vault config
func GetVaultConfig(c echo.Context, addr string, ucanCID string) *types.Config {
	return &types.Config{
		MotrToken:      ucanCID,
		MotrAddress:    addr,
		IpfsGatewayUrl: "http://localhost:80",
		SonrApiUrl:     "http://localhost:1317",
		SonrRpcUrl:     "http://localhost:26657",
		SonrChainId:    GetChainID(c),
		VaultSchema:    GetVaultSchema(c),
	}
}

// GetVaultSchema returns the default vault schema
func GetVaultSchema(c echo.Context) *types.Schema {
	return types.DefaultSchema()
}

// SetVaultAddress sets the address of the vault
func SetVaultAddress(c echo.Context, address string) error {
	return cookie.Write(c, cookie.SonrAddress, address)
}

// SetVaultAuthorization sets the UCAN CID of the vault
func SetVaultAuthorization(c echo.Context, ucanCID string) error {
	header.Write(c, header.Authorization, formatAuth(ucanCID))
	return nil
}
