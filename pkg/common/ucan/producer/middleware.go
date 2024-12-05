//go:build !js && !wasm
// +build !js,!wasm

package producer

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common"
)

// UCANMiddleware returns middleware to validate UCANMiddleware tokens
func UCANMiddleware(ipfs common.IPFSClient, parser common.UCANParser) echo.MiddlewareFunc {
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
