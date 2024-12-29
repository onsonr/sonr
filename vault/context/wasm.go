//go:build js && wasm
// +build js,wasm

package context

import (
	"encoding/base64"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

// AI! Fix any lint errors
func WASMMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract WASM context from headers
		if wasmCtx := c.Request().Header.Get("X-Wasm-Context"); wasmCtx != "" {
			if ctx, err := DecodeWasmContext(wasmCtx); err == nil {
				c.Set("wasm_context", ctx)
			}
		}
		return next(c)
	}
}

// decodeWasmContext decodes the WASM context from a base64 encoded string
func DecodeWasmContext(ctx string) (map[string]any, error) {
	decoded, err := base64.StdEncoding.DecodeString(ctx)
	if err != nil {
		return nil, err
	}
	var ctxData map[string]any
	err = json.Unmarshal(decoded, &ctxData)
	return ctxData, err
}
