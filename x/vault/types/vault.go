package types

import (
	"github.com/ipfs/boxo/files"

	"github.com/onsonr/sonr/config/dwn"
	vault "github.com/onsonr/sonr/x/vault/internal"
)

type Vault struct {
	FS files.Node
}

func NewVault(cnfg *dwn.Config, chainID string) (*Vault, error) {
	fileMap, err := vault.NewVaultDirectory(cnfg)
	if err != nil {
		return nil, err
	}
	return &Vault{
		FS: fileMap,
	}, nil
}
