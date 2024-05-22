package controller

import (
	"errors"
	"fmt"

	"github.com/di-dao/sonr/crypto"
	"github.com/di-dao/sonr/crypto/core/protocol"
	"github.com/di-dao/sonr/crypto/kss"
	"github.com/di-dao/sonr/crypto/mpc"
	"github.com/di-dao/sonr/crypto/signatures/ecdsa"
	"github.com/di-dao/sonr/crypto/tecdsa/dklsv1"
)

// Controller is the interface for the controller
type Controller interface {
	PublicKey() crypto.PublicKey
	Refresh() error
	Sign(msg []byte) ([]byte, error)
}

// controller is the controller for the DID scheme
type controller struct {
	usrKs kss.User
	valKs kss.Val
}

// New creates a new controller
func New(kss kss.Set) Controller {
	c := &controller{
		usrKs: kss.Usr(),
		valKs: kss.Val(),
	}
	return c
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
