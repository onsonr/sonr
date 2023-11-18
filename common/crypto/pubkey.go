package crypto

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"

	tmcrypto "github.com/cometbft/cometbft/libs/bytes"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	secp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	mb "github.com/multiformats/go-multibase"
	"github.com/multiformats/go-varint"
	"lukechampine.com/blake3"
)

type (
	// Address is a type alias for tmcrypto.Address in libs/bytes.
	Address = tmcrypto.HexBytes
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

// NewSecp256k1PubKey takes a secp256k1.PubKey and returns a PubKey
func NewSecp256k1PubKey(pk *secp256k1.PubKey) *PubKey {
	return NewPubKey(pk.Bytes(), KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019)
}

//
// CryptoTypes Implementation of PubKey interface
//

// Address Creating a new method called Address() that returns an Address type.
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


// Multibase returns the Base58 encoding the key.
func (pk *PubKey) Multibase() string {
	b58BKeyStr, err := mb.Encode(mb.Base58BTC, pk.Bytes())
	if err != nil {
		return ""
	}
	return b58BKeyStr
}

// Bytes Returning the key in bytes.
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

// Equals Comparing the two keys.
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

// Secp256k1 Returning the secp256k1 public key.
func (pk *PubKey) Secp256k1() (*secp256k1.PubKey, error) {
	if len(pk.Bytes()) != 33 {
		return nil, errors.New("invalid public key length")
	}

	pubKey := &secp256k1.PubKey{Key: pk.Bytes()}
	return pubKey, nil
}

// Secp256k1AnyProto returns the pubkey for cosmos transactions
func (pk *PubKey) Secp256k1AnyProto() (*codectypes.Any, error) {
	scpk, err := pk.Secp256k1()
	if err != nil {
		return nil, err
	}
	return codectypes.NewAnyWithValue(scpk)
}

// Type Returning the type of the key.
func (pk *PubKey) Type() string {
	return pk.KeyType
}
