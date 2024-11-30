package gateway

import (
	"github.com/ipfs/boxo/files"
	"github.com/onsonr/sonr/pkg/gateway/embed"
)

const (
	DWNConfigFileName     = "dwn.json"
	IndexHTMLFileName     = "index.html"
	ServiceWorkerFileName = "sw.js"
)

// spawnVaultDirectory creates a new directory with the default files
func setupVaultDirectory(cfgBz []byte) files.Directory {
	return files.NewMapDirectory(map[string]files.Node{
		DWNConfigFileName:     files.NewBytesFile(cfgBz),
		IndexHTMLFileName:     files.NewBytesFile(embed.IndexHTML),
		ServiceWorkerFileName: files.NewBytesFile(embed.WorkerJS),
	})
}
