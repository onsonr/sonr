package dwn

import (
	_ "embed"
	"encoding/json"

	"github.com/ipfs/boxo/files"
	"github.com/onsonr/sonr/internal/dwn/gen"
	"github.com/onsonr/sonr/pkg/nebula/components/vaultindex"
)

//go:embed app.wasm
var dwnWasmData []byte

//go:embed sw.js
var swJSData []byte

var (
	dwnWasmFile = files.NewBytesFile(dwnWasmData)
	swJSFile    = files.NewBytesFile(swJSData)
)

const (
	kConfigJSONFileName    = "config.json"
	kServiceWorkerFileName = "sw.js"
	kAppWasmFileName       = "app.wasm"
	kIndexFileName         = "index.html"
)

// NewVaultDirectory creates a new directory with the default files
func NewVaultDirectory(cnfg *gen.Config) (files.Node, error) {
	idxFile, err := vaultindex.BuildFile(cnfg)
	if err != nil {
		return nil, err
	}
	cnfgFile, err := createJSONConfig(cnfg)
	if err != nil {
		return nil, err
	}
	fileMap := map[string]files.Node{
		kServiceWorkerFileName: swJSFile,
		kAppWasmFileName:       dwnWasmFile,
		kIndexFileName:         idxFile,
		kConfigJSONFileName:    cnfgFile,
	}
	return files.NewMapDirectory(fileMap), nil
}

func createJSONConfig(cnfg *gen.Config) (files.Node, error) {
	bz, err := json.Marshal(cnfg)
	if err != nil {
		return nil, err
	}
	return files.NewBytesFile(bz), nil
}
