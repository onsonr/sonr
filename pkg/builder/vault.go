package builder

import (
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/ipfs/boxo/files"
	"github.com/onsonr/crypto/mpc"
	"github.com/onsonr/sonr/pkg/vault"
)

type Vault struct {
	FS    files.Node
	ValKs mpc.Share
}

func NewVault(subject string, origin string, chainID string) (*Vault, error) {
	shares, err := mpc.GenerateKeyshares()
	var (
		valKs = shares[0]
		usrKs = shares[1]
	)
	usrKsJSON, err := usrKs.Marshal()
	if err != nil {
		return nil, err
	}
	sonrAddr, err := bech32.ConvertAndEncode("idx", valKs.GetPublicKey())
	if err != nil {
		return nil, err
	}

	cnfg := vault.NewConfig(usrKsJSON, sonrAddr, chainID, DefaultSchema())
	cnfgFile, err := vault.MarshalConfigFile(cnfg)
	if err != nil {
		return nil, err
	}

	idxFile, err := vault.IndexHTMLFile(cnfg)
	if err != nil {
		return nil, err
	}

	fileMap := map[string]files.Node{
		"config.json": cnfgFile,
		"motr.mjs":    vault.MotrMJSFile(),
		"sw.js":       vault.SWJSFile(),
		"app.wasm":    vault.DWNWasmFile(),
		"index.html":  idxFile,
	}

	return &Vault{
		FS:    files.NewMapDirectory(fileMap),
		ValKs: valKs,
	}, nil
}
