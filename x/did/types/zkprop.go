package types

import (
	"fmt"

	"github.com/onsonr/crypto/accumulator"
	"github.com/onsonr/crypto/core/curves"
)

// Accumulator is the accumulator for the ZKP
type Accumulator []byte

// Element is the element for the BLS scheme
type Element = accumulator.Element

// Witness is the witness for the ZKP
type Witness []byte

// NewProof creates a new Proof which is used for ZKP
func NewProof(id, controller, issuer, property string, pubKey []byte) (*Proof, error) {
	input := append(pubKey, []byte(property)...)
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

	return &Proof{
		Id:          id,
		Controller:  controller,
		Issuer:      issuer,
		Property:    property,
		Accumulator: keyBytes,
	}, nil
}

// CreateAccumulator creates a new accumulator for a Proof
func CreateAccumulator(proof *Proof, values ...string) error {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	acc, err := new(accumulator.Accumulator).New(curve)
	if err != nil {
		return err
	}

	secretKey := new(accumulator.SecretKey)
	if err := secretKey.UnmarshalBinary(proof.Accumulator); err != nil {
		return err
	}

	fin, _, err := acc.Update(secretKey, ConvertValuesToZeroKnowledgeElements(values), nil)
	if err != nil {
		return err
	}

	accBytes, err := fin.MarshalBinary()
	if err != nil {
		return fmt.Errorf("failed to marshal accumulator: %w", err)
	}

	proof.Accumulator = accBytes
	return nil
}

// CreateWitness creates a witness for the accumulator in a Proof for a given value
func CreateWitness(proof *Proof, value string) ([]byte, error) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	element := curve.Scalar.Hash([]byte(value))

	secretKey := new(accumulator.SecretKey)
	if err := secretKey.UnmarshalBinary(proof.Accumulator); err != nil {
		return nil, err
	}

	accObj := new(accumulator.Accumulator)
	if err := accObj.UnmarshalBinary(proof.Accumulator); err != nil {
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
	return witnessBytes, nil
}

// VerifyWitness proves that a value is a member of the accumulator
func VerifyWitness(proof *Proof, acc Accumulator, witness Witness) error {
	secretKey := new(accumulator.SecretKey)
	if err := secretKey.UnmarshalBinary([]byte(proof.Id)); err != nil {
		return err
	}

	curve := curves.BLS12381(&curves.PointBls12381G1{})
	publicKey, err := secretKey.GetPublicKey(curve)
	if err != nil {
		return err
	}

	accObj := new(accumulator.Accumulator)
	if err := accObj.UnmarshalBinary(proof.Accumulator); err != nil {
		return fmt.Errorf("failed to unmarshal accumulator: %w", err)
	}

	witnessObj := new(accumulator.MembershipWitness)
	if err := witnessObj.UnmarshalBinary(witness); err != nil {
		return fmt.Errorf("failed to unmarshal witness: %w", err)
	}

	return witnessObj.Verify(publicKey, accObj)
}

// UpdateAccumulator updates the accumulator in a Proof with new values
func UpdateAccumulator(proof *Proof, addValues []string, removeValues []string) error {
	secretKey := new(accumulator.SecretKey)
	if err := secretKey.UnmarshalBinary(proof.Accumulator); err != nil {
		return err
	}

	accObj := new(accumulator.Accumulator)
	if err := accObj.UnmarshalBinary(proof.Accumulator); err != nil {
		return fmt.Errorf("failed to unmarshal accumulator: %w", err)
	}

	updatedAcc, _, err := accObj.Update(secretKey, ConvertValuesToZeroKnowledgeElements(addValues), ConvertValuesToZeroKnowledgeElements(removeValues))
	if err != nil {
		return err
	}

	updatedAccBytes, err := updatedAcc.MarshalBinary()
	if err != nil {
		return fmt.Errorf("failed to marshal updated accumulator: %w", err)
	}

	proof.Accumulator = updatedAccBytes
	return nil
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
