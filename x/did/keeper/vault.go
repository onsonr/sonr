package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/core/coreiface/options"
	"github.com/onsonr/crypto/mpc"
	"github.com/onsonr/sonr/internal/vfs"
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
	cnfg, err := vfs.NewDWNConfigFile(usrKsJSON, sonrAddr, chainID)
	if err != nil {
		return nil, err
	}
	fileMap := map[string]files.Node{
		"config.json": cnfg,
		"sw.js":       vfs.SWJSFile(),
		"app.wasm":    vfs.DWNWasmFile(),
		"index.html":  vfs.IndexFile(),
	}
	return &Vault{
		FS:    files.NewMapDirectory(fileMap),
		ValKs: valKs,
	}, nil
}

// AssembleVault assembles the initial vault
func (k Keeper) AssembleVault(ctx Context, subject string, origin string) (string, int64, error) {
	v, err := NewVault(subject, origin, "sonr-testnet")
	if err != nil {
		return "", 0, err
	}
	cid, err := k.ipfsClient.Unixfs().Add(context.Background(), v.FS)
	if err != nil {
		return "", 0, err
	}
	return cid.String(), ctx.CalculateExpiration(time.Second * 15), nil
}

// PinVaultController pins the initial vault to the local IPFS node
func (k Keeper) PinVaultController(_ sdk.Context, cid string, address string) (bool, error) {
	// Resolve the path
	path, err := path.NewPath(cid)
	if err != nil {
		return false, err
	}

	// 1. Initialize vault.db sqlite database in local IPFS with Mount

	// 2. Insert the InitialWalletAccounts

	// 3. Publish the path to the IPNS
	_, err = k.ipfsClient.Name().Publish(context.Background(), path, options.Name.Key(address))
	if err != nil {
		return false, err
	}

	// 4. Insert the accounts into x/auth

	// 5. Insert the controller into state
	return true, nil
}
