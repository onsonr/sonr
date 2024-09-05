package types

import (
	"fmt"

	"github.com/onsonr/crypto/accumulator"
	"github.com/onsonr/crypto/core/curves"
)

// Element is the element for the BLS scheme
type Element = accumulator.Element

// NewProperty creates a new Property which is used for ZKP
func NewProperty(propertyKey string, pubKey []byte) (*Property, error) {
	input := append(pubKey, []byte(propertyKey)...)
	hash := []byte(input)

	curve := curves.BLS12381(&curves.PointBls12381G1{})
	key, err := new(accumulator.SecretKey).New(curve, hash[:])
	if err != nil {
		return nil, fmt.Errorf("failed to create secret key: %w", err)
	}

	keyBytes, err := key.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal secret key: %w", err)
	}

	return &Property{Key: keyBytes}, nil
}

// CreateAccumulator creates a new accumulator
func CreateAccumulator(prop *Property, values ...string) (*Accumulator, error) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	acc, err := new(accumulator.Accumulator).New(curve)
	if err != nil {
		return nil, err
	}

	secretKey := new(accumulator.SecretKey)
	if err := secretKey.UnmarshalBinary(prop.Key); err != nil {
		return nil, err
	}

	fin, _, err := acc.Update(secretKey, ConvertValuesToZeroKnowledgeElements(values), nil)
	if err != nil {
		return nil, err
	}

	accBytes, err := fin.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal accumulator: %w", err)
	}

	return &Accumulator{Accumulator: accBytes}, nil
}

// CreateWitness creates a witness for the accumulator for a given value
func CreateWitness(prop *Property, acc *Accumulator, value string) (*Witness, error) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	element := curve.Scalar.Hash([]byte(value))

	secretKey := new(accumulator.SecretKey)
	if err := secretKey.UnmarshalBinary(prop.Key); err != nil {
		return nil, err
	}

	accObj := new(accumulator.Accumulator)
	if err := accObj.UnmarshalBinary(acc.Accumulator); err != nil {
		return nil, fmt.Errorf("failed to unmarshal accumulator: %w", err)
	}

	mw, err := new(accumulator.MembershipWitness).New(element, accObj, secretKey)
	if err != nil {
		return nil, err
	}

	witnessBytes, err := mw.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal witness: %w", err)
	}

	return &Witness{Witness: witnessBytes}, nil
}

// VerifyWitness proves that a value is a member of the accumulator
func VerifyWitness(prop *Property, acc *Accumulator, witness *Witness) error {
	secretKey := new(accumulator.SecretKey)
	if err := secretKey.UnmarshalBinary(prop.Key); err != nil {
		return err
	}

	curve := curves.BLS12381(&curves.PointBls12381G1{})
	publicKey, err := secretKey.GetPublicKey(curve)
	if err != nil {
		return err
	}

	accObj := new(accumulator.Accumulator)
	if err := accObj.UnmarshalBinary(acc.Accumulator); err != nil {
		return fmt.Errorf("failed to unmarshal accumulator: %w", err)
	}

	witnessObj := new(accumulator.MembershipWitness)
	if err := witnessObj.UnmarshalBinary(witness.Witness); err != nil {
		return fmt.Errorf("failed to unmarshal witness: %w", err)
	}

	return witnessObj.Verify(publicKey, accObj)
}

// UpdateAccumulator updates the accumulator with new values
func UpdateAccumulator(prop *Property, acc *Accumulator, addValues []string, removeValues []string) (*Accumulator, error) {
	secretKey := new(accumulator.SecretKey)
	if err := secretKey.UnmarshalBinary(prop.Key); err != nil {
		return nil, err
	}

	accObj := new(accumulator.Accumulator)
	if err := accObj.UnmarshalBinary(acc.Accumulator); err != nil {
		return nil, fmt.Errorf("failed to unmarshal accumulator: %w", err)
	}

	updatedAcc, _, err := accObj.Update(secretKey, ConvertValuesToZeroKnowledgeElements(addValues), ConvertValuesToZeroKnowledgeElements(removeValues))
	if err != nil {
		return nil, err
	}

	updatedAccBytes, err := updatedAcc.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal updated accumulator: %w", err)
	}

	return &Accumulator{Accumulator: updatedAccBytes}, nil
}

// ConvertValuesToZeroKnowledgeElements converts a slice of strings to a slice of accumulator elements
func ConvertValuesToZeroKnowledgeElements(values []string) []Element {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	elements := make([]accumulator.Element, len(values))
	for i, value := range values {
		elements[i] = curve.Scalar.Hash([]byte(value))
	}
	return elements
}
