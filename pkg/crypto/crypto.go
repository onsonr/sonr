package crypto

import (
	"errors"
	"fmt"

	"strings"

	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	mb "github.com/multiformats/go-multibase"
	"github.com/multiformats/go-varint"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"

	types "github.com/sonrhq/core/types/crypto"
)

// SNRPubKey is a type alias for common.SNRPubKey in pkg/common.
type SNRPubKey = types.SNRPubKey

// CoinType is a type alias for types.CoinType in pkg/crypto/internal/types.
type CoinType = types.CoinType

// KeyType is a type alias for types.KeyType in pkg/crypto/internal/types.
type KeyType = types.KeyType

// PubKey is a type alias for types.PubKey in pkg/crypto/internal/types.
type PubKey = types.PubKey

// WebauthnCredential is a type alias for types.WebauthnCredential in pkg/crypto/internal/types.
type WebauthnCredential = types.WebauthnCredential

// WebauthnAuthenticator is a type alias for types.WebauthnAuthenticator in pkg/crypto/internal/types.
type WebauthnAuthenticator = types.WebauthnAuthenticator

// BTCCoinType is the CoinType for Bitcoin.
const BTCCoinType = types.CoinType_CoinType_BITCOIN

// ETHCoinType is the CoinType for Ethereum.
const ETHCoinType = types.CoinType_CoinType_ETHEREUM

// LTCCoinType is the CoinType for Litecoin.
const LTCCoinType = types.CoinType_CoinType_LITECOIN

// DOGECoinType is the CoinType for Dogecoin.
const DOGECoinType = types.CoinType_CoinType_DOGE

// SONRCoinType is the CoinType for Sonr.
const SONRCoinType = types.CoinType_CoinType_SONR

// COSMOSCoinType is the CoinType for Cosmos.
const COSMOSCoinType = types.CoinType_CoinType_COSMOS

// FILECOINCoinType is the CoinType for Filecoin.
const FILCoinType = types.CoinType_CoinType_FILECOIN

// HNSCoinType is the CoinType for Handshake.
const HNSCoinType = types.CoinType_CoinType_HNS

// TestCoinType is the CoinType for Testnet.
const TestCoinType = types.CoinType_CoinType_TESTNET

// SOLCoinType is the CoinType for Solana.
const SOLCoinType = types.CoinType_CoinType_SOLANA

// XRPCoinType is the CoinType for XRP.
const XRPCoinType = types.CoinType_CoinType_XRP

// AllCoinTypes is a slice of all CoinTypes.
var AllCoinTypes = types.AllCoinTypes

// CoinTypeFromAddrPrefix returns the CoinType from the public key address prefix (btc, eth).
func CoinTypeFromAddrPrefix(str string) CoinType {
	return types.CoinTypeFromAddrPrefix(str)
}

// CoinTypeFromBipPath returns the CoinType from the BIP Path (0, 60).
func CoinTypeFromBipPath(i int32) CoinType {
	return types.CoinTypeFromBipPath(i)
}

// CoinTypeFromDidMethod returns the CoinType from the DID Method (btc, eth).
func CoinTypeFromDidMethod(str string) CoinType {
	return types.CoinTypeFromDidMethod(str)
}

// CoinTypeFromName returns the CoinType from the Blockchain name (Bitcoin, Ethereum).
func CoinTypeFromName(str string) CoinType {
	return types.CoinTypeFromName(str)
}

// CoinTypeFromTicker returns the CoinType from the tokens Ticker (BTC, ETH).
func CoinTypeFromTicker(str string) CoinType {
	return types.CoinTypeFromTicker(str)
}

// Secp256k1KeyType is the key type for secp256k1.
const Secp256k1KeyType = types.KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019

// Ed25519KeyType is the key type for ed25519.
const Ed25519KeyType = types.KeyType_KeyType_ED25519_VERIFICATION_KEY_2018

// RSAKeyType is the key type for RSA.
const RSAKeyType = types.KeyType_KeyType_RSA_VERIFICATION_KEY_2018

// WebAuthnKeyType is the key type for WebAuthn.
const WebAuthnKeyType = types.KeyType_KeyType_WEB_AUTHN_AUTHENTICATION_2018

// NewPubKeyFromCmpConfig takes a `cmp.Config` and returns a `PubKey`
func NewPubKeyFromCmpConfig(config *cmp.Config) (*PubKey, error) {
	skPP, ok := config.PublicPoint().(*curve.Secp256k1Point)
	if !ok {
		return nil, errors.New("invalid public point")
	}
	bz, err := skPP.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return types.NewPubKey(bz, Secp256k1KeyType), nil
}


// NewSecp256k1PubKey takes a byte array of raw public key bytes and returns a PubKey.
func NewSecp256k1PubKey(bz []byte) *PubKey {
	return types.NewPubKey(bz, Secp256k1KeyType)
}

// NewEd25519PubKey takes a byte array of raw public key bytes and returns a PubKey.
func NewEd25519PubKey(bz []byte) *PubKey {
	return types.NewPubKey(bz, Ed25519KeyType)
}

// NewRSAPubKey takes a byte array of raw public key bytes and returns a PubKey.
func NewRSAPubKey(bz []byte) *PubKey {
	return types.NewPubKey(bz, RSAKeyType)
}

// NewWebAuthnPubKey takes a byte array of raw public key bytes and returns a PubKey.
func NewWebAuthnPubKey(bz []byte) *PubKey {
	return types.NewPubKey(bz, WebAuthnKeyType)
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
	kt, err := types.KeyTypeFromMulticodec(code)
	if err != nil {
		return nil, err
	}
	return types.NewPubKey(data[n:], kt), nil
}

// PubKeyFromBytes takes a byte array and returns a PubKey
func PubKeyFromBytes(bz []byte) (*PubKey, error) {
	code, n, err := varint.FromUvarint(bz)
	if err != nil {
		return nil, err
	}
	kt, err := types.KeyTypeFromMulticodec(code)
	if err != nil {
		return nil, err
	}
	return types.NewPubKey(bz[n:], kt), nil
}

// PubKeyFromCommon takes a common.SNRPubKey and returns a PubKey
func PubKeyFromCommon(pk SNRPubKey) (*PubKey, error) {
	t, err := types.KeyTypeFromPrettyString(pk.Type())
	if err != nil {
		return nil, fmt.Errorf("error retreiving key type from PubKey interface: %w", err)
	}
	return types.NewPubKey(pk.Raw(), t), nil
}

// PubKeyFromWebAuthn takes a webauthncose.Key and returns a PubKey
func PubKeyFromWebAuthn(cred *types.WebauthnCredential) (*PubKey, error) {
	if cred == nil {
		return nil, errors.New("credential is nil")
	}
	pub, err := webauthncose.ParsePublicKey(cred.PublicKey)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case webauthncose.EC2PublicKeyData:
		return NewSecp256k1PubKey(pub.XCoord), nil
	case webauthncose.OKPPublicKeyData:
		return NewEd25519PubKey(pub.XCoord), nil
	default:
		return nil, fmt.Errorf("unsupported public key type: %T", pub)
	}
}
