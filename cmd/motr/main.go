//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/onsonr/sonr/pkg/core/dwn"
	"github.com/onsonr/sonr/pkg/core/dwn/server"
)

var (
	env    *dwn.Environment
	config *dwn.Config
	srv    server.Server
	err    error
)

func main() {
	// Load dwn config
	if config, err = dwn.LoadJSONConfig(); err != nil {
		panic(err)
	}

	srv = server.New(env, config)
	srv.Serve()
}
