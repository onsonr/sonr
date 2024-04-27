package keeper

import (
	"encoding/hex"

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

// Property is the property for the BLS scheme
type Property string

// String returns the string representation of the property
func (p Property) String() string {
	return string(p)
}

// Update updates the property with the given value
func (p Property) Update(key string, properties map[string]string) {
	properties[key] = p.String()
}

// Witness is the witness for the BLS scheme
type Witness string

// String returns the string representation of the witness
func (w Witness) String() string {
	return string(w)
}

// CreateAccumulator creates a new accumulator
func (s *SecretKey) CreateAccumulator() (*Accumulator, error) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	acc, err := new(accumulator.Accumulator).New(curve)
	if err != nil {
		return nil, err
	}
	return &Accumulator{Accumulator: acc}, nil
}

// DeserializeAccumulator opens an accumulator
func DeserializeAccumulator(hexAcc string) (*Accumulator, error) {
	acc, err := hex.DecodeString(hexAcc)
	if err != nil {
		return nil, err
	}
	e := new(accumulator.Accumulator)
	err = e.UnmarshalBinary(acc)
	if err != nil {
		return nil, err
	}
	return &Accumulator{Accumulator: e}, nil
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

// Accumulator is the secret key for the BLS scheme
type Accumulator struct {
	*accumulator.Accumulator
}

// AddValue adds a value to the accumulator
func (a *Accumulator) AddValues(k *SecretKey, values ...string) error {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	elements := []accumulator.Element{}
	for _, value := range values {
		element := curve.Scalar.Hash([]byte(value))
		elements = append(elements, element)
	}

	acc, _, err := a.Accumulator.Update(k.SecretKey, elements, nil)
	if err != nil {
		return err
	}
	a.Accumulator = acc
	return nil
}

// RemoveValue removes a value from the accumulator
func (a *Accumulator) RemoveValues(k *SecretKey, values ...string) error {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	elements := []accumulator.Element{}
	for _, value := range values {
		element := curve.Scalar.Hash([]byte(value))
		elements = append(elements, element)
	}
	acc, _, err := a.Accumulator.Update(k.SecretKey, nil, elements)
	if err != nil {
		return err
	}
	a.Accumulator = acc
	return nil
}

// CreateWitness creates a witness for the accumulator for a given value
func (a *Accumulator) CreateWitness(k *SecretKey, value string) (Property, Witness, error) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	element := curve.Scalar.Hash([]byte(value))
	mw, err := new(accumulator.MembershipWitness).New(element, a.Accumulator, k.SecretKey)
	if err != nil {
		return "", "", err
	}
	wtstr, err := encodeWitness(mw)
	if err != nil {
		return "", "", err
	}
	prop, err := encodeProperty(a)
	if err != nil {
		return "", "", err
	}
	return prop, wtstr, nil
}

// VerifyElement verifies an element against the accumulator and public key
func (a *Accumulator) VerifyElement(pk *accumulator.PublicKey, witness Witness) bool {
	mw, err := decodeWitness(witness)
	if err != nil {
		return false
	}
	err = mw.Verify(pk, a.Accumulator)
	return err == nil
}

// encoodeProperty encodes the accumulator to a base58 string
func encodeProperty(acc *Accumulator) (Property, error) {
	bz, err := acc.MarshalBinary()
	if err != nil {
		return "", err
	}
	return Property(base58.Encode(bz)), nil
}

// encodeWitness encodes the witness to a base58 string
func encodeWitness(witness *accumulator.MembershipWitness) (Witness, error) {
	witBz, err := witness.MarshalBinary()
	if err != nil {
		return "", err
	}
	return Witness(base58.Encode(witBz)), nil
}

// decodeProperty decodes the accumulator from a base58 string
func decodeProperty(prop Property) (*Accumulator, error) {
	accBz, err := hex.DecodeString(prop.String())
	if err != nil {
		return nil, err
	}
	e := new(accumulator.Accumulator)
	err = e.UnmarshalBinary(accBz)
	if err != nil {
		return nil, err
	}
	return &Accumulator{Accumulator: e}, nil
}

// decodeWitness decodes the witness from a base58 string
func decodeWitness(witness Witness) (*accumulator.MembershipWitness, error) {
	bz, err := base58.Decode(witness.String())
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
