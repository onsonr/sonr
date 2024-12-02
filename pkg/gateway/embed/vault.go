package embed

import (
	"github.com/ipfs/boxo/files"
)

const (
	AppManifestFileName   = "app.webmanifest"
	DWNConfigFileName     = "dwn.json"
	IndexHTMLFileName     = "index.html"
	MainJSFileName        = "main.js"
	ServiceWorkerFileName = "sw.js"
)

// spawnVaultDirectory creates a new directory with the default files
func NewFS(cfgBz []byte, manifestBz []byte) files.Directory {
	return files.NewMapDirectory(map[string]files.Node{
		AppManifestFileName:   files.NewBytesFile(manifestBz),
		DWNConfigFileName:     files.NewBytesFile(cfgBz),
		IndexHTMLFileName:     files.NewBytesFile(IndexHTML),
		MainJSFileName:        files.NewBytesFile(MainJS),
		ServiceWorkerFileName: files.NewBytesFile(WorkerJS),
	})
}
