//go:build !wasm
// +build !wasm

package main

import (
	"context"

	"github.com/sonr-io/sonr/internal/highway"
)

func main() {
	// Start the app
	hw, err := highway.NewHighway(context.Background())
	if err != nil {
		panic(err)
	}
	hw.Serve()
	select {}
}
