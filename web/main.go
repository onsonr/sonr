//go:build wasm

package main

import (
	"github.com/syumai/workers"
)

// # Sonr.ID
//
// This App is used as an IPFS gateway proxy for exissting Sonr DWN nodes
func main() {
	// TODO: Route from /ipfs/cid and /ipns/did to the gateway
	// Configure the server
	workers.Serve(nil)
}
