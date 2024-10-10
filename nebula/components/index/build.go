package index

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/ipfs/boxo/files"
	"github.com/onsonr/sonr/internal/dwn/gen"
)

func BuildFile(cnfg *gen.Config) (files.Node, error) {
	_, err := json.Marshal(cnfg)
	if err != nil {
		return nil, err
	}
	w := bytes.NewBuffer(nil)
	err = IndexFile().Render(context.Background(), w)
	if err != nil {
		return nil, err
	}
	indexData := w.Bytes()
	return files.NewBytesFile(indexData), nil
}
