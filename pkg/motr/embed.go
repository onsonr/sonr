package motr

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"

	"github.com/ipfs/boxo/files"

	"github.com/onsonr/sonr/app/nebula/views/vault"
	"github.com/onsonr/sonr/pkg/core/dwn"
)

const (
	FileNameConfigJSON = "dwn.json"
	FileNameIndexHTML  = "index.html"
)

// NewVaultDirectory creates a new directory with the default files
func NewVaultDirectory(cnfg *dwn.Config) (files.Node, error) {
	idxFile, err := buildVaultFile(cnfg)
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
	}
	return files.NewMapDirectory(fileMap), nil
}

// buildVaultFile builds the index.html file for the vault
func buildVaultFile(cnfg *dwn.Config) (files.Node, error) {
	w := bytes.NewBuffer(nil)
	err := vault.IndexFile().Render(context.Background(), w)
	if err != nil {
		return nil, err
	}
	indexData := w.Bytes()
	return files.NewBytesFile(indexData), nil
}
