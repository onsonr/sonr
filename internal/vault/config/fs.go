package config

import (
	"encoding/json"

	"github.com/ipfs/boxo/files"
	"github.com/onsonr/sonr/internal/vault/embed"
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
func NewVaultFS(cfg *Config) (files.Directory, error) {
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
		IndexHTMLFileName:     files.NewBytesFile(embed.IndexHTML),
		MainJSFileName:        files.NewBytesFile(embed.MainJS),
		ServiceWorkerFileName: files.NewBytesFile(embed.WorkerJS),
	}), nil
}

// NewVaultConfig returns the default vault config
func NewVaultConfig(addr string, ucanCID string) *Config {
	return &Config{
		MotrToken:      ucanCID,
		MotrAddress:    addr,
		IpfsGatewayUrl: "http://localhost:80",
		SonrApiUrl:     "http://localhost:1317",
		SonrRpcUrl:     "http://localhost:26657",
		SonrChainId:    "sonr-testnet-1",
		VaultSchema:    DefaultSchema(),
	}
}
