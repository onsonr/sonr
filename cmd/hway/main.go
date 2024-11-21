//go:build js && wasm

package main

import (
	"github.com/onsonr/sonr/pkg/gateway"
)

func main() {
	s := gateway.New()
	s.Serve()
}
