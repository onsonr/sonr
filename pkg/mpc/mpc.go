package mpc

import (
	"github.com/sonr-io/core/pkg/crypto"
	"github.com/sonr-io/core/pkg/mpc/base"
	models "github.com/sonr-io/core/pkg/mpc/base"
	v1algo "github.com/sonr-io/core/pkg/mpc/protocol/dkls"
	v1types "github.com/sonr-io/core/pkg/mpc/types/v1"
)

// AccountV1 is a type alias for the AccountV1 struct in the base package.
type AccountV1 = base.AccountV1

// KeyshareV1 is a type alias for the Keyshare struct in the v1types package.
type KeyshareV1 = v1types.Keyshare

// KeyshareSet is a type alias for the KeyshareSet struct in the v1types package.
type KeyshareSet = v1types.KeyshareSet

// EncKeyshareSet is a type alias for the EncKeyshareSet struct in the v1types package.
type EncKeyshareSet = v1types.EncKeyshareSet

// ZKSet is a type alias for the ZKSet struct in the v1types package.
type ZKSet = v1types.ZKSet

// GenerateV2 generates a new account with a given ID.
func GenerateV2(name string, ct crypto.CoinType) (*models.AccountV1, KeyshareSet, error) {
	return models.NewAccountV1(name, ct)
}

// KeygenV1 generates a keyshare set.
func KeygenV1() (KeyshareSet, error) {
	kss, err := v1algo.DKLSKeygen()
	if err != nil {
		return v1types.EmptyKeyshareSet(), err
	}
	return kss, nil
}

// NewKSS creates a new keyshare set from a list of keyshares.
func NewKSS(pub *KeyshareV1, priv *KeyshareV1) KeyshareSet {
	return v1types.NewKSS(pub, priv)
}

// NewZKSet creates a new zero-knowledge set from a list of zero-knowledge proofs.
func NewZKSet(pubKey *crypto.Secp256k1PubKey, initialIds ...string) (ZKSet, error) {
	return v1types.CreateZkSet(pubKey, initialIds...)
}

// LoadZKSet loads a zero-knowledge set from a list of zero-knowledge proofs.
func LoadZKSet(str string) (ZKSet, error) {
	return v1types.OpenZkSet(str)
}
