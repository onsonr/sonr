package keeper

import (
	"errors"
	"fmt"

	gocache "github.com/patrickmn/go-cache"
	"github.com/sonr-hq/sonr/pkg/node/config"
	"github.com/sonr-hq/sonr/pkg/vault/core/wallet"
	v1 "github.com/sonr-hq/sonr/pkg/vault/types/v1"
	"github.com/sonr-hq/sonr/x/identity/types"
)

type VaultBank struct {
	// The IPFS node that the vault is running on
	node config.IPFSNode

	// The wallet that the vault is using
	cache *gocache.Cache
}

// Creates a new Vault
func CreateBank(node config.IPFSNode, cache *gocache.Cache) *VaultBank {
	return &VaultBank{
		node:  node,
		cache: cache,
	}
}

func (v *VaultBank) StartRegistration(entry *Session) (string, string, error) {
	optsJson, err := entry.BeginRegistration()
	if err != nil {
		return "", "", err
	}
	v.putEntryIntoCache(entry)
	return optsJson, entry.ID, nil
}

func (v *VaultBank) FinishRegistration(sessionId string, credsJson string) (*types.DidDocument, *v1.WalletConfig, error) {
	// Get Session
	entry, err := v.getEntryFromCache(sessionId)
	if err != nil {
		return nil, nil, err
	}
	didDoc, err := entry.FinishRegistration(credsJson)
	if err != nil {
		return nil, nil, err
	}
	// Create a new offline wallet
	wallet, err := wallet.NewWallet()
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("Failed to create new offline wallet using MPC: %s", err))
	}

	primAcc, err := wallet.PrimaryAccount()
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("Failed to get primary account: %s", err))
	}
	didDoc.AddAssertion(primAcc.GetAssertionMethod())
	return didDoc, wallet.WalletConfig(), nil
}
