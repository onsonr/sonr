package vfs

import (
	"github.com/ipfs/boxo/files"
)

var (
	kServiceWorkerFileName = "server/sw.js"
	kVaultFileName         = "server/vault.wasm"
	kIndexFileName         = "index.html"
)

func AssembleDirectory() files.Directory {
	fileMap := map[string]files.Node{
		kVaultFileName:         DWNWasmFile(),
		kServiceWorkerFileName: SWJSFile(),
	}

	return files.NewMapDirectory(fileMap)
}
