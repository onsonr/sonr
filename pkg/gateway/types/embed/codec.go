package embed

import (
	"encoding/json"

	"github.com/ipfs/boxo/files"
	config "github.com/onsonr/sonr/internal/config/motr"
	"github.com/onsonr/sonr/internal/models"
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
func NewVaultFS(cfg *config.Config) (files.Directory, error) {
	manifestBz, err := models.NewWebManifest()
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
		IndexHTMLFileName:     files.NewBytesFile(IndexHTML),
		MainJSFileName:        files.NewBytesFile(MainJS),
		ServiceWorkerFileName: files.NewBytesFile(WorkerJS),
	}), nil
}

// NewVaultConfig returns the default vault config
func NewVaultConfig(addr string, ucanCID string) *config.Config {
	return &config.Config{
		MotrToken:      ucanCID,
		MotrAddress:    addr,
		IpfsGatewayUrl: "http://localhost:80",
		SonrApiUrl:     "http://localhost:1317",
		SonrRpcUrl:     "http://localhost:26657",
		SonrChainId:    "sonr-testnet-1",
		VaultSchema:    DefaultSchema(),
	}
}

// DefaultSchema returns the default schema
func DefaultSchema() *config.Schema {
	return &config.Schema{
		Version:    SchemaVersion,
		Account:    getSchema(&models.Account{}),
		Asset:      getSchema(&models.Asset{}),
		Chain:      getSchema(&models.Chain{}),
		Credential: getSchema(&models.Credential{}),
		Grant:      getSchema(&models.Grant{}),
		Keyshare:   getSchema(&models.Keyshare{}),
		Profile:    getSchema(&models.Profile{}),
	}
}
