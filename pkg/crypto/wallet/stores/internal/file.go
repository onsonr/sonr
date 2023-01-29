package internal

import (
	"github.com/sonrhq/core/pkg/crypto/wallet"
	vaultv1 "github.com/sonrhq/core/x/identity/types/vault/v1"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
	bolt "go.etcd.io/bbolt"
)

type FileStore struct {
	accConfig *vaultv1.AccountConfig
	path      string
	db        *bolt.DB
	bucketKey []byte
	*empty
}

func NewFileStore(p string, accCfg *vaultv1.AccountConfig) (wallet.Store, error) {
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
		empty:     &empty{},
		bucketKey: []byte(accCfg.DID()),
	}
	return ds, nil
}

func (ds *FileStore) GetShare(name string) (*cmp.Config, error) {
	sc := cmp.EmptyConfig(curve.Secp256k1{})
	ds.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(ds.bucketKey)
		v := b.Get([]byte(name))
		return sc.UnmarshalBinary(v)
	})
	return sc, nil
}

func (ds *FileStore) SetShare(sc *cmp.Config) error {
	ds.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(ds.bucketKey)
		v, err := sc.MarshalBinary()
		if err != nil {
			return err
		}
		return b.Put([]byte(sc.ID), v)
	})
	return nil
}

// JWKClaims returns the JWKClaims for the store to be signed by the identity
func (ds *FileStore) JWKClaims(acc wallet.Account) (string, error) {
	return "", nil
}

// VerifyJWKClaims verifies the JWKClaims for the store
func (ds *FileStore) VerifyJWKClaims(claims string, acc wallet.Account) error {
	return nil
}
