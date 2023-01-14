package bank

import (
	"errors"
	"fmt"

	gocache "github.com/patrickmn/go-cache"
	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/node/config"
	"github.com/sonr-hq/sonr/pkg/vault/internal/network"
	"github.com/sonr-hq/sonr/pkg/vault/internal/session"
	"github.com/sonr-hq/sonr/x/identity/types"
)

type VaultBank struct {
	// The IPFS node that the vault is running on
	node config.IPFSNode

	// The wallet that the vault is using
	cache *gocache.Cache

	// Completed wallets from channel
	done chan common.Wallet
}

// Creates a new Vault
func CreateBank(node config.IPFSNode, cache *gocache.Cache) *VaultBank {
	return &VaultBank{
		node:  node,
		cache: cache,
	}
}

func (v *VaultBank) StartRegistration(entry *session.Session) (string, string, error) {
	optsJson, err := entry.BeginRegistration()
	if err != nil {
		return "", "", err
	}
	v.putEntryIntoCache(entry)
	return optsJson, entry.ID, nil
}

func (v *VaultBank) FinishRegistration(sessionId string, credsJson string) (*types.DidDocument, network.OfflineWallet, error) {
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
	wallet, err := v.buildWallet("snr")
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("Failed to create new offline wallet using MPC: %s", err))
	}
	return didDoc, wallet, nil
}
