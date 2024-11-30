package gateway

import (
	"github.com/ipfs/boxo/files"
	"github.com/onsonr/sonr/pkg/gateway/embed"
)

const (
	AppManifestFileName   = "app.webmanifest"
	DWNConfigFileName     = "dwn.json"
	IndexHTMLFileName     = "index.html"
	MainJSFileName        = "main.js"
	ServiceWorkerFileName = "sw.js"
)

// spawnVaultDirectory creates a new directory with the default files
func setupVaultDirectory(cfgBz []byte, manifestBz []byte) files.Directory {
	return files.NewMapDirectory(map[string]files.Node{
		AppManifestFileName:   files.NewBytesFile(manifestBz),
		DWNConfigFileName:     files.NewBytesFile(cfgBz),
		IndexHTMLFileName:     files.NewBytesFile(embed.IndexHTML),
		MainJSFileName:        files.NewBytesFile(embed.MainJS),
		ServiceWorkerFileName: files.NewBytesFile(embed.WorkerJS),
	})
}
