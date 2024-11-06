package nebula

import (
	"bytes"
	"context"

	"github.com/ipfs/boxo/files"
	"github.com/onsonr/sonr/internal/dwn/gen"
	"github.com/onsonr/sonr/pkg/nebula/views"
)

// BuildVaultFile builds the index.html file for the vault
func BuildVaultFile(cnfg *gen.Config) (files.Node, error) {
	w := bytes.NewBuffer(nil)
	err := views.VaultIndexFile().Render(context.Background(), w)
	if err != nil {
		return nil, err
	}
	indexData := w.Bytes()
	return files.NewBytesFile(indexData), nil
}
