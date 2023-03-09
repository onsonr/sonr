package internal

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sonrhq/core/pkg/wallet"
	"github.com/sonrhq/core/pkg/wallet/accounts"
	vaultv1 "github.com/sonrhq/core/x/identity/types/vault/v1"
	"github.com/ucan-wg/go-ucan"
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
	acc, err := accounts.Load(accCfg)
	if err != nil {
		return nil, err
	}
	err = ds.PutAccount(acc, accCfg.DID())
	if err != nil {
		return nil, err
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
	caps := ucan.NewNestedCapabilities("DELEGATOR", "AUTHENTICATOR", "CREATE", "READ", "UPDATE")
	att := ucan.Attenuations{
		{Cap: caps.Cap("AUTHENTICATOR"), Rsc: ucan.NewStringLengthResource("mpc/acc", "*")},
		{Cap: caps.Cap("SUPER_USER"), Rsc: ucan.NewStringLengthResource("mpc/acc", "b5:world_bank_population:*")},
	}
	zero := time.Time{}
	origin, err := acc.NewOriginToken(acc.PubKey().DID(), att, nil, zero, zero)
	if err != nil {
		return "", err
	}
	return origin, nil
}

// VerifyJWKClaims verifies the JWKClaims for the store
func (ds *MemoryStore) VerifyJWKClaims(claims string, acc wallet.Account) error {
	p := exampleParser()
	_, err := p.ParseAndVerify(context.Background(), claims)
	if err != nil {
		return err
	}
	return nil
}
