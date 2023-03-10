package types

import (
	"bytes"
	"errors"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	mb "github.com/multiformats/go-multibase"
	"github.com/multiformats/go-varint"
	common "github.com/sonrhq/core/pkg/common"
	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	tmcrypto "github.com/tendermint/tendermint/crypto"
)

type (
	Address   = tmcrypto.Address
	SNRPubKey = common.SNRPubKey
)

//
// Constructors
//

// NewPubKey takes a byte array and returns a PubKey
func NewPubKey(bz []byte, kt KeyType) *PubKey {
	pk := &PubKey{}
	pk.Key = bz
	pk.KeyType = kt.PrettyString()
	return pk
}

//
// CryptoTypes Implementation of PubKey interface
//

// Creating a new method called Address() that returns an Address type.
func (pk *PubKey) Address() Address {
	return tmcrypto.AddressHash(pk.Bytes())
}

// Bech32 returns the bech32 encoding of the key.
func (pk *PubKey) Bech32(pfix string) (string, error) {
	return bech32.ConvertAndEncode(pfix, pk.Address().Bytes())
}

// DID returns the DID of the verification method
func (pk *PubKey) DID(opts ...common.DIDOption) string {
	c := common.DefaultDidUriConfig()
	opts = append(opts, common.WithIdentifier(pk.Multibase()))
	return c.Apply(opts...)
}

// Multibase returns the Base58 encoding the key.
func (pk *PubKey) Multibase() string {
	b58BKeyStr, err := mb.Encode(mb.Base58BTC, pk.Bytes())
	if err != nil {
		return ""
	}
	return b58BKeyStr
}

// Returning the key in bytes.
func (pk *PubKey) Bytes() []byte {
	raw := pk.Key
	// Get the multicodec code for the key type
	kt, err := KeyTypeFromPrettyString(pk.KeyType)
	if err != nil {
		return nil
	}
	t := kt.MulticodecType()
	size := varint.UvarintSize(t)
	data := make([]byte, size+len(raw))
	n := varint.PutUvarint(data, t)
	copy(data[n:], raw)
	return pk.Key
}

// Comparing the two keys.
func (pk *PubKey) Equals(other cryptotypes.PubKey) bool {
	if other == nil {
		return false
	}
	return bytes.Equal(pk.Bytes(), other.Bytes())
}

// Raw returns the raw key without the type.
func (pk *PubKey) Raw() []byte {
	return pk.Key
}

// // Returning the type of the key.
func (pk *PubKey) Type() string {
	return pk.KeyType
}

// Verifying the signature of the message.
func (pk *PubKey) VerifySignature(msg []byte, sig []byte) bool {
	if pk.KeyType == KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019.FormatString() {
		pp := &curve.Secp256k1Point{}
		if err := pp.UnmarshalBinary(pk.Key); err != nil {
			return false
		}
		signature, err := deserializeSignature(sig)
		if err != nil {
			return false
		}
		return signature.Verify(pp, msg)
	}
	if pk.KeyType == KeyType_KeyType_WEB_AUTHN_AUTHENTICATION_2018.FormatString() {
		keyFace, err := webauthncose.ParsePublicKey(pk.Key)
		if err != nil {
			return false
		}
		switch keyFace.(type) {
		case webauthncose.OKPPublicKeyData:
			key := keyFace.(webauthncose.OKPPublicKeyData)
			ok, err := key.Verify(msg, sig)
			if err != nil {
				return false
			}
			return ok
		case webauthncose.EC2PublicKeyData:
			key := keyFace.(webauthncose.EC2PublicKeyData)
			ok, err := key.Verify(msg, sig)
			if err != nil {
				return false
			}
			return ok
		case webauthncose.RSAPublicKeyData:
			key := keyFace.(webauthncose.RSAPublicKeyData)
			ok, err := key.Verify(msg, sig)
			if err != nil {
				return false
			}
			return ok
		default:
			return false
		}
	}
	return false
}

// SerializeSignature marshals an ECDSA signature to DER format for use with the CMP protocol
func serializeSignature(sig *ecdsa.Signature) ([]byte, error) {
	rBytes, err := sig.R.MarshalBinary()
	if err != nil {
		return nil, err
	}
	sBytes, err := sig.S.MarshalBinary()
	if err != nil {
		return nil, err
	}

	sigBytes := make([]byte, 65)
	// 0 pad the byte arrays from the left if they aren't big enough.
	copy(sigBytes[33-len(rBytes):33], rBytes)
	copy(sigBytes[65-len(sBytes):65], sBytes)
	return sigBytes, nil
}

// - The R and S values must be in the valid range for secp256k1 scalars:
//   - Negative values are rejected
//   - Zero is rejected
//   - Values greater than or equal to the secp256k1 group order are rejected
func deserializeSignature(sigStr []byte) (*ecdsa.Signature, error) {
	rBytes := sigStr[:33]
	sBytes := sigStr[33:65]

	sig := ecdsa.EmptySignature(curve.Secp256k1{})
	if err := sig.R.UnmarshalBinary(rBytes); err != nil {
		return nil, errors.New("malformed signature: R is not in the range [1, N-1]")
	}

	// S must be in the range [1, N-1].  Notice the check for the maximum number
	// of bytes is required because SetByteSlice truncates as noted in its
	// comment so it could otherwise fail to detect the overflow.
	if err := sig.S.UnmarshalBinary(sBytes); err != nil {
		return nil, errors.New("malformed signature: S is not in the range [1, N-1]")
	}

	// Create and return the signature.
	return &sig, nil
}