package types

import (
	"github.com/ipfs/boxo/files"

	"github.com/onsonr/sonr/pkg/dwn"
	vault "github.com/onsonr/sonr/x/vault/types/internal"
)

type Vault struct {
	FS files.Node
}

func NewVault(keyshareJSON string, adddress string, chainID string, schema *dwn.Schema) (*Vault, error) {
	dwnCfg := &dwn.Config{
		Motr:   createMotrConfig(keyshareJSON, adddress, "sonr.id"),
		Ipfs:   defaultIPFSConfig(),
		Sonr:   defaultSonrConfig(chainID),
		Schema: schema,
	}
	fileMap, err := vault.NewVaultDirectory(dwnCfg)
	if err != nil {
		return nil, err
	}
	return &Vault{
		FS: fileMap,
	}, nil
}

func createMotrConfig(keyshareJSON string, adddress string, origin string) *dwn.Motr {
	return &dwn.Motr{
		Keyshare: keyshareJSON,
		Address:  adddress,
		Origin:   origin,
	}
}

func defaultIPFSConfig() *dwn.IPFS {
	return &dwn.IPFS{
		ApiUrl:     "https://api.sonr-ipfs.land",
		GatewayUrl: "https://ipfs.sonr.land",
	}
}

func defaultSonrConfig(chainID string) *dwn.Sonr {
	return &dwn.Sonr{
		ApiUrl:       "https://api.sonr.land",
		GrpcUrl:      "https://grpc.sonr.land",
		RpcUrl:       "https://rpc.sonr.land",
		WebSocketUrl: "wss://rpc.sonr.land/ws",
		ChainId:      chainID,
	}
}
