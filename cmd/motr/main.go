//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/onsonr/sonr/web/vault/server"
	"github.com/onsonr/sonr/web/vault/types"
)

var (
	env    *types.Environment
	config *types.Config
	srv    server.Server
	err    error
)

func main() {
	// Load dwn config

	srv = server.New(env, config)
	srv.Serve()
}
