//go:build !js && !wasm
// +build !js,!wasm

package producer

import (
	"github.com/onsonr/sonr/crypto/didkey"
	"github.com/onsonr/sonr/pkg/common/ipfs"

	"github.com/labstack/echo/v4"
)

// UCANMiddleware returns middleware to validate UCANMiddleware tokens
func UCANMiddleware(ipfs ipfs.Client, parser *didkey.TokenParser) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := ProducerContext{
				Context:     c,
				IPFSClient:  ipfs,
				TokenParser: parser,
			}
			return next(ctx)
		}
	}
}
