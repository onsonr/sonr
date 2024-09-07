package vfs

import (
	"github.com/ipfs/boxo/files"
)

var (
	kServiceWorkerFileName = "sw.js"
	kVaultFileName         = "vault.wasm"
	kIndexFileName         = "index.html"
)

func AssembleDirectory() files.Directory {
	fileMap := map[string]files.Node{
		kVaultFileName:         DWNWasmFile(),
		kServiceWorkerFileName: SWJSFile(),
		kIndexFileName:         IndexHTMLFile(),
	}

	return files.NewMapDirectory(fileMap)
}
