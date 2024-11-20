//go:build js && wasm

package main

import (
	"github.com/onsonr/sonr/cmd/hway/server"
)

func main() {
	s := server.New()
	s.Serve()
}
