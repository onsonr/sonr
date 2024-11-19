package types

import (
	"github.com/ipfs/boxo/files"

	"github.com/onsonr/sonr/pkg/motr"
	"github.com/onsonr/sonr/pkg/motr/config"
)

type Vault struct {
	FS files.Node
}

func NewVault(keyshareJSON string, adddress string, chainID string, schema *config.Schema) (*Vault, error) {
	dwnCfg := &config.Config{
		MotrKeyshare:   keyshareJSON,
		MotrAddress:    adddress,
		IpfsGatewayUrl: "https://ipfs.sonr.land",
		SonrApiUrl:     "https://api.sonr.land",
		SonrRpcUrl:     "https://rpc.sonr.land",
		SonrChainId:    chainID,
		VaultSchema:    schema,
	}
	fileMap, err := motr.NewVaultDirectory(dwnCfg)
	if err != nil {
		return nil, err
	}
	return &Vault{
		FS: fileMap,
	}, nil
}
