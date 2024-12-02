package config

import (
	"encoding/json"

	"github.com/ipfs/boxo/files"
	"github.com/onsonr/sonr/pkg/vault/config/internal"
	"github.com/onsonr/sonr/pkg/vault/types"
)

const SchemaVersion = 1
const (
	AppManifestFileName   = "app.webmanifest"
	DWNConfigFileName     = "dwn.json"
	IndexHTMLFileName     = "index.html"
	MainJSFileName        = "main.js"
	ServiceWorkerFileName = "sw.js"
)

// spawnVaultDirectory creates a new directory with the default files
func NewFS(cfg *types.Config) (files.Directory, error) {
	manifestBz, err := newWebManifestBytes()
	if err != nil {
		return nil, err
	}
	cnfBz, err := json.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	return files.NewMapDirectory(map[string]files.Node{
		AppManifestFileName:   files.NewBytesFile(manifestBz),
		DWNConfigFileName:     files.NewBytesFile(cnfBz),
		IndexHTMLFileName:     files.NewBytesFile(internal.IndexHTML),
		MainJSFileName:        files.NewBytesFile(internal.MainJS),
		ServiceWorkerFileName: files.NewBytesFile(internal.WorkerJS),
	}), nil
}
