package embed

import (
	_ "embed"
	"encoding/json"

	"github.com/ipfs/boxo/files"

	dwn "github.com/onsonr/sonr/pkg/motr/config"
	"github.com/onsonr/sonr/pkg/nebula"
)

const (
	FileNameConfigJSON = "dwn.json"
	FileNameIndexHTML  = "index.html"
	FileNameWorkerJS   = "sw.js"
)

//go:embed sw.js
var swJSData []byte

// NewVaultDirectory creates a new directory with the default files
func NewVaultDirectory(cnfg *dwn.Config) (files.Node, error) {
	idxFile, err := nebula.BuildVaultFile(cnfg)
	if err != nil {
		return nil, err
	}
	cnfgBz, err := json.Marshal(cnfg)
	if err != nil {
		return nil, err
	}
	fileMap := map[string]files.Node{
		FileNameConfigJSON: files.NewBytesFile(cnfgBz),
		FileNameIndexHTML:  idxFile,
		FileNameWorkerJS:   files.NewBytesFile(swJSData),
	}
	return files.NewMapDirectory(fileMap), nil
}
