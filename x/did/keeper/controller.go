package keeper

import (
	"github.com/di-dao/core/crypto/accumulator"
	"github.com/di-dao/core/crypto/core/curves"
)

type controller struct {
	valKS *ValidatorKeyshare
	usrKS *UserKeyshare
}

// NewSecretKey creates a new secret key
func (c *controller) DeriveSecretKey(seed []byte) (*SecretKey, error) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	key, err := new(accumulator.SecretKey).New(curve, seed[:])
	if err != nil {
		return nil, err
	}
	return &SecretKey{SecretKey: key}, nil
}
