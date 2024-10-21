package vaultindex

import (
	"bytes"
	"context"

	"github.com/ipfs/boxo/files"
	"github.com/onsonr/sonr/internal/dwn/gen"
)

// BuildFile builds the index.html file for the vault
func BuildFile(cnfg *gen.Config) (files.Node, error) {
	w := bytes.NewBuffer(nil)
	err := IndexFile().Render(context.Background(), w)
	if err != nil {
		return nil, err
	}
	indexData := w.Bytes()
	return files.NewBytesFile(indexData), nil
}
