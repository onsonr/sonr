package producer

import (
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/crypto/ucan"
	"github.com/onsonr/sonr/crypto/ucan/store"
	"github.com/onsonr/sonr/pkg/common/ipfs"

	"github.com/labstack/echo/v4"
)

// Middleware returns middleware to spawn controllers and validate UCAN tokens
func Middleware(ipfs ipfs.Client, perms ucan.Permissions) echo.MiddlewareFunc {
	// Setup token store and parser
	store := store.NewIPFSTokenStore(ipfs)
	parser := ucan.NewTokenParser(perms.GetConstructor(), store, store)

	// Return middleware
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := ProducerContext{
				Context:     c,
				IPFSClient:  ipfs,
				TokenParser: parser,
				TokenStore:  store,
			}
			return next(ctx)
		}
	}
}

func NewKeyset(c echo.Context) (mpc.Keyset, error) {
	ks, err := mpc.NewKeyset()
	if err != nil {
		return nil, err
	}
	return ks, nil
}
