package secret

import (
	"errors"
	"fmt"

	"github.com/onsonr/hway/crypto"
	"github.com/onsonr/hway/crypto/accumulator"
	"github.com/onsonr/hway/crypto/core/curves"
)

// PrimaryKey is the secret key for the BLS scheme
type PrimaryKey struct {
	*accumulator.SecretKey
}

// Element is the element for the BLS scheme
type Element = accumulator.Element

// NewKey creates a new primary key
func NewKey(propertyKey string, pubKey crypto.PublicKey) (*PrimaryKey, error) {
	// Concatenate the controller's public key and the property key
	input := append(pubKey.Bytes(), []byte(propertyKey)...)
	hash := []byte(input)

	// Use the hash as the seed for the secret key
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	key, err := new(accumulator.SecretKey).New(curve, hash[:])
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to create secret key"))
	}
	return &PrimaryKey{SecretKey: key}, nil
}

// CreateAccumulator creates a new accumulator
func (s *PrimaryKey) CreateAccumulator(values ...string) (*accumulator.Accumulator, error) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	acc, err := new(accumulator.Accumulator).New(curve)
	if err != nil {
		return nil, err
	}
	fin, _, err := acc.Update(s.SecretKey, convertValuesToElements(values), nil)
	if err != nil {
		return nil, err
	}
	return fin, nil
}

// CreateWitness creates a witness for the accumulator for a given value
func (s *PrimaryKey) CreateWitness(acc *accumulator.Accumulator, value string) (*accumulator.MembershipWitness, error) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	element := curve.Scalar.Hash([]byte(value))
	mw, err := new(accumulator.MembershipWitness).New(element, acc, s.SecretKey)
	if err != nil {
		return nil, err
	}
	return mw, nil
}

// ProveMembership proves that a value is a member of the accumulator
func (s *PrimaryKey) VerifyWitness(acc *accumulator.Accumulator, witness *accumulator.MembershipWitness) error {
	return witness.Verify(s.PublicKey(), acc)
}

// PublicKey returns the public key for the secret key
func (s *PrimaryKey) PublicKey() *accumulator.PublicKey {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	pk, err := s.GetPublicKey(curve)
	if err != nil {
		panic(err)
	}
	return pk
}

// UpdateAccumulator updates the accumulator with new values
func (s *PrimaryKey) UpdateAccumulator(acc *accumulator.Accumulator, addValues, removeValues []string) (*accumulator.Accumulator, error) {
	acc, _, err := acc.Update(s.SecretKey, convertValuesToElements(addValues), convertValuesToElements(removeValues))
	if err != nil {
		return nil, err
	}
	return acc, nil
}

// MarshalAccumulator takes a *accumulator.Accumulator and returns a byte slice
func MarshalAccumulator(acc *accumulator.Accumulator) ([]byte, error) {
	return acc.MarshalBinary()
}

// UnmarshalAccumulator takes a byte slice and returns a *accumulator.Accumulator
func UnmarshalAccumulator(data []byte) (*accumulator.Accumulator, error) {
	acc := new(accumulator.Accumulator)
	err := acc.UnmarshalBinary(data)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func convertValuesToElements(values []string) []accumulator.Element {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	elements := []accumulator.Element{}
	for _, value := range values {
		element := curve.Scalar.Hash([]byte(value))
		elements = append(elements, element)
	}
	return elements
}
