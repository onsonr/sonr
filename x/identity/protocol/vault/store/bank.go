package store

import (
	"errors"
	"fmt"

	gocache "github.com/patrickmn/go-cache"
	"github.com/sonr-hq/sonr/pkg/node/config"
	"github.com/sonr-hq/sonr/x/identity/protocol/vault/wallet"
	"github.com/sonr-hq/sonr/x/identity/types"
	v1 "github.com/sonr-hq/sonr/x/identity/types/vault/v1"
)

type VaultBank struct {
	// The IPFS node that the vault is running on
	node config.IPFSNode

	// The wallet that the vault is using
	cache *gocache.Cache
}

// Creates a new Vault
func InitBank(node config.IPFSNode, cache *gocache.Cache) *VaultBank {
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

func (v *VaultBank) StartLogin(entry *Session) (string, string, error) {
	optsJson, err := entry.BeginLogin()
	if err != nil {
		return "", "", err
	}
	v.putEntryIntoCache(entry)
	return optsJson, entry.ID, nil
}

func (v *VaultBank) FinishLogin(sessionId string, credsJson string) (bool, error) {
	// Get Session
	entry, err := v.getEntryFromCache(sessionId)
	if err != nil {
		return false, err
	}
	didDoc, err := entry.FinishLogin(credsJson)
	if err != nil {
		return false, err
	}
	return didDoc, nil
}

func (v *VaultBank) getEntryFromCache(id string) (*Session, error) {
	val, ok := v.cache.Get(id)
	if !ok {
		return nil, errors.New("Failed to find entry for ID")
	}
	e, ok := val.(*Session)
	if !ok {
		return nil, errors.New("Invalid type for session entry")
	}
	return e, nil
}

func (v *VaultBank) putEntryIntoCache(entry *Session) error {
	if entry == nil {
		return errors.New("Entry cannot be nil to put into cache")
	}
	return v.cache.Add(entry.ID, entry, -1)
}
