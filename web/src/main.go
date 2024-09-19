//go:build wasm

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/web/src/pages"
	"github.com/syumai/workers"
)

// # Sonr.ID
//
// This App is used as an IPFS gateway proxy for exissting Sonr DWN nodes
func main() {
	// TODO: Route from /ipfs/cid and /ipns/did to the gateway
	// 1. Display Generic Homepage
	e := echo.New()
	// Configure the server
	e.GET("/", pages.HomeView)
	e.GET("/allocate", pages.AllocateView)
	// 2. Present Terms Agreement and Checkbox to Accept
	// 3. Collect UserAgent, Set-Cookie, and Client Headers
	// 4. Redirect to DWN Node
	workers.Serve(e)
}
