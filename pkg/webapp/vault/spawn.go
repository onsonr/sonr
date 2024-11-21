package vault

import (
	"github.com/ipfs/boxo/files"

	"github.com/onsonr/sonr/pkg/common/middleware/render"
	"github.com/onsonr/sonr/pkg/core/dwn"
)

const (
	kFileNameConfigJSON = "dwn.json"
	kFileNameIndexHTML  = "index.html"
)

type Vault = files.Directory

func SpawnVault(keyshareJSON string, adddress string, chainID string, schema *dwn.Schema) (Vault, error) {
	dwnCfg := &dwn.Config{
		MotrKeyshare:   keyshareJSON,
		MotrAddress:    adddress,
		IpfsGatewayUrl: "https://ipfs.sonr.land",
		SonrApiUrl:     "https://api.sonr.land",
		SonrRpcUrl:     "https://rpc.sonr.land",
		SonrChainId:    chainID,
		VaultSchema:    schema,
	}
	return setupVaultDirectory(dwnCfg)
}

// spawnVaultDirectory creates a new directory with the default files
func setupVaultDirectory(cnfg *dwn.Config) (files.Directory, error) {
	idxf, err := render.TemplRawBytes(IndexFile())
	if err != nil {
		return nil, err
	}

	cnf, err := cnfg.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return files.NewMapDirectory(map[string]files.Node{
		kFileNameConfigJSON: files.NewBytesFile(cnf),
		kFileNameIndexHTML:  files.NewBytesFile(idxf),
	}), nil
}
