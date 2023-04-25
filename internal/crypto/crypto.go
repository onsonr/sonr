package crypto

import (
	"encoding/base64"
	"errors"
	"fmt"

	"strings"

	mb "github.com/multiformats/go-multibase"
	"github.com/multiformats/go-varint"
	"github.com/shengdoushi/base58"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"

	types "github.com/sonrhq/core/types/crypto"
)

// CoinType is a type alias for types.CoinType in pkg/crypto/internal/types.
type CoinType = types.CoinType

// KeyType is a type alias for types.KeyType in pkg/crypto/internal/types.
type KeyType = types.KeyType

// PubKey is a type alias for types.PubKey in pkg/crypto/internal/types.
type PubKey = types.PubKey

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

func DeriveBIP44(config *cmp.Config, coinType types.CoinType, account, change, addressIndex uint32) (*cmp.Config, error) {
	purpose := uint32(44)

	// m / purpose'
	configPurpose, err := config.DeriveBIP32(purpose | 0x80000000)
	if err != nil {
		return nil, err
	}

	// m / purpose' / coin_type'
	configCoinType, err := configPurpose.DeriveBIP32(uint32(coinType.BipPath()) | 0x80000000)
	if err != nil {
		return nil, err
	}

	// m / purpose' / coin_type' / account'
	configAccount, err := configCoinType.DeriveBIP32(account | 0x80000000)
	if err != nil {
		return nil, err
	}

	// m / purpose' / coin_type' / account' / change
	configChange, err := configAccount.DeriveBIP32(change)
	if err != nil {
		return nil, err
	}

	// m / purpose' / coin_type' / account' / change / address_index
	configAddress, err := configChange.DeriveBIP32(addressIndex)
	if err != nil {
		return nil, err
	}

	return configAddress, nil
}

const hardenedOffset uint32 = 0x80000000

func formatBip44Path(coinType types.CoinType, idx int) []uint32 {
	purpose := uint32(44)
	coinTypeNum := uint32(coinType.BipPath())
	account := uint32(0)
	change := uint32(0)
	addressIndex := uint32(idx)
	return []uint32{
		purpose + hardenedOffset,
		coinTypeNum + hardenedOffset,
		account + hardenedOffset,
		change,
		addressIndex,
	}
}
