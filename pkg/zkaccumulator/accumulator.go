package zkaccumulator

import (
	"fmt"
	"strings"

	"github.com/sonr-io/kryptology/pkg/accumulator"
	"github.com/sonr-io/kryptology/pkg/core/curves"

	"github.com/sonrhq/sonr/internal/crypto"
	"github.com/sonrhq/sonr/pkg/didcommon"
)

// DIDAccumulator is a ZKSet accumulator for a DID
type DIDAccumulator string

// NewAccumulator creates a new accumulator for the DID
func NewAccumulator(id didcommon.Identifier, key string) DIDAccumulator {
	return DIDAccumulator(fmt.Sprintf("%s!%s", id.String(), key))
}

// Key returns the key of the resource
func (zk DIDAccumulator) Key() string {
	ptrs := strings.Split(string(zk), "!")
	if len(ptrs) < 2 {
		return ""
	}
	return ptrs[1]
}

// Value Returns the value for the Properties key
func (zk DIDAccumulator) Value() string {
	return zk.Identifier().GetKey(zk.Key())
}

// Identifier returns the identifier of the resource
func (zk DIDAccumulator) Identifier() didcommon.Identifier {
	ptrs := strings.Split(string(zk), "!")
	if len(ptrs) < 2 {
		return ""
	}
	return didcommon.Identifier(ptrs[0])
}

// Add adds a new element to the accumulator
func (zk DIDAccumulator) Add(id string, key didcommon.SecretKey) error {
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
func (zk DIDAccumulator) Remove(id string, key didcommon.SecretKey) error {
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

// Validate checks if the accumulator contains an element
func (zk DIDAccumulator) Validate(id string, key didcommon.SecretKey) (bool, error) {
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
func (zk DIDAccumulator) acc() (*accumulator.Accumulator, error) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	if zk.Value() == "" {
		e, err := new(accumulator.Accumulator).New(curve)
		if err != nil {
			return nil, err
		}
		return e, nil
	}
	accBz, err := crypto.Base64Decode(zk.Value())
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
