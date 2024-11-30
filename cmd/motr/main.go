//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/cmd/motr/internal"
	"github.com/onsonr/sonr/pkg/common/session"
	"github.com/onsonr/sonr/web/vault"
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
	e.Use(internal.WasmContextMiddleware)
	vault.ServeStatic(e)
	vault.RegisterAPI(e)

	internal.ServeFetch(e)
}
