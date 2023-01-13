package vault

import (
	"context"
	"errors"
	"fmt"

	gocache "github.com/patrickmn/go-cache"
	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/node"
	"github.com/sonr-hq/sonr/pkg/vault/internal/fs"
	"github.com/sonr-hq/sonr/pkg/vault/internal/mpc"
	"github.com/sonr-hq/sonr/pkg/vault/internal/network"
	"github.com/sonr-hq/sonr/pkg/vault/internal/session"
	"github.com/sonr-hq/sonr/x/identity/types"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
)

type VaultBank struct {
	// The IPFS node that the vault is running on
	node node.Node

	// The wallet that the vault is using
	cache *gocache.Cache

	// Completed wallets from channel
	done chan common.Wallet
}

// Creates a new Vault
func NewVaultBank(node node.Node, cache *gocache.Cache) *VaultBank {
	return &VaultBank{
		node:  node,
		cache: cache,
	}
}

func (v *VaultBank) StartRegistration(rpid string, aka string) (string, string, error) {
	entry, err := session.NewEntry(rpid, aka)
	if err != nil {
		return "", "", err
	}
	optsJson, err := entry.BeginRegistration()
	if err != nil {
		return "", "", err
	}
	v.cache.Set(entry.ID, entry, gocache.DefaultExpiration)
	return optsJson, entry.ID, nil
}

func (v *VaultBank) FinishRegistration(sessionId string, credsJson string) (*types.DidDocument, error) {
	// Get Session
	entry, err := session.GetEntry(sessionId, v.cache)
	if err != nil {
		return nil, err
	}
	didDoc, err := entry.FinishRegistration(credsJson)
	if err != nil {
		return nil, err
	}
	// Create a new offline wallet
	wallet, vfs, err := buildWallet(context.Background(), "snr", v.node)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to create new offline wallet using MPC: %s", err))
	}
	err = didDoc.SetRootWallet(wallet)
	if err != nil {
		return nil, err
	}
	didDoc.AddService(vfs.Service())
	return didDoc, nil
}

// It creates a new wallet with two participants, one of which is the current participant, and returns
// the wallet
func buildWallet(ctx context.Context, prefix string, node node.Node) (common.Wallet, fs.VaultFS, error) {
	participants := party.IDSlice{"current", "vault"}
	net := network.NewOfflineNetwork(participants)
	wsl, err := mpc.Keygen("current", 1, net, prefix)
	if err != nil {
		return nil, nil, err
	}
	wallet := network.OfflineWallet(wsl)
	vaultfs, err := fs.New(ctx, wallet.Address(), node)
	if err != nil {
		return nil, nil, err
	}

	// Create a new OfflineWallet from the WalletShares
	for _, share := range wsl {
		buf, err := share.Marshal()
		if err != nil {
			return nil, nil, err
		}
		err = vaultfs.StoreShare(buf, string(share.SelfID()), "insecure-password-testnet")
		if err != nil {
			return nil, nil, err
		}

	}
	return wallet, vaultfs, nil
}

func loadWallet(ctx context.Context, didDoc *types.DidDocument, node node.Node) (common.Wallet, fs.VaultFS, error) {
	if s := didDoc.GetVaultService(); s != nil {
		_, err := fs.New(ctx, didDoc.Address(), node, fs.WithIPFSPath(s.CID()))
		if err != nil {
			return nil, nil, err
		}
		//cfgs, err := vaultfs.LoadShares()
	}
	return nil, nil, errors.New("Unimplemented")

}

// Loads an OfflineWallet from a []*WalletShareConfig and returns a `common.Wallet` interface
func loadOfflineWallet(shareConfigs []*common.WalletShareConfig) (common.Wallet, error) {
	// Convert the WalletShareConfigs to WalletShares
	ws := make([]common.WalletShare, 0)
	for i, shareConfig := range shareConfigs {
		if s, err := mpc.LoadWalletShare(shareConfig); err != nil {
			return nil, err
		} else {
			ws[i] = s
		}
	}
	return network.OfflineWallet(ws), nil
}
