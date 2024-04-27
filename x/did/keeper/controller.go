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
	Link(key, value string) (Witness, error)
	PublicKey() *types.PublicKey
	Refresh() error
	Sign(msg []byte) ([]byte, error)
	Unlink(key, value string) (Witness, error)
	Validate(key string, w Witness) bool
	Verify(msg, sig []byte) bool
}

// controller is the controller for the DID scheme
type controller struct {
	proofKey *accumulator.PublicKey

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
	sk, err := DeriveSecretKey(c)
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
	sk, err := DeriveSecretKey(c)
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

// PublicKey returns the public key for the shares
func (c *controller) PublicKey() *types.PublicKey {
	return c.usrKS.PublicKey()
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
		return errors.Join(fmt.Errorf("Error Starting Keyshare MPC Protocol"), err)
	}
	newAlice, err := valRefresh.Result(protocol.Version1)
	if err != nil {
		return errors.Join(fmt.Errorf("Error Getting Validator Result"), err)
	}
	newBob, err := usrRefresh.Result(protocol.Version1)
	if err != nil {
		return errors.Join(fmt.Errorf("Error Getting User Result"), err)
	}
	usrKs, err := createUserKeyshare(newAlice)
	if err != nil {
		return errors.Join(fmt.Errorf("Error Creating User Keyshare"), err)
	}
	valKs, err := createValidatorKeyshare(newBob)
	if err != nil {
		return errors.Join(fmt.Errorf("Error Creating Validator Keyshare"), err)
	}
	c.usrKS = usrKs
	c.valKS = valKs
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
		return nil, errors.Join(fmt.Errorf("Error Starting Keyshare MPC Protocol"), err)
	}
	// Output
	resultMessage, err := usrSign.Result(protocol.Version1)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("Error Getting User Sign Result"), err)
	}
	sig, err := dklsv1.DecodeSignature(resultMessage)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("Error Decoding Signature"), err)
	}
	return ecdsa.SerializeSecp256k1Signature(sig)
}

// Unlink unlinks the property from the controller
func (c *controller) Unlink(key string, value string) (Witness, error) {
	propStr, ok := c.properties[key]
	if !ok {
		return "", fmt.Errorf("property not found")
	}
	sk, err := DeriveSecretKey(c)
	if err != nil {
		return "", errors.Join(err, fmt.Errorf("failed to get secret key"))
	}
	a, err := decodeProperty(Property(propStr))
	if err != nil {
		return "", errors.Join(err, fmt.Errorf("failed to decode property"))
	}
	err = a.RemoveValues(sk, value)
	if err != nil {
		return "", errors.Join(err, fmt.Errorf("failed to remove values"))
	}
	prop, witness, err := a.CreateWitness(sk, value)
	if err != nil {
		return "", errors.Join(err, fmt.Errorf("failed to create witness"))
	}
	prop.Update(key, c.properties)
	return witness, nil
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
	return c.usrKS.PublicKey().VerifySignature(msg, sig)
}

//
// 2. Utility Functions
//

// DeriveSecretKey derives the secret key from the keyshares
func DeriveSecretKey(c *controller) (*SecretKey, error) {
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
