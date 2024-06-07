package types

import (
	"errors"
	"fmt"

	"github.com/di-dao/sonr/crypto"
	"github.com/di-dao/sonr/crypto/accumulator"
	"github.com/di-dao/sonr/crypto/core/curves"
)

type Properties map[string]*accumulator.Accumulator

func NewProperties() Properties {
	return make(Properties)
}

// Check validates the witness
func (p Properties) Check(publicKey crypto.PublicKey, key string, witness []byte) bool {
	sk, err := DeriveSecretKey(key, publicKey)
	if err != nil {
		return false
	}
	acc, ok := p[key]
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
func (p Properties) Set(publicKey crypto.PublicKey, key, value string) ([]byte, error) {
	sk, err := DeriveSecretKey(key, publicKey)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to get secret key"))
	}
	acc, err := sk.CreateAccumulator(value)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to create accumulator"))
	}
	p[key] = acc
	witness, err := sk.CreateWitness(acc, value)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to create witness"))
	}
	return witness.MarshalBinary()
}

// Remove unlinks the property from the controller
func (p Properties) Remove(publicKey crypto.PublicKey, key, value string) error {
	sk, err := DeriveSecretKey(key, publicKey)
	if err != nil {
		return err
	}
	acc, ok := p[key]
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
	p[key] = newAcc
	return nil
}

// Marshal the properties to a byte map
func (p Properties) Marshal() (map[string][]byte, error) {
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

// Unmarshal the properties from a byte map
func (p Properties) Unmarshal(m map[string][]byte) error {
	for k, v := range m {
		acc := &accumulator.Accumulator{}
		err := acc.UnmarshalBinary(v)
		if err != nil {
			return err
		}
		p[k] = acc
	}
	return nil
}

// deriveSecretKey derives the secret key from the keyshares
func DeriveSecretKey(propertyKey string, pubKey crypto.PublicKey) (*crypto.SecretKey, error) {
	// Concatenate the controller's public key and the property key
	input := append(pubKey.Bytes(), []byte(propertyKey)...)
	hash := []byte(input)

	// Use the hash as the seed for the secret key
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	key, err := new(accumulator.SecretKey).New(curve, hash[:])
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to create secret key"))
	}
	return &crypto.SecretKey{SecretKey: key}, nil
}
