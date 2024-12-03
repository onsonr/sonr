package signer

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/mpc"
)

type SignerContext struct {
	echo.Context
	signer    mpc.KeyshareSource
	keyset    mpc.Keyset
	hasSigner bool
	hasKeyset bool
}

func UseMiddleware() echo.MiddlewareFunc {
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
	return sc
}
