package vault

import (
	"errors"

	"github.com/labstack/echo/v4"
)

var (
	ErrFailedIPFSClient        = errors.New("failed to create IPFS client")
	ErrFailedValNodeConn       = errors.New("failed to connect to validator node")
	ErrFailedMatrixConn        = errors.New("failed to connect to matrix HomeServer")
	ErrFailedMatrixClientConn  = errors.New("failed to establish matrix client connection")
	ErrFailedServiceResolution = errors.New("failed to resolve service origin from the blockchain")
	ErrFailedWebauthnOptions   = errors.New("failed to create webauthn options")
)

type vault struct {
	echo.Context
}

func Vault(ctx echo.Context) *vault {
	return &vault{
		Context: ctx,
	}
}
