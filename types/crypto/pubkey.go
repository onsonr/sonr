package crypto

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"hash"

	"github.com/btcsuite/btcd/btcec"
	secp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	mb "github.com/multiformats/go-multibase"
	"github.com/multiformats/go-varint"
	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	tmcrypto "github.com/tendermint/tendermint/crypto"
	"lukechampine.com/blake3"
)

type (
	Address   = tmcrypto.Address
)


// `SNRPubKey` is a `PubKey` that has a `DID` and a `Multibase`
// @property {string} DID - The DID of the SNR
// @property {string} Multibase - The multibase encoding of the DID.
type SNRPubKey interface {
	cryptotypes.PubKey

	Bech32(pfix string) (string, error)
	Multibase() string
	Raw() []byte
}


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
	sckp, err := pk.Secp256k1()
	if err != nil {
		return nil
	}
	return sckp.Address()
}

// Base64 returns the base64 encoding of the key.
func (pk *PubKey) Base64() string {
	return base64.RawStdEncoding.EncodeToString(pk.Bytes())
}

// Blake3 returns the blake3 hash of the key.
func (pk *PubKey) Blake3() string {
	hasher := blake3.New(32, nil)
	hasher.Write(pk.Bytes())
	return hex.EncodeToString(hasher.Sum(nil))
}

// Btcec returns the btcec public key.
func (pk *PubKey) Btcec() (*btcec.PublicKey, error) {
	pubKey, err := btcec.ParsePubKey(pk.Bytes(), btcec.S256())
	if err != nil {
		return nil, errors.New("failed to parse public key")
	}
	return pubKey, nil
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

// Returning the secp256k1 public key.
func (pk *PubKey) Secp256k1() (*secp256k1.PubKey, error) {
	if len(pk.Bytes()) != 33 {
		return nil, errors.New("invalid public key length")
	}

	pubKey := &secp256k1.PubKey{Key: pk.Bytes()}
	return pubKey, nil
}

// Returning the type of the key.
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
		return verifyWebAuthnSignature(pk.Key, msg, sig)
	}
	return false
}

// Calculate the hash of hasher over buf.
func calcHash(buf []byte, hasher hash.Hash) []byte {
	hasher.Write(buf)
	return hasher.Sum(nil)
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

func verifyWebAuthnSignature(msg []byte, sig []byte, key []byte) bool {
	keyFace, err := webauthncose.ParsePublicKey(key)
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
