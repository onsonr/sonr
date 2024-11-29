package types

import (
	"encoding/json"

	"github.com/ipfs/boxo/files"
	"github.com/onsonr/sonr/pkg/crypto/mpc"
	"github.com/onsonr/sonr/web/vault/types"
	"github.com/onsonr/sonr/x/dwn/types/static"
)

const (
	kFileNameConfigJSON = "dwn.json"
	kFileNameIndexHTML  = "index.html"
	kFileNameWorkerJS   = "sw.js"
)

type vault struct {
	Config *types.Config
	Schema *types.Schema
	Source KeyshareSource
	ks     mpc.Keyset
}

func SpawnVault(chainID string, schema *types.Schema) (files.Directory, error) {
	ks, err := mpc.NewKeyset()
	if err != nil {
		return nil, err
	}
	src, err := createKeySource(ks)
	if err != nil {
		return nil, err
	}
	tk, err := src.OriginToken()
	if err != nil {
		return nil, err
	}
	kscid, err := tk.CID()
	if err != nil {
		return nil, err
	}
	dwnCfg := &types.Config{
		MotrKeyshare:   kscid.String(),
		MotrAddress:    src.Address(),
		IpfsGatewayUrl: "https://ipfs.sonr.land",
		SonrApiUrl:     "https://api.sonr.land",
		SonrRpcUrl:     "https://rpc.sonr.land",
		SonrChainId:    chainID,
		VaultSchema:    schema,
	}
	cnf, err := json.Marshal(dwnCfg)
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
