package internal

import (
	"context"
	"fmt"
	"time"

	"berty.tech/go-orbit-db/iface"
	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/pkg/crypto/wallet"
	vaultv1 "github.com/sonrhq/core/x/identity/types/vault/v1"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
	"github.com/ucan-wg/go-ucan"
)

type IPFSStore struct {
	accConfig   *vaultv1.AccountConfig
	ipfsKVStore iface.KeyValueStore
	*empty
}

func NewIPFSStore(node common.IPFSNode, accCfg *vaultv1.AccountConfig) (wallet.Store, error) {
	docs, err := node.LoadKeyValueStore(accCfg.DID())
	if err != nil {

		return nil, err
	}
	ds := &IPFSStore{
		accConfig:   accCfg,
		ipfsKVStore: docs,
		empty:       &empty{},
	}
	return ds, nil
}

func (ds *IPFSStore) GetShare(name string) (*cmp.Config, error) {
	bz, err := ds.ipfsKVStore.Get(context.Background(), name)
	if err != nil {
		return nil, err
	}

	sc := cmp.EmptyConfig(curve.Secp256k1{})
	if err := sc.UnmarshalBinary(bz); err != nil {
		return nil, err
	}
	return sc, nil
}

func (ds *IPFSStore) SetShare(sc *cmp.Config) error {
	bz, err := sc.MarshalBinary()
	if err != nil {
		return err
	}
	_, err = ds.ipfsKVStore.Put(context.Background(), string(sc.ID), bz)
	if err != nil {
		return err
	}
	return nil
}

// JWKClaims returns the JWKClaims for the store to be signed by the identity
func (ds *IPFSStore) JWKClaims(acc wallet.Account) (string, error) {
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

func exampleParser() *ucan.TokenParser {
	caps := ucan.NewNestedCapabilities("DELEGATOR", "AUTHENTICATOR", "CREATE", "READ", "UPDATE")

	ac := func(m map[string]interface{}) (ucan.Attenuation, error) {
		var (
			cap string
			rsc ucan.Resource
		)
		for key, vali := range m {
			val, ok := vali.(string)
			if !ok {
				return ucan.Attenuation{}, fmt.Errorf(`expected attenuation value to be a string`)
			}

			if key == ucan.CapKey {
				cap = val
			} else {
				rsc = ucan.NewStringLengthResource(key, val)
			}
		}

		return ucan.Attenuation{
			Rsc: rsc,
			Cap: caps.Cap(cap),
		}, nil
	}

	store := ucan.NewMemTokenStore()
	return ucan.NewTokenParser(ac, ucan.StringDIDPubKeyResolver{}, store.(ucan.CIDBytesResolver))
}

// VerifyJWKClaims verifies the JWKClaims for the store
func (ds *IPFSStore) VerifyJWKClaims(claims string, acc wallet.Account) error {
	return nil
}
