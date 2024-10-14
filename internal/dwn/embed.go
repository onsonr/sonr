package dwn

import (
	_ "embed"
	"encoding/json"

	"github.com/ipfs/boxo/files"
	"github.com/onsonr/sonr/internal/dwn/gen"
	"github.com/onsonr/sonr/pkg/nebula/components/vaultindex"
)

const (
	FileNameAppWASM    = "app.wasm"
	FileNameConfigJSON = "dwn.json"
	FileNameIndexHTML  = "index.html"
	FileNameWorkerJS   = "sw.js"
)

//go:embed app.wasm
var dwnWasmData []byte

//go:embed sw.js
var swJSData []byte

// NewVaultDirectory creates a new directory with the default files
func NewVaultDirectory(cnfg *gen.Config) (files.Node, error) {
	idxFile, err := vaultindex.BuildFile(cnfg)
	if err != nil {
		return nil, err
	}
	cnfgBz, err := json.Marshal(cnfg)
	if err != nil {
		return nil, err
	}
	fileMap := map[string]files.Node{
		FileNameAppWASM:    files.NewBytesFile(dwnWasmData),
		FileNameConfigJSON: files.NewBytesFile(cnfgBz),
		FileNameIndexHTML:  idxFile,
		FileNameWorkerJS:   files.NewBytesFile(swJSData),
	}
	return files.NewMapDirectory(fileMap), nil
}
