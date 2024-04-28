package keeper

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
	Link(key, value string) (string, error)
	PublicKey() *types.PublicKey
	Refresh() error
	Sign(msg []byte) ([]byte, error)
	Unlink(key, value string) (string, error)
	Validate(key string, w string) (bool, error)
}

// controller is the controller for the DID scheme
type controller struct {
	usrKS *UserKeyshare
	valKS *ValidatorKeyshare

	properties map[string]string
}

// CreateController creates a new controller
func CreateController(uks *UserKeyshare, vks *ValidatorKeyshare) (Controller, error) {
	c := &controller{
		properties: make(map[string]string),
		usrKS:      uks,
		valKS:      vks,
	}
	return c, nil
}

// Link links a property to the controller
func (c *controller) Link(key, value string) (string, error) {
	sk, err := DeriveSecretKey(c, key)
	if err != nil {
		return "", errors.Join(err, fmt.Errorf("failed to get secret key"))
	}
	acc, err := sk.CreateAccumulator()
	if err != nil {
		return "", errors.Join(err, fmt.Errorf("failed to create accumulator"))
	}
	prop, err := zkAddValues(sk, acc, value)
	if err != nil {
		return "", errors.Join(err, fmt.Errorf("failed to add values"))
	}
	witness, err := zkCreateWitness(sk, acc, value)
	if err != nil {
		return "", errors.Join(err, fmt.Errorf("failed to create witness"))
	}
	c.properties[key] = prop
	return witness, nil
}

// PublicKey returns the public key for the shares
func (c *controller) PublicKey() *types.PublicKey {
	pub := c.valKS.PublicKey()
	return pub
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
	err = StartKsProtocol(valRefresh, usrRefresh)
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
	c.usrKS = createUserKeyshare(newAlice)
	c.valKS = createValidatorKeyshare(newBob)
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
	err = StartKsProtocol(valSign, usrSign)
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
func (c *controller) Unlink(key string, value string) (string, error) {
	acc, ok := c.properties[key]
	if !ok {
		return "", fmt.Errorf("property not found")
	}
	sk, err := DeriveSecretKey(c, key)
	if err != nil {
		return "", errors.Join(err, fmt.Errorf("failed to get secret key"))
	}
	prop, err := zkRemoveValues(sk, acc, value)
	if err != nil {
		return "", errors.Join(err, fmt.Errorf("failed to remove values"))
	}
	witness, err := zkCreateWitness(sk, acc, value)
	if err != nil {
		return "", errors.Join(err, fmt.Errorf("failed to create witness"))
	}
	c.properties[key] = prop
	return witness, nil
}

// Validate validates that the property is linked to the controller
func (c *controller) Validate(key string, w string) (bool, error) {
	acc, ok := c.properties[key]
	if !ok {
		return false, fmt.Errorf("property not found")
	}
	sk, err := DeriveSecretKey(c, key)
	if err != nil {
		return false, errors.Join(err, fmt.Errorf("failed to get secret key"))
	}
	pub, err := sk.PublicKey()
	if err != nil {
		return false, errors.Join(err, fmt.Errorf("failed to get public key"))
	}
	return zkVerifyElement(pub, acc, w)
}

//
// 2. Utility Functions
//

// DeriveSecretKey derives the secret key from the keyshares
func DeriveSecretKey(c *controller, propertyKey string) (*SecretKey, error) {
	propHash := types.Blake3Hash([]byte(propertyKey))
	seed, err := c.Sign(propHash)
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

// StartKsProtocol runs the keyshare protocol between two parties
func StartKsProtocol(firstParty protocol.Iterator, secondParty protocol.Iterator) error {
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
