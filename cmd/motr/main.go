//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
	"syscall/js"

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

func processConfig(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return nil
	}

	configString := args[0].String()
	if err := json.Unmarshal([]byte(configString), &config); err != nil {
		println("Error parsing config:", err.Error())
		return nil
	}
	return nil
}

func main() {
	// Load dwn config
	js.Global().Set("processConfig", js.FuncOf(processConfig))

	e := echo.New()
	e.Use(session.MotrMiddleware(config))
	e.Use(internal.WasmContextMiddleware)
	vault.RegisterAPI(e)

	internal.ServeFetch(e)
}
