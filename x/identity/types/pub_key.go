package types

import (
	"bytes"
	"errors"
	fmt "fmt"
	"strings"

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
	Address = tmcrypto.Address
)

//
// Constructors
//

// NewPubKey takes a byte array and returns a PubKey
func NewPubKey(bz []byte, kt KeyType) *PubKey {
	pk := &PubKey{}
	pk.Key = bz
	pk.KeyType = kt
	return pk
}

// It takes a string of a DID, decodes it from base58, unmarshals it into a PubKey, and returns the PubKey
func PubKeyFromDID(did string) (*PubKey, error) {
	ptrs := strings.Split(did, ":")
	keystr := ptrs[len(ptrs)-1]

	enc, data, err := mb.Decode(keystr)
	if err != nil {
		return nil, fmt.Errorf("decoding multibase: %w", err)
	}

	if enc != mb.Base58BTC {
		return nil, fmt.Errorf("unexpected multibase encoding: %s", mb.EncodingToStr[enc])
	}

	code, n, err := varint.FromUvarint(data)
	if err != nil {
		return nil, err
	}
	kt, err := KeyTypeFromMulticodec(code)
	if err != nil {
		return nil, err
	}
	return NewPubKey(data[n:], kt), nil
}

// PubKeyFromCurvePoint takes a curve point and returns a PubKey
func PubKeyFromCurvePoint(p *curve.Secp256k1Point) (*PubKey, error) {
	bz, err := p.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return NewPubKey(bz, KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019), nil
}

// PubKeyFromBytes takes a byte array and returns a PubKey
func PubKeyFromBytes(bz []byte) (*PubKey, error) {
	code, n, err := varint.FromUvarint(bz)
	if err != nil {
		return nil, err
	}
	kt, err := KeyTypeFromMulticodec(code)
	if err != nil {
		return nil, err
	}
	return NewPubKey(bz[n:], kt), nil
}

// PubKeyFromCommon takes a common.SNRPubKey and returns a PubKey
func PubKeyFromCommon(pk common.SNRPubKey) (*PubKey, error) {
	t, err := KeyTypeFromPrettyString(pk.Type())
	if err != nil {
		return nil, fmt.Errorf("error retreiving key type from PubKey interface: %w", err)
	}
	return NewPubKey(pk.Raw(), t), nil
}

// PubKeyFromWebAuthn takes a webauthncose.Key and returns a PubKey
func PubKeyFromWebAuthn(cred *common.WebauthnCredential) (*PubKey, error) {
	if cred == nil {
		return nil, errors.New("credential is nil")
	}
	pub, err := webauthncose.ParsePublicKey(cred.PublicKey)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *webauthncose.EC2PublicKeyData:
		return NewPubKey(pub.XCoord, KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019), nil
	case *webauthncose.OKPPublicKeyData:
		return NewPubKey(pub.XCoord, KeyType_KeyType_ED25519_VERIFICATION_KEY_2018), nil
	default:
		return nil, fmt.Errorf("unsupported public key type: %T", pub)
	}
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
	return bech32.ConvertAndEncode(pfix, pk.Bytes())
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
	t := pk.KeyType.MulticodecType()
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
	return pk.KeyType.PrettyString()
}

// Verifying the signature of the message.
func (pk *PubKey) VerifySignature(msg []byte, sig []byte) bool {
	if pk.KeyType == KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019 {
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
	if pk.KeyType == KeyType_KeyType_WEB_AUTHN_AUTHENTICATION_2018 {
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

// VerificationMethod applies the given options and builds a verification method from this Key
func (pk *PubKey) VerificationMethod(opts ...VerificationMethodOption) (*VerificationMethod, error) {
	vm := &VerificationMethod{
		Id:                 pk.DID(),
		Type:               pk.KeyType,
		PublicKeyMultibase: pk.Multibase(),
	}
	for _, opt := range opts {
		if err := opt(vm); err != nil {
			return nil, err
		}
	}
	return vm, nil
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

// -- We represent those as raw public key bytes prefixed with public key
// -- multiformat code.
// | secp256k1  "0xe7"
// | Ed25519    "0xed"
// | P256       "0x1200"
// | P384       "0x1201"
// | P512       "0x1202"
// | RSA        "0x1205"
//
// MulticodecType returns the multicodec code for the key type
func (kt KeyType) MulticodecType() uint64 {
	switch kt {
	case KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019:
		return 0xe7
	case KeyType_KeyType_ED25519_VERIFICATION_KEY_2018:
		return 0xed
	case KeyType_KeyType_JSON_WEB_KEY_2020:
		return 0x1200
	case KeyType_KeyType_RSA_VERIFICATION_KEY_2018:
		return 0x1205
	default:
		return 0
	}
}

// PrettyString returns the string representation of the key type
func (kt KeyType) PrettyString() string {
	switch kt {
	case KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019:
		return "secp256k1"
	case KeyType_KeyType_ED25519_VERIFICATION_KEY_2018:
		return "Ed25519"
	case KeyType_KeyType_JSON_WEB_KEY_2020:
		return "JWK"
	case KeyType_KeyType_RSA_VERIFICATION_KEY_2018:
		return "RSA"
	default:
		return "unknown"
	}
}

// KeyTypeFromMulticodec returns the key type
func KeyTypeFromMulticodec(code uint64) (KeyType, error) {
	switch code {
	case 0xe7:
		return KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019, nil
	case 0xed:
		return KeyType_KeyType_ED25519_VERIFICATION_KEY_2018, nil
	case 0x1200:
		return KeyType_KeyType_JSON_WEB_KEY_2020, nil
	case 0x1205:
		return KeyType_KeyType_RSA_VERIFICATION_KEY_2018, nil
	default:
		return KeyType_KeyType_UNSPECIFIED, fmt.Errorf("unknown key type code: %d", code)
	}
}

// KeyTypeFromPrettyString returns the key type
func KeyTypeFromPrettyString(s string) (KeyType, error) {
	switch s {
	case "secp256k1":
		return KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019, nil
	case "Ed25519":
		return KeyType_KeyType_ED25519_VERIFICATION_KEY_2018, nil
	case "JWK":
		return KeyType_KeyType_JSON_WEB_KEY_2020, nil
	case "RSA":
		return KeyType_KeyType_RSA_VERIFICATION_KEY_2018, nil
	default:
		return KeyType_KeyType_UNSPECIFIED, fmt.Errorf("unknown key type: %s", s)
	}
}
