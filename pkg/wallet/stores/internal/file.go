package internal

import (
	"context"
	"time"

	"github.com/sonrhq/core/pkg/wallet"
	"github.com/sonrhq/core/pkg/wallet/accounts"
	vaultv1 "github.com/sonrhq/core/x/identity/types/vault/v1"
	"github.com/ucan-wg/go-ucan"
	bolt "go.etcd.io/bbolt"
)

type FileStore struct {
	accConfig *vaultv1.AccountConfig
	path      string
	db        *bolt.DB
	bucketKey []byte
	pwd       []byte
}

func NewFileStore(p string, pwd []byte, accCfg *vaultv1.AccountConfig) (wallet.Store, error) {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(p, 0600, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	ds := &FileStore{
		accConfig: accCfg,
		path:      p,
		db:        db,
		bucketKey: []byte(accCfg.DID()),
		pwd:       pwd,
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

func (ds *FileStore) GetAccount(name string) (wallet.Account, error) {
	var acc wallet.Account
	ds.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(ds.bucketKey)
		v := b.Get([]byte(name))
		w, err := accounts.LoadFromBytes(v)
		if err != nil {
			return err
		}
		acc = w
		return nil
	})
	return acc, nil
}

func (ds *FileStore) PutAccount(w wallet.Account, name string) error {
	ds.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(ds.bucketKey)
		v, err := w.Marshal()
		if err != nil {
			return err
		}
		return b.Put([]byte(name), v)
	})
	return nil
}

// JWKClaims returns the JWKClaims for the store to be signed by the identity
func (ds *FileStore) JWKClaims(acc wallet.Account) (string, error) {
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
func (ds *FileStore) VerifyJWKClaims(claims string, acc wallet.Account) error {
	p := exampleParser()
	_, err := p.ParseAndVerify(context.Background(), claims)
	if err != nil {
		return err
	}
	return nil
}
