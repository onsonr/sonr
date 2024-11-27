package types

import (
	"github.com/ipfs/boxo/files"

	"github.com/onsonr/sonr/pkg/core/dwn"
	"github.com/onsonr/sonr/x/dwn/types/static"
)

const (
	kFileNameConfigJSON = "dwn.json"
	kFileNameIndexHTML  = "index.html"
	kFileNameWorkerJS   = "sw.js"
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
	cnf, err := dwnCfg.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return setupVaultDirectory(cnf), nil
}

// spawnVaultDirectory creates a new directory with the default files
func setupVaultDirectory(cfgBz []byte) files.Directory {
	return files.NewMapDirectory(map[string]files.Node{
		kFileNameConfigJSON: files.NewBytesFile(cfgBz),
		kFileNameIndexHTML:  files.NewBytesFile(static.IndexHTML),
		kFileNameWorkerJS:   files.NewBytesFile(static.WorkerJS),
	})
}
