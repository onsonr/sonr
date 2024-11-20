package dwn

import (
	"github.com/ipfs/boxo/files"

	"github.com/onsonr/sonr/cmd/motr/handlers/view"
	"github.com/onsonr/sonr/pkg/common/middleware/render"
)

const (
	FileNameConfigJSON = "dwn.json"
	FileNameIndexHTML  = "index.html"
)

// spawnVaultDirectory creates a new directory with the default files
func SpawnVault(cnfg *Config) (files.Directory, error) {
	idxf, err := render.TemplFileNode(view.IndexFile())
	if err != nil {
		return nil, err
	}

	cnf, err := cnfg.MarshalFileNode()
	if err != nil {
		return nil, err
	}

	return files.NewMapDirectory(map[string]files.Node{
		FileNameConfigJSON: cnf,
		FileNameIndexHTML:  idxf,
	}), nil
}
