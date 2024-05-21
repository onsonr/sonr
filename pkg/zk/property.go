package zk

import (
	"errors"
	"fmt"

	"github.com/di-dao/core/crypto/accumulator"
	"github.com/di-dao/core/crypto/core/curves"
	"github.com/di-dao/core/x/did/types"
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

// deriveSecretKey derives the secret key from the keyshares
func DeriveSecretKey(propertyKey string, pubKey *types.PublicKey) (*SecretKey, error) {
	// Concatenate the controller's public key and the property key
	input := append(pubKey.Bytes(), []byte(propertyKey)...)
	hash := blake3Hash(input)

	// Use the hash as the seed for the secret key
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	key, err := new(accumulator.SecretKey).New(curve, hash[:])
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to create secret key"))
	}
	return &SecretKey{SecretKey: key}, nil
}

// Blake3Hash returns the blake3 hash of the input bytes
func blake3Hash(bz []byte) []byte {
	bz32 := blake3.Sum256(bz)
	return bz32[:]
}
