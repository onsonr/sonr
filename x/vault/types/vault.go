package types

import (
	"github.com/ipfs/boxo/files"

	"github.com/onsonr/sonr/internal/dwn"
	dwngen "github.com/onsonr/sonr/internal/dwn/gen"
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
	fileMap, err := dwn.NewVaultDirectory(dwnCfg)
	if err != nil {
		return nil, err
	}
	return &Vault{
		FS: fileMap,
	}, nil
}
