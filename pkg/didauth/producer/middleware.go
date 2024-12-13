package producer

import (
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/crypto/ucan"
	"github.com/onsonr/sonr/pkg/ipfsapi"

	"github.com/labstack/echo/v4"
)

// Middleware returns middleware to spawn controllers and validate UCAN tokens
func Middleware(ipc ipfsapi.Client, perms ucan.Permissions) echo.MiddlewareFunc {
	// Setup token store and parser
	store := ipfsapi.NewUCANStore(ipc)
	parser := ucan.NewTokenParser(perms.GetConstructor(), store, store)

	// Return middleware
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := ProducerContext{
				Context:     c,
				IPFSClient:  ipc,
				TokenParser: parser,
				TokenStore:  store,
			}
			return next(ctx)
		}
	}
}

func NewKeyset(c echo.Context) (mpc.Enclave, error) {
	ks, err := mpc.GenEnclave()
	if err != nil {
		return nil, err
	}
	return ks, nil
}
