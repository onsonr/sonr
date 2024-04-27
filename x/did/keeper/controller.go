package keeper

import (
	"errors"
	"fmt"

	"github.com/di-dao/core/crypto/accumulator"
	"github.com/di-dao/core/crypto/core/curves"
	"github.com/di-dao/core/x/did/types"
)

type controller struct {
	valKS *ValidatorKeyshare
	usrKS *UserKeyshare
}

// Refresh refreshes the keyshares
func (c *controller) Refresh() error {
	valRefresh, err := c.valKS.GetRefreshFunc()
	if err != nil {
		return errors.Join(err, fmt.Errorf("failed to get validator refresh function"))
	}
	usrRefresh, err := c.usrKS.GetRefreshFunc()
	if err != nil {
		return errors.Join(err, fmt.Errorf("failed to get user refresh function"))
	}
	newVal, newUsr, err := refreshKSS(valRefresh, usrRefresh)
	if err != nil {
		return errors.Join(err, fmt.Errorf("failed to refresh keyshares"))
	}
	c.valKS = newVal
	c.usrKS = newUsr
	return nil
}

// NewSecretKey creates a new secret key
func (c *controller) SecretKey() (*SecretKey, error) {
	seed, err := c.getAnonSeed()
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to get anon seed"))
	}
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	key, err := new(accumulator.SecretKey).New(curve, seed[:])
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to create secret key"))
	}
	return &SecretKey{SecretKey: key}, nil
}

// Sign signs the message with the keyshares
func (c *controller) Sign(msg []byte) ([]byte, error) {
	valSign, err := c.valKS.GetSignFunc(msg)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to get validator sign function"))
	}
	usrSign, err := c.usrKS.GetSignFunc(msg)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to get user sign function"))
	}
	return signKSS(valSign, usrSign)
}

// Verify verifies the signature
func (c *controller) Verify(msg, sig []byte) bool {
	return c.valKS.PublicKey().VerifySignature(msg, sig)
}

// util function for getting the anonymous seed
func (c *controller) getAnonSeed() ([]byte, error) {
	addr, err := types.GetIDXAddress(c.usrKS.pubKey)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to get IDX address"))
	}
	return c.Sign(addr.Bytes())
}
