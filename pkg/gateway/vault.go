package gateway

import (
	"github.com/ipfs/boxo/files"
	"github.com/onsonr/sonr/web/vault/embed"
)

const (
	kFileNameConfigJSON = "dwn.json"
	kFileNameIndexHTML  = "index.html"
	kFileNameWorkerJS   = "sw.js"
)

// spawnVaultDirectory creates a new directory with the default files
func setupVaultDirectory(cfgBz []byte) files.Directory {
	return files.NewMapDirectory(map[string]files.Node{
		kFileNameConfigJSON: files.NewBytesFile(cfgBz),
		kFileNameIndexHTML:  files.NewBytesFile(embed.IndexHTML),
		kFileNameWorkerJS:   files.NewBytesFile(embed.WorkerJS),
	})
}
