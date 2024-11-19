package static

import (
	"bytes"
	"context"

	"github.com/ipfs/boxo/files"

	"github.com/onsonr/sonr/app/nebula/views/vault"
	dwn "github.com/onsonr/sonr/pkg/motr/config"
)

// BuildVaultFile builds the index.html file for the vault
func BuildVaultFile(cnfg *dwn.Config) (files.Node, error) {
	w := bytes.NewBuffer(nil)
	err := vault.IndexFile().Render(context.Background(), w)
	if err != nil {
		return nil, err
	}
	indexData := w.Bytes()
	return files.NewBytesFile(indexData), nil
}
