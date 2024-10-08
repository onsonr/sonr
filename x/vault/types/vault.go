package types

import (
	"github.com/ipfs/boxo/files"

	"github.com/onsonr/sonr/pkg/dwn"
	dwngen "github.com/onsonr/sonr/pkg/dwn/gen"
)

type Vault struct {
	FS files.Node
}

func NewVault(keyshareJSON string, adddress string, chainID string, schema *dwngen.Schema) (*Vault, error) {
	dwnCfg := &dwngen.Config{
		Motr:   createMotrConfig(keyshareJSON, adddress, "sonr.id"),
		Ipfs:   defaultIPFSConfig(),
		Sonr:   defaultSonrConfig(chainID),
		Schema: schema,
	}
	fileMap, err := dwn.NewVaultDirectory(dwnCfg)
	if err != nil {
		return nil, err
	}
	return &Vault{
		FS: fileMap,
	}, nil
}

func createMotrConfig(keyshareJSON string, adddress string, origin string) *dwngen.Motr {
	return &dwngen.Motr{
		Keyshare: keyshareJSON,
		Address:  adddress,
		Origin:   origin,
	}
}

func defaultIPFSConfig() *dwngen.IPFS {
	return &dwngen.IPFS{
		ApiUrl:     "https://api.sonr-ipfs.land",
		GatewayUrl: "https://ipfs.sonr.land",
	}
}

func defaultSonrConfig(chainID string) *dwngen.Sonr {
	return &dwngen.Sonr{
		ApiUrl:       "https://api.sonr.land",
		GrpcUrl:      "https://grpc.sonr.land",
		RpcUrl:       "https://rpc.sonr.land",
		WebSocketUrl: "wss://rpc.sonr.land/ws",
		ChainId:      chainID,
	}
}
