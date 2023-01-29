package internal

import (
	"fmt"
	"sync"

	"github.com/sonrhq/core/pkg/crypto/wallet"
	vaultv1 "github.com/sonrhq/core/x/identity/types/vault/v1"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

type MemoryStore struct {
	accConfig *vaultv1.AccountConfig
	configs   map[string]*cmp.Config
	sync.Mutex
	*empty
}

func NewMemoryStore(accCfg *vaultv1.AccountConfig) (wallet.Store, error) {
	ds := &MemoryStore{
		accConfig: accCfg,
		configs:   make(map[string]*cmp.Config),
		empty:     &empty{},
	}
	return ds, nil
}

func (ds *MemoryStore) GetShare(name string) (*cmp.Config, error) {
	ds.Lock()
	defer ds.Unlock()
	s, ok := ds.configs[name]
	if !ok {
		return nil, fmt.Errorf("share not found")
	}
	return s, nil
}

func (ds *MemoryStore) SetShare(sc *cmp.Config) error {
	ds.Lock()
	defer ds.Unlock()
	ds.configs[string(sc.ID)] = sc
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
