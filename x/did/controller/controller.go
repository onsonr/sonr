package controller

import (
	"errors"
	"fmt"

	"github.com/di-dao/core/crypto/accumulator"
	"github.com/di-dao/core/crypto/core/curves"
	"github.com/di-dao/core/crypto/core/protocol"
	"github.com/di-dao/core/crypto/signatures/ecdsa"
	"github.com/di-dao/core/crypto/tecdsa/dklsv1"
	"github.com/di-dao/core/x/did/types"
)

// Controller is the interface for the controller
type Controller interface {
	Set(key, value string) ([]byte, error)
	PublicKey() *types.PublicKey
	Refresh() error
	Sign(msg []byte) ([]byte, error)
	Remove(key, value string) error
	Check(key string, w []byte) bool
}

// controller is the controller for the DID scheme
type controller struct {
	usrKs      types.UserKeyshare
	valKs      types.ValidatorKeyshare
	props      []*types.Property
	properties map[string]*accumulator.Accumulator
}

// Create creates a new controller
func Create(kss types.KeyshareSet) (Controller, error) {
	c := &controller{
		properties: make(map[string]*accumulator.Accumulator),
		usrKs:      kss.Usr(),
		valKs:      kss.Val(),
	}
	return c, nil
}
func (c *controller) Set(key string, value string) ([]byte, error) {
	sk, err := c.deriveSecretKey(key)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to get secret key"))
	}
	acc, err := sk.CreateAccumulator(value)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to create accumulator"))
	}
	c.setAccumulator(key, acc)
	witness, err := sk.CreateWitness(acc, value)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to create witness"))
	}
	return witness.MarshalBinary()
}

// PublicKey returns the public key for the shares
func (c *controller) PublicKey() *types.PublicKey {
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
	err = runMpcProtocol(valRefresh, usrRefresh)
	if err != nil {
		return errors.Join(fmt.Errorf("error Starting Keyshare MPC Protocol"), err)
	}
	newAlice, err := valRefresh.Result(protocol.Version1)
	if err != nil {
		return errors.Join(fmt.Errorf("error Getting Validator Result"), err)
	}
	newBob, err := usrRefresh.Result(protocol.Version1)
	if err != nil {
		return errors.Join(fmt.Errorf("error Getting User Result"), err)
	}
	kss := types.NewKeyshareSet(newAlice, newBob)
	c.valKs = kss.Val()
	c.usrKs = kss.Usr()
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
	err = runMpcProtocol(valSign, usrSign)
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

// Unlink unlinks the property from the controller
func (c *controller) Remove(key string, value string) error {
	sk, err := c.deriveSecretKey(key)
	if err != nil {
		return err
	}
	acc, err := c.getAccumulator(key)
	if err != nil {
		return err
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
	return c.setAccumulator(key, newAcc)
}

// Validate validates the witness
func (c *controller) Check(key string, witness []byte) bool {
	sk, err := c.deriveSecretKey(key)
	if err != nil {
		return false
	}
	acc, err := c.getAccumulator(key)
	if err != nil {
		return false
	}
	wit := &accumulator.MembershipWitness{}
	err = wit.UnmarshalBinary(witness)
	if err != nil {
		return false
	}
	return sk.VerifyWitness(acc, wit) == nil
}

// deriveSecretKey derives the secret key from the keyshares
func (c *controller) deriveSecretKey(propertyKey string) (*SecretKey, error) {
	// Get the controller's public key
	controllerPubKey := c.PublicKey()

	// Concatenate the controller's public key and the property key
	input := append(controllerPubKey.Bytes(), []byte(propertyKey)...)
	hash := types.Blake3Hash(input)

	// Use the hash as the seed for the secret key
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	key, err := new(accumulator.SecretKey).New(curve, hash[:])
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to create secret key"))
	}
	return &SecretKey{SecretKey: key}, nil
}

func (c *controller) getAccumulator(key string) (*accumulator.Accumulator, error) {
	acc, ok := c.properties[key]
	if !ok {
		return nil, fmt.Errorf("property not found")
	}
	return acc, nil
}

func (c *controller) setAccumulator(key string, acc *accumulator.Accumulator) error {
	c.properties[key] = acc
	return nil
}

//
// 3. Utility Functions
//

// runMpcProtocol runs the keyshare protocol between two parties
func runMpcProtocol(firstParty protocol.Iterator, secondParty protocol.Iterator) error {
	var (
		message *protocol.Message
		aErr    error
		bErr    error
	)

	for aErr != protocol.ErrProtocolFinished || bErr != protocol.ErrProtocolFinished {
		// Crank each protocol forward one iteration
		message, bErr = firstParty.Next(message)
		if bErr != nil && bErr != protocol.ErrProtocolFinished {
			return errors.Join(fmt.Errorf("validator failed to process mpc message"), bErr)
		}

		message, aErr = secondParty.Next(message)
		if aErr != nil && aErr != protocol.ErrProtocolFinished {
			return errors.Join(fmt.Errorf("user failed to process mpc message"), aErr)
		}
	}
	if aErr == protocol.ErrProtocolFinished && bErr == protocol.ErrProtocolFinished {
		return nil
	}
	if aErr != nil && bErr != nil {
		return fmt.Errorf("both parties failed: %v, %v", aErr, bErr)
	}
	if aErr != nil {
		return fmt.Errorf("validator keyshare failed: %v", aErr)
	}
	if bErr != nil {
		return fmt.Errorf("user keyshare failed: %v", bErr)
	}
	return nil
}
