package bank

import (
	"context"
	"errors"

	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/node"
	"github.com/sonr-hq/sonr/pkg/vault/internal/fs"
	"github.com/sonr-hq/sonr/pkg/vault/internal/mpc"
	"github.com/sonr-hq/sonr/pkg/vault/internal/network"
	"github.com/sonr-hq/sonr/x/identity/types"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
)

// It creates a new wallet with two participants, one of which is the current participant, and returns
// the wallet
func (v *VaultBank) buildWallet(prefix string) (common.Wallet, error) {
	participants := party.IDSlice{"current", "vault"}
	net := network.NewOfflineNetwork(participants)
	wsl, err := mpc.Keygen("current", 1, net, prefix)
	if err != nil {
		return nil, err
	}
	wallet := network.OfflineWallet(wsl)
	// Create a new OfflineWallet from the WalletShares
	// for _, share := range wsl {
	// 	_, err := share.Marshal()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	// err = vaultfs.StoreShare(buf, string(share.SelfID()), "insecure-password-testnet")
	// 	// if err != nil {
	// 	// 	return nil, err
	// 	// }
	// }
	return wallet, nil
}

func (v *VaultBank) loadWallet(ctx context.Context, didDoc *types.DidDocument, node node.Node) (common.Wallet, fs.VaultFS, error) {
	if s := didDoc.GetVaultService(); s != nil {
		//cfgs, err := vaultfs.LoadShares()
	}
	return nil, nil, errors.New("Unimplemented")

}

// Loads an OfflineWallet from a []*WalletShareConfig and returns a `common.Wallet` interface
func (v *VaultBank) loadOfflineWallet(shareConfigs []*common.WalletShareConfig) (common.Wallet, error) {
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
