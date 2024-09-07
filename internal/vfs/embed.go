package vfs

import (
	_ "embed"

	"github.com/ipfs/boxo/files"
)

//go:embed dwn.wasm
var dwnWasmData []byte

func DWNWasmFile() files.Node {
	return files.NewBytesFile(dwnWasmData)
}

//go:embed sw.js
var swJSData []byte

func SWJSFile() files.Node {
	return files.NewBytesFile(swJSData)
}

//go:embed index.html
var indexHTMLData []byte

func IndexHTMLFile() files.Node {
	return files.NewBytesFile(indexHTMLData)
}
