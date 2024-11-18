package motr

import (
	_ "embed"
	"encoding/json"

	"github.com/ipfs/boxo/files"

	"github.com/onsonr/sonr/pkg/motr/config"
	"github.com/onsonr/sonr/pkg/motr/static"
)

const (
	FileNameConfigJSON = "dwn.pkl"
	FileNameIndexHTML  = "index.html"
	FileNameWorkerJS   = "sw.js"
)

//go:embed static/sw.js
var swJSData []byte

// NewVaultDirectory creates a new directory with the default files
func NewVaultDirectory(cnfg *config.Config) (files.Node, error) {
	idxFile, err := static.BuildVaultFile(cnfg)
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
