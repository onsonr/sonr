package vault

import (
	"context"

	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/node/ipfs"
	"github.com/sonr-hq/sonr/pkg/vault/fs"
	"github.com/sonr-hq/sonr/pkg/vault/mpc"
	"github.com/sonr-hq/sonr/pkg/vault/network"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
)

// It creates a new wallet with two participants, one of which is the current participant, and returns
// the wallet
func NewWallet(ctx context.Context, prefix string, node *ipfs.IPFS) (common.Wallet, error) {
	getPrefix := func() string {
		if len(prefix) == 0 {
			return "snr"
		}
		return prefix
	}
	participants := party.IDSlice{"current", "vault"}
	net := network.NewOfflineNetwork(participants)
	wsl, err := mpc.Keygen("current", 1, net, getPrefix())
	if err != nil {
		return nil, err
	}
	wallet := network.OfflineWallet(wsl)
	vaultfs, err := fs.New(ctx, wallet.Address(), node.CoreAPI)
	if err != nil {
		return nil, err
	}

	// Create a new OfflineWallet from the WalletShares
	for _, share := range wsl {
		buf, err := share.Marshal()
		if err != nil {
			return nil, err
		}
		err = vaultfs.StoreShare(buf, string(share.SelfID()))
		if err != nil {
			return nil, err
		}

	}
	return wallet, nil
}

// Loads an OfflineWallet from a []*WalletShareConfig and returns a `common.Wallet` interface
func LoadOfflineWallet(shareConfigs []*common.WalletShareConfig) (common.Wallet, error) {
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
