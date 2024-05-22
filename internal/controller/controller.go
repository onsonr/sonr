package controller

import (
	"errors"
	"fmt"

	"github.com/di-dao/core/crypto"
	"github.com/di-dao/core/crypto/accumulator"
	"github.com/di-dao/core/crypto/core/protocol"
	"github.com/di-dao/core/crypto/kss"
	"github.com/di-dao/core/crypto/mpc"
	"github.com/di-dao/core/crypto/signatures/ecdsa"
	"github.com/di-dao/core/crypto/tecdsa/dklsv1"
	"github.com/di-dao/core/crypto/zk"
	"github.com/di-dao/core/x/did/types"
)

// controller is the controller for the DID scheme
type controller struct {
	usrKs      kss.User
	valKs      kss.Val
	properties zk.Properties
}

// Create creates a new controller
func Create(kss kss.Set) (types.ControllerI, error) {
	c := &controller{
		properties: make(map[string]*accumulator.Accumulator),
		usrKs:      kss.Usr(),
		valKs:      kss.Val(),
	}
	return c, nil
}

// PublicKey returns the public key for the shares
func (c *controller) PublicKey() crypto.PublicKey {
	pub := c.valKs.PublicKey()
	return pub
}

// Refresh refreshes the keyshares
func (c *controller) Refresh() error {
	valRefresh, err := c.valKs.GetRefreshFunc()
	if err != nil {
		return errors.Join(err, fmt.Errorf("failed to get validator refresh function"))
	}
	usrRefresh, err := c.usrKs.GetRefreshFunc()
	if err != nil {
		return errors.Join(err, fmt.Errorf("failed to get user refresh function"))
	}
	err = mpc.RunProtocol(valRefresh, usrRefresh)
	if err != nil {
		return errors.Join(fmt.Errorf("error Starting Keyshare MPC Protocol"), err)
	}
	return nil
}

// Sign signs the message with the keyshares
func (c *controller) Sign(msg []byte) ([]byte, error) {
	valSign, err := c.valKs.GetSignFunc(msg)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to get validator sign function"))
	}
	usrSign, err := c.usrKs.GetSignFunc(msg)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to get user sign function"))
	}
	err = mpc.RunProtocol(valSign, usrSign)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("error Starting Keyshare MPC Protocol"), err)
	}
	// Output
	resultMessage, err := usrSign.Result(protocol.Version1)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("error Getting User Sign Result"), err)
	}
	sig, err := dklsv1.DecodeSignature(resultMessage)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("error Decoding Signature"), err)
	}
	return ecdsa.SerializeSecp256k1Signature(sig)
}

// Validate validates the witness
func (c *controller) Check(key string, witness []byte) bool {
	sk, err := zk.DeriveSecretKey(key, c.PublicKey())
	if err != nil {
		return false
	}
	acc, ok := c.properties[key]
	if !ok {
		return false
	}
	wit := &accumulator.MembershipWitness{}
	err = wit.UnmarshalBinary(witness)
	if err != nil {
		return false
	}
	return sk.VerifyWitness(acc, wit) == nil
}

// Set sets the property for the controller
func (c *controller) Set(key string, value string) ([]byte, error) {
	sk, err := zk.DeriveSecretKey(key, c.PublicKey())
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to get secret key"))
	}
	acc, err := sk.CreateAccumulator(value)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to create accumulator"))
	}
	c.properties[key] = acc
	witness, err := sk.CreateWitness(acc, value)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to create witness"))
	}
	return witness.MarshalBinary()
}

// Unlink unlinks the property from the controller
func (c *controller) Remove(key string, value string) error {
	sk, err := zk.DeriveSecretKey(key, c.PublicKey())
	if err != nil {
		return err
	}
	acc, ok := c.properties[key]
	if !ok {
		return fmt.Errorf("property not found")
	}
	witness, err := sk.CreateWitness(acc, value)
	if err != nil {
		return err
	}
	// no need to continue if the property is not linked
	err = sk.VerifyWitness(acc, witness)
	if err != nil {
		return nil
	}
	newAcc, err := sk.UpdateAccumulator(acc, []string{}, []string{value})
	if err != nil {
		return err
	}
	c.properties[key] = newAcc
	return nil
}
