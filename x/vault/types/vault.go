package types

import (
	"github.com/ipfs/boxo/files"

	"github.com/onsonr/sonr/pkg/core/dwn"
)

const (
	FileNameConfigJSON = "dwn.json"
	FileNameIndexHTML  = "index.html"
)

type Vault = files.Directory

func NewVault(keyshareJSON string, adddress string, chainID string, schema *dwn.Schema) (Vault, error) {
	dwnCfg := &dwn.Config{
		MotrKeyshare:   keyshareJSON,
		MotrAddress:    adddress,
		IpfsGatewayUrl: "https://ipfs.sonr.land",
		SonrApiUrl:     "https://api.sonr.land",
		SonrRpcUrl:     "https://rpc.sonr.land",
		SonrChainId:    chainID,
		VaultSchema:    schema,
	}
	return dwn.SpawnVault(dwnCfg)
}
