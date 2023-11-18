package crypto

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	secp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	mb "github.com/multiformats/go-multibase"
	"github.com/multiformats/go-varint"
	"github.com/shengdoushi/base58"
)

// Secp256k1PubKey is a type alias for secp256k1.PubKey in pkg/crypto/keys/secp256k1.
type Secp256k1PubKey = secp256k1.PubKey

// BTCCoinType is the CoinType for Bitcoin.
const BTCCoinType = CoinType_CoinType_BITCOIN

// ETHCoinType is the CoinType for Ethereum.
const ETHCoinType = CoinType_CoinType_ETHEREUM

// LTCCoinType is the CoinType for Litecoin.
const LTCCoinType = CoinType_CoinType_LITECOIN

// DOGECoinType is the CoinType for Dogecoin.
const DOGECoinType = CoinType_CoinType_DOGE

// SONRCoinType is the CoinType for Sonr.
const SONRCoinType = CoinType_CoinType_SONR

// COSMOSCoinType is the CoinType for Cosmos.
const COSMOSCoinType = CoinType_CoinType_COSMOS

// FILCoinType is the CoinType for Filecoin.
const FILCoinType = CoinType_CoinType_FILECOIN

// HNSCoinType is the CoinType for Handshake.
const HNSCoinType = CoinType_CoinType_HNS

// TestCoinType is the CoinType for Testnet.
const TestCoinType = CoinType_CoinType_TESTNET

// SOLCoinType is the CoinType for Solana.
const SOLCoinType = CoinType_CoinType_SOLANA

// XRPCoinType is the CoinType for XRP.
const XRPCoinType = CoinType_CoinType_XRP

// Secp256k1KeyType is the key type for secp256k1.
const Secp256k1KeyType = KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019

// Ed25519KeyType is the key type for ed25519.
const Ed25519KeyType = KeyType_KeyType_ED25519_VERIFICATION_KEY_2018

// RSAKeyType is the key type for RSA.
const RSAKeyType = KeyType_KeyType_RSA_VERIFICATION_KEY_2018

// WebAuthnKeyType is the key type for WebAuthn.
const WebAuthnKeyType = KeyType_KeyType_WEB_AUTHN_AUTHENTICATION_2018

// Base58Encode takes a byte array and returns a base68 encoded string.
func Base58Encode(bz []byte) string {
	return base58.Encode(bz, base58.BitcoinAlphabet)
}

// Base58Decode takes a base68 encoded string and returns a byte array.
func Base58Decode(str string) ([]byte, error) {
	return base58.Decode(str, base58.BitcoinAlphabet)
}

// Base64Encode takes a byte array and returns a base64 encoded string.
func Base64Encode(bz []byte) string {
	return base64.RawURLEncoding.EncodeToString(bz)
}

// Base64Decode takes a base64 encoded string and returns a byte array.
func Base64Decode(str string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(str)
}

// HexEncode takes a byte array and returns a hex encoded string.
func HexEncode(bz []byte) string {
	return hex.EncodeToString(bz)
}

// HexDecode takes a hex encoded string and returns a byte array.
func HexDecode(str string) ([]byte, error) {
	return hex.DecodeString(str)
}

// NewEd25519PubKey takes a byte array of raw public key bytes and returns a PubKey.
func NewEd25519PubKey(bz []byte) *PubKey {
	return NewPubKey(bz, Ed25519KeyType)
}

// NewRSAPubKey takes a byte array of raw public key bytes and returns a PubKey.
func NewRSAPubKey(bz []byte) *PubKey {
	return NewPubKey(bz, RSAKeyType)
}

// NewWebAuthnPubKey takes a byte array of raw public key bytes and returns a PubKey.
func NewWebAuthnPubKey(bz []byte) *PubKey {
	return NewPubKey(bz, WebAuthnKeyType)
}

// PubKeyFromDID takes a string of a DID, decodes it from base58, unmarshals it into a PubKey, and returns the PubKey
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
