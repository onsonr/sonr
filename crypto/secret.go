package crypto

import (
	"github.com/di-dao/sonr/crypto/accumulator"
	"github.com/di-dao/sonr/crypto/core/curves"
)

// SecretKey is the secret key for the BLS scheme
type SecretKey struct {
	*accumulator.SecretKey
}

// Element is the element for the BLS scheme
type Element = accumulator.Element

// CreateAccumulator creates a new accumulator
func (s *SecretKey) CreateAccumulator(values ...string) (*accumulator.Accumulator, error) {
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
func (s *SecretKey) CreateWitness(acc *accumulator.Accumulator, value string) (*accumulator.MembershipWitness, error) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	element := curve.Scalar.Hash([]byte(value))
	mw, err := new(accumulator.MembershipWitness).New(element, acc, s.SecretKey)
	if err != nil {
		return nil, err
	}
	return mw, nil
}

// ProveMembership proves that a value is a member of the accumulator
func (s *SecretKey) VerifyWitness(acc *accumulator.Accumulator, witness *accumulator.MembershipWitness) error {
	return witness.Verify(s.PublicKey(), acc)
}

// PublicKey returns the public key for the secret key
func (s *SecretKey) PublicKey() *accumulator.PublicKey {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	pk, err := s.GetPublicKey(curve)
	if err != nil {
		panic(err)
	}
	return pk
}

// UpdateAccumulator updates the accumulator with new values
func (s *SecretKey) UpdateAccumulator(acc *accumulator.Accumulator, addValues, removeValues []string) (*accumulator.Accumulator, error) {
	acc, _, err := acc.Update(s.SecretKey, convertValuesToElements(addValues), convertValuesToElements(removeValues))
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
