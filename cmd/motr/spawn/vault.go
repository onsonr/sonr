package spawn

import (
	"github.com/ipfs/boxo/files"

	"github.com/onsonr/sonr/cmd/motr/handlers/view"
	"github.com/onsonr/sonr/pkg/common/middleware/render"
	"github.com/onsonr/sonr/pkg/core/dwn"
)

const (
	FileNameConfigJSON = "dwn.json"
	FileNameIndexHTML  = "index.html"
)

// spawnVaultDirectory creates a new directory with the default files
func NewVault(cnfg *dwn.Config) (files.Node, error) {
	idxFile, err := render.TemplFileNode(view.IndexFile())
	if err != nil {
		return nil, err
	}

	cnfgBz, err := cnfg.MarshalJSON()
	if err != nil {
		return nil, err
	}

	fileMap := map[string]files.Node{
		FileNameConfigJSON: files.NewBytesFile(cnfgBz),
		FileNameIndexHTML:  idxFile,
	}
	return files.NewMapDirectory(fileMap), nil
}
