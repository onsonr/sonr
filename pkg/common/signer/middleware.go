package signer

import (
	"github.com/ipfs/kubo/client/rpc"
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/mpc"
)

type SignerContext struct {
	echo.Context
	ipfs *rpc.HttpApi

	signer    mpc.KeyshareSource
	keyset    mpc.Keyset
	hasIPFS   bool
	hasSigner bool
	hasKeyset bool
}

func Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := initContext(c)
			return next(cc)
		}
	}
}

func FromContext(c echo.Context) (*SignerContext, error) {
	cc, ok := c.(*SignerContext)
	if !ok {
		return nil, echo.NewHTTPError(401, "invalid signer context")
	}
	return cc, nil
}

func initContext(c echo.Context) *SignerContext {
	sc := &SignerContext{
		Context:   c,
		hasSigner: false,
		hasKeyset: false,
	}
	api, err := rpc.NewLocalApi()
	if err != nil {
		sc.hasIPFS = false
		return sc
	}
	sc.ipfs = api
	sc.hasIPFS = true
	return sc
}
