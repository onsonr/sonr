//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common/middleware/session"
	"github.com/onsonr/sonr/web/vault"
	"github.com/onsonr/sonr/web/vault/bridge"
	"github.com/onsonr/sonr/web/vault/types"
)

var (
	env    *types.Environment
	config *types.Config
	err    error
)

func main() {
	// Load dwn config
	e := echo.New()
	// if config, err = dwn.LoadJSONConfig(); err != nil {
	// 	panic(err)
	// }

	e.Use(session.MotrMiddleware(config))
	e.Use(bridge.WasmContextMiddleware)
	vault.RegisterRoutes(e)
	bridge.ServeFetch(e)
}
