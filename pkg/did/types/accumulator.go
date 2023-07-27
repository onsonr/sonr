package types

import (
	"fmt"
	"strings"

	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/kryptology/pkg/accumulator"
	"github.com/sonrhq/kryptology/pkg/core/curves"
)

// DIDAccumulator is a ZKSet accumulator for a DID
type DIDAccumulator string

// NewAccumulator creates a new accumulator for the DID
func NewAccumulator(id DIDIdentifier, key string) DIDAccumulator {
	return DIDAccumulator(fmt.Sprintf("%s!%s", id.String(), key))
}

// Key returns the key of the resource
func (d DIDAccumulator) Key() string {
	ptrs := strings.Split(string(d), "!")
	if len(ptrs) < 2 {
		return ""
	}
	return ptrs[1]
}

// Returns the value for the Properties key
func (d DIDAccumulator) Value() string {
	return d.Identifier().GetKey(d.Key())
}

// Identifier returns the identifier of the resource
func (d DIDAccumulator) Identifier() DIDIdentifier {
	ptrs := strings.Split(string(d), "!")
	if len(ptrs) < 2 {
		return ""
	}
	return DIDIdentifier(ptrs[0])
}

// Add adds a new element to the accumulator
func (zk DIDAccumulator) Add(id string, key DIDSecretKey) error {
	accKey, err := key.AccumulatorKey()
	if err != nil {
		return err
	}
	e, err := zk.acc()
	if err != nil {
		return err
	}
	res, err := e.Add(accKey, stringToZkElement(id))
	if err != nil {
		return err
	}
	bz, err := res.MarshalBinary()
	if err != nil {
		return err
	}
	newVal := crypto.Base64Encode(bz)
	zk.Identifier().SetKey(zk.Key(), newVal)
	return nil
}

// Remove removes an element from the accumulator
func (zk DIDAccumulator) Remove(id string, key DIDSecretKey) error {
	accKey, err := key.AccumulatorKey()
	if err != nil {
		return err
	}
	e, err := zk.acc()
	if err != nil {
		return err
	}
	res, err := e.Remove(accKey, stringToZkElement(id))
	if err != nil {
		return err
	}
	bz, err := res.MarshalBinary()
	if err != nil {
		return err
	}
	newVal := crypto.Base64Encode(bz)
	zk.Identifier().SetKey(zk.Key(), newVal)
	return nil
}

// Contains checks if the accumulator contains an element
func (zk DIDAccumulator) Validate(id string, key DIDSecretKey) (bool, error) {
	accKey, err := key.AccumulatorKey()
	if err != nil {
		return false, err
	}
	e, err := zk.acc()
	if err != nil {
		return false, err
	}
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	pub, err := accKey.GetPublicKey(curve)
	if err != nil {
		return false, err
	}
	wit, err := new(accumulator.MembershipWitness).New(stringToZkElement(id), e, accKey)
	if err != nil {
		return false, err
	}
	err = wit.Verify(pub, e)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// String returns the string representation of the property
func (zk DIDAccumulator) String() string {
	return string(zk)
}

// acc returns the accumulator for the ZKSet
func (s DIDAccumulator) acc() (*accumulator.Accumulator, error) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	if s.Value() == "" {
		e, err := new(accumulator.Accumulator).New(curve)
		if err != nil {
			return nil, err
		}
		return e, nil
	}
	accBz, err := crypto.Base64Decode(s.Value())
	if err != nil {
		return nil, err
	}
	e, err := new(accumulator.Accumulator).New(curve)
	if err != nil {
		return nil, err
	}
	err = e.UnmarshalBinary(accBz)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func stringToZkElement(str string) accumulator.Element {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	return curve.Scalar.Hash([]byte(str))
}
