// +build wasm

package main

import (
	"github.com/sonr-io/sonr/pkg/crypto"
)

func main() {
	_, err := crypto.GenerateWallet()
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
