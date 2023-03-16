package crypto

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"hash"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	mb "github.com/multiformats/go-multibase"
	"github.com/multiformats/go-varint"
	common "github.com/sonrhq/core/types/common"
	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	tmcrypto "github.com/tendermint/tendermint/crypto"
	"golang.org/x/crypto/ripemd160"
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
	// Get base64 encoding of the key
	b64 := pk.Base64()
	// Get the sha256 hash of the key
	hasher := sha256.New()
	hasher.Write([]byte(b64))
	// Get the ripemd160 hash of the sha256 hash
	hasherRIPEMD160 := ripemd160.New()
	hasherRIPEMD160.Write(hasher.Sum(nil))
	// Return the ripemd160 hash
	return tmcrypto.Address(hasherRIPEMD160.Sum(nil))
}

// AddrString returns the address of the key.
func (pk *PubKey) AddrString(ct CoinType) (string, error) {
	if ct.IsEthereum() {
		return pk.Keccak256(), nil
	}
	return pk.Bech32(ct.AddrPrefix())
}

// Base64 returns the base64 encoding of the key.
func (pk *PubKey) Base64() string {
	return base64.RawStdEncoding.EncodeToString(pk.Bytes())
}

// Bech32 returns the bech32 encoding of the key. This is used for the Cosmos address.
func (pk *PubKey) Bech32(pfix string) (string, error) {
	return bech32.ConvertAndEncode(pfix, pk.Bytes())
}

// Keccak256 returns the keccak256 hash of the key. This is used for the Ethereum address.
func (pk *PubKey) Keccak256() string {
	hash := ethcrypto.Keccak256(pk.Bytes()[1:])
	addressBytes := hash[len(hash)-20:]
	// Convert the address bytes to a hexadecimal string
	address := hex.EncodeToString(addressBytes)
	return "0x" + address
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
