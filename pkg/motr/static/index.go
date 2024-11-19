package static

import (
	"bytes"
	"context"

	"github.com/ipfs/boxo/files"

	dwn "github.com/onsonr/sonr/pkg/motr/config"
	"github.com/onsonr/sonr/pkg/nebula/views"
)

// BuildVaultFile builds the index.html file for the vault
func BuildVaultFile(cnfg *dwn.Config) (files.Node, error) {
	w := bytes.NewBuffer(nil)
	err := views.VaultIndexFile().Render(context.Background(), w)
	if err != nil {
		return nil, err
	}
	indexData := w.Bytes()
	return files.NewBytesFile(indexData), nil
}
