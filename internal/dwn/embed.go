package dwn

import (
	_ "embed"

	"github.com/ipfs/boxo/files"
	"github.com/onsonr/sonr/internal/dwn/gen"
	"github.com/onsonr/sonr/nebula/components/index"
)

//go:embed app.wasm
var dwnWasmData []byte

//go:embed sw.js
var swJSData []byte

var (
	dwnWasmFile = files.NewBytesFile(dwnWasmData)
	swJSFile    = files.NewBytesFile(swJSData)
)

// NewVaultDirectory creates a new directory with the default files
func NewVaultDirectory(cnfg *gen.Config) (files.Node, error) {
	idxFile, err := index.BuildFile(cnfg)
	if err != nil {
		return nil, err
	}
	fileMap := map[string]files.Node{
		"sw.js":      swJSFile,
		"app.wasm":   dwnWasmFile,
		"index.html": idxFile,
	}
	return files.NewMapDirectory(fileMap), nil
}
