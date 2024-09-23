package vfs

import (
	_ "embed"

	"github.com/ipfs/boxo/files"
)

//go:embed dwn.wasm
var dwnWasmData []byte

//go:embed sw.js
var swJSData []byte

func DWNWasmFile() files.Node {
	return files.NewBytesFile(dwnWasmData)
}

// Use ServiceWorkerJS template to generate the service worker file
func SWJSFile() files.Node {
	return files.NewBytesFile(swJSData)
}
