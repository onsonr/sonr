package types

import (
	"github.com/ipfs/boxo/files"

	dwngen "github.com/onsonr/sonr/pkg/motr/config"
	"github.com/onsonr/sonr/pkg/motr/embed"
)

type Vault struct {
	FS files.Node
}

func NewVault(keyshareJSON string, adddress string, chainID string, schema *dwngen.Schema) (*Vault, error) {
	dwnCfg := &dwngen.Config{
		MotrKeyshare:   keyshareJSON,
		MotrAddress:    adddress,
		IpfsGatewayUrl: "https://ipfs.sonr.land",
		SonrApiUrl:     "https://api.sonr.land",
		SonrRpcUrl:     "https://rpc.sonr.land",
		SonrChainId:    chainID,
		VaultSchema:    schema,
	}
	fileMap, err := embed.NewVaultDirectory(dwnCfg)
	if err != nil {
		return nil, err
	}
	return &Vault{
		FS: fileMap,
	}, nil
}
