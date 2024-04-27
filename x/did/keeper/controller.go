package keeper

import (
	"errors"
	"fmt"

	"github.com/di-dao/core/crypto/accumulator"
	"github.com/di-dao/core/crypto/core/curves"
	"github.com/di-dao/core/x/did/types"
)

type Controller interface {
	Link(key, value string) (Witness, error)
	Refresh() error
	Sign(msg []byte) ([]byte, error)
	Validate(key string, w Witness) bool
	Verify(msg, sig []byte) bool
}

type controller struct {
	proofKey *accumulator.PublicKey

	usrKS *UserKeyshare
	valKS *ValidatorKeyshare

	properties map[string]string
}

func CreateController(uks *UserKeyshare, vks *ValidatorKeyshare) (Controller, error) {
	c := &controller{
		properties: make(map[string]string),
		usrKS:      uks,
		valKS:      vks,
	}
	sk, err := deriveSecretKey(c)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to get secret key"))
	}
	prfKey, err := sk.PublicKey()
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to get public key"))
	}
	c.proofKey = prfKey
	return c, nil
}

// Link links a property to the controller
func (c *controller) Link(key, value string) (Witness, error) {
	sk, err := deriveSecretKey(c)
	if err != nil {
		return "", errors.Join(err, fmt.Errorf("failed to get secret key"))
	}
	acc, err := sk.CreateAccumulator()
	if err != nil {
		return "", errors.Join(err, fmt.Errorf("failed to create accumulator"))
	}
	err = acc.AddValues(sk, value)
	if err != nil {
		return "", errors.Join(err, fmt.Errorf("failed to add values"))
	}
	prop, witness, err := acc.CreateWitness(sk, value)
	if err != nil {
		return "", errors.Join(err, fmt.Errorf("failed to create witness"))
	}
	prop.Update(key, c.properties)
	return witness, nil
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

	c.properties = make(map[string]string)
	c.valKS = newVal
	c.usrKS = newUsr
	return nil
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

// Validate validates that the property is linked to the controller
func (c *controller) Validate(key string, w Witness) bool {
	propStr, ok := c.properties[key]
	if !ok {
		return false
	}
	a, err := decodeProperty(Property(propStr))
	if err != nil {
		return false
	}
	mw, err := decodeWitness(w)
	if err != nil {
		return false
	}
	err = mw.Verify(c.proofKey, a.Accumulator)
	if err != nil {
		return false
	}
	return true
}

// Verify verifies the signature
func (c *controller) Verify(msg, sig []byte) bool {
	return c.valKS.PublicKey().VerifySignature(msg, sig)
}

// deriveSecretKey derives the secret key from the keyshares
func deriveSecretKey(c *controller) (*SecretKey, error) {
	addr, err := types.GetIDXAddress(c.usrKS.pubKey)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to get IDX address"))
	}
	seed, err := c.Sign(addr.Bytes())
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
