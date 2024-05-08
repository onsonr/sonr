package controller

import (
	"errors"
	"fmt"

	"github.com/di-dao/core/crypto/accumulator"
	"github.com/di-dao/core/crypto/core/curves"
	"lukechampine.com/blake3"
)

type Properties map[string]*accumulator.Accumulator

func ConvertByteMapToProperties(p map[string][]byte) (Properties, error) {
	results := make(Properties, len(p))
	for k, v := range p {
		acc := &accumulator.Accumulator{}
		err := acc.UnmarshalBinary(v)
		if err != nil {
			return nil, err
		}
		results[k] = acc
	}
	return results, nil
}

func ConvertPropertiesToByteMap(p Properties) (map[string][]byte, error) {
	results := make(map[string][]byte, len(p))
	for k, v := range p {
		b, err := v.MarshalBinary()
		if err != nil {
			return nil, err
		}
		results[k] = b
	}
	return results, nil
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

// deriveSecretKey derives the secret key from the keyshares
func (c *controller) deriveSecretKey(propertyKey string) (*SecretKey, error) {
	// Get the controller's public key
	controllerPubKey := c.PublicKey()

	// Concatenate the controller's public key and the property key
	input := append(controllerPubKey.Bytes(), []byte(propertyKey)...)
	hash := blake3Hash(input)

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

// Blake3Hash returns the blake3 hash of the input bytes
func blake3Hash(bz []byte) []byte {
	bz32 := blake3.Sum256(bz)
	return bz32[:]
}
