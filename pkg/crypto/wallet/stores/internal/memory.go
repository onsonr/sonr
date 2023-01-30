package internal

import (
	"fmt"
	"sync"

	"github.com/sonrhq/core/pkg/crypto/wallet"
	vaultv1 "github.com/sonrhq/core/x/identity/types/vault/v1"
)

type MemoryStore struct {
	accConfig *vaultv1.AccountConfig
	configs   map[string]wallet.Account
	sync.Mutex
}

func NewMemoryStore(accCfg *vaultv1.AccountConfig) (wallet.Store, error) {
	ds := &MemoryStore{
		accConfig: accCfg,
		configs:   make(map[string]wallet.Account),
	}

	return ds, nil
}

func (ds *MemoryStore) GetAccount(name string) (wallet.Account, error) {
	ds.Lock()
	defer ds.Unlock()
	s, ok := ds.configs[name]
	if !ok {
		return nil, fmt.Errorf("account not found. name: %s", name)
	}
	return s, nil
}

func (ds *MemoryStore) PutAccount(sc wallet.Account, name string) error {
	ds.Lock()
	defer ds.Unlock()
	ds.configs[name] = sc
	return nil
}

// JWKClaims returns the JWKClaims for the store to be signed by the identity
func (ds *MemoryStore) JWKClaims(acc wallet.Account) (string, error) {
	return "", nil
}

// VerifyJWKClaims verifies the JWKClaims for the store
func (ds *MemoryStore) VerifyJWKClaims(claims string, acc wallet.Account) error {
	return nil
}
