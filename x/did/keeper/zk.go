package keeper

import (
	"github.com/mr-tron/base58"

	"github.com/di-dao/core/crypto/accumulator"
	"github.com/di-dao/core/crypto/core/curves"
)

// SecretKey is the secret key for the BLS scheme
type SecretKey struct {
	*accumulator.SecretKey
}

// Element is the element for the BLS scheme
type Element = accumulator.Element

// CreateAccumulator creates a new accumulator
func (s *SecretKey) CreateAccumulator() (string, error) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	acc, err := new(accumulator.Accumulator).New(curve)
	if err != nil {
		return "", err
	}
	return encodeAccumulator(acc), nil
}

// DeserializeAccumulator opens an accumulator
func DeserializeAccumulator(encodedAcc string) (*accumulator.Accumulator, error) {
	return decodeAccumulator(encodedAcc)
}

// PublicKey returns the public key for the secret key
func (s *SecretKey) PublicKey() (*accumulator.PublicKey, error) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	pk, err := s.SecretKey.GetPublicKey(curve)
	if err != nil {
		return nil, err
	}
	return pk, nil
}

// AddValues adds values to the accumulator
func zkAddValues(k *SecretKey, encodedAcc string, values ...string) (string, error) {
	acc, err := decodeAccumulator(encodedAcc)
	if err != nil {
		return "", err
	}

	curve := curves.BLS12381(&curves.PointBls12381G1{})
	elements := []accumulator.Element{}
	for _, value := range values {
		element := curve.Scalar.Hash([]byte(value))
		elements = append(elements, element)
	}

	updatedAcc, _, err := acc.Update(k.SecretKey, elements, nil)
	if err != nil {
		return "", err
	}
	return encodeAccumulator(updatedAcc), nil
}

// RemoveValues removes values from the accumulator
func zkRemoveValues(k *SecretKey, encodedAcc string, values ...string) (string, error) {
	acc, err := decodeAccumulator(encodedAcc)
	if err != nil {
		return "", err
	}

	curve := curves.BLS12381(&curves.PointBls12381G1{})
	elements := []accumulator.Element{}
	for _, value := range values {
		element := curve.Scalar.Hash([]byte(value))
		elements = append(elements, element)
	}

	updatedAcc, _, err := acc.Update(k.SecretKey, nil, elements)
	if err != nil {
		return "", err
	}
	return encodeAccumulator(updatedAcc), nil
}

// CreateWitness creates a witness for the accumulator for a given value
func zkCreateWitness(k *SecretKey, encodedAcc string, value string) (string, error) {
	acc, err := decodeAccumulator(encodedAcc)
	if err != nil {
		return "", err
	}

	curve := curves.BLS12381(&curves.PointBls12381G1{})
	element := curve.Scalar.Hash([]byte(value))
	mw, err := new(accumulator.MembershipWitness).New(element, acc, k.SecretKey)
	if err != nil {
		return "", err
	}
	return encodeWitness(mw), nil
}

// VerifyElement verifies an element against the accumulator and public key
func zkVerifyElement(pk *accumulator.PublicKey, encodedAcc string, witness string) (bool, error) {
	acc, err := decodeAccumulator(encodedAcc)
	if err != nil {
		return false, err
	}

	mw, err := decodeWitness(witness)
	if err != nil {
		return false, err
	}
	err = mw.Verify(pk, acc)
	return err == nil, err
}

// encodeAccumulator encodes the accumulator to a base58 string
func encodeAccumulator(acc *accumulator.Accumulator) string {
	bz, _ := acc.MarshalBinary()
	return base58.Encode(bz)
}

// encodeWitness encodes the witness to a base58 string
func encodeWitness(witness *accumulator.MembershipWitness) string {
	witBz, _ := witness.MarshalBinary()
	return base58.Encode(witBz)
}

// decodeAccumulator decodes the accumulator from a base58 string
func decodeAccumulator(encodedAcc string) (*accumulator.Accumulator, error) {
	accBz, err := base58.Decode(encodedAcc)
	if err != nil {
		return nil, err
	}
	acc := new(accumulator.Accumulator)
	err = acc.UnmarshalBinary(accBz)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

// decodeWitness decodes the witness from a base58 string
func decodeWitness(encodedWitness string) (*accumulator.MembershipWitness, error) {
	bz, err := base58.Decode(encodedWitness)
	if err != nil {
		return nil, err
	}
	mw := new(accumulator.MembershipWitness)
	err = mw.UnmarshalBinary(bz)
	if err != nil {
		return nil, err
	}
	return mw, nil
}
