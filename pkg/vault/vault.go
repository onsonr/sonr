package vault

import (
	"context"
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	gocache "github.com/patrickmn/go-cache"
	cl "github.com/sonr-hq/sonr/pkg/client"
	"github.com/sonr-hq/sonr/pkg/common"
	ipfs "github.com/sonr-hq/sonr/pkg/node/ipfs"
	"github.com/sonr-hq/sonr/pkg/vault/fs"
	"github.com/sonr-hq/sonr/pkg/vault/mpc"
	"github.com/sonr-hq/sonr/pkg/vault/network"
	"github.com/sonr-hq/sonr/pkg/vault/session"
	"github.com/sonr-hq/sonr/x/identity/types"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
)

// `VaultBank` is a struct that contains an IPFS node, a cache, a channel, a client context, and a
// client stub.
// @property node - The IPFS node that the vault is running on
// @property cache - This is a cache that stores the wallets that the vault has already processed.
// @property done - This is a channel that will be used to send completed wallets to the vault.
// @property cctx - The client context that the vault is running on.
// @property client - This is the client stub that the vault uses to communicate with the client.
type VaultBank struct {
	// The IPFS node that the vault is running on
	node ipfs.IPFS

	// The wallet that the vault is using
	cache *gocache.Cache

	// Completed wallets from channel
	done chan common.Wallet

	// Client Context
	cctx client.Context

	// Client Stub
	client *cl.ClientStub
}

// `WalletResult` is a struct with three fields: `Wallet`, `Vault`, and `Err`.
//
// The `Wallet` field is of type `common.Wallet`, which is a struct with two fields: `Name` and
// `Address`.
//
// The `Vault` field is of type `fs.VaultFS`, which is a struct with two fields: `Path` and `Password`.
//
// The `Err` field is of type `error`.
// @property Wallet - The wallet object that was created.
// @property Vault - The vault that the wallet is stored in.
// @property {error} Err - This is the error that occurred during the wallet creation process.
type WalletResult struct {
	Wallet common.Wallet
	Vault  fs.VaultFS
	Err    error
}

// Creates a new Vault
func NewVaultBank(cctx client.Context, node ipfs.IPFS, cache *gocache.Cache) *VaultBank {
	stub := cl.NewStub(cctx)
	return &VaultBank{
		node:   node,
		cache:  cache,
		cctx:   cctx,
		client: stub,
	}
}

// Creating a new session entry and returning the options json and the session id.
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

// Creating a new wallet with two participants, one of which is the current participant, and returns
// the wallet
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
	done := make(chan *WalletResult)
	go func() {
		wallet, vfs, err := buildWallet(context.Background(), "snr", v.node, "testnet-not-so-secret-password")
		done <- &WalletResult{
			Wallet: wallet,
			Vault:  vfs,
			Err:    err,
		}
	}()
	result := <-done
	err = didDoc.SetRootWallet(result.Wallet)
	if err != nil {
		return nil, err
	}
	didDoc.AddService(result.Vault.Service())
	return didDoc, nil
}

func Authenticate(didDoc *types.DidDocument, node ipfs.IPFS) {
	didDoc.Service.FindByID(fmt.Sprintf("did:ipfs:%s", didDoc.Address()))
}

// It creates a new wallet with two participants, one of which is the current participant, and returns
// the wallet
func buildWallet(ctx context.Context, prefix string, node ipfs.IPFS, password string) (common.Wallet, fs.VaultFS, error) {
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
		err = vaultfs.StoreShare(buf, string(share.SelfID()), password)
		if err != nil {
			return nil, nil, err
		}

	}
	return wallet, vaultfs, nil
}

// `loadWallet` loads a wallet from a DID Document and a password
func loadWallet(ctx context.Context, didDoc *types.DidDocument, node ipfs.IPFS, password string) (common.Wallet, fs.VaultFS, error) {
	if s := didDoc.GetVaultService(); s != nil {
		vaultfs, err := fs.New(ctx, didDoc.Address(), node, fs.WithIPFSPath(s.CID()))
		if err != nil {
			return nil, nil, err
		}
		cfgs, err := vaultfs.LoadShares(password)
		if err != nil {
			return nil, nil, err
		}

		// Load the OfflineWallet
		wallet, err := loadOfflineWallet(cfgs)
		if err != nil {
			return nil, nil, err
		}
		return wallet, vaultfs, nil
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
