package types

import (
	"encoding/hex"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/mr-tron/base58/base58"
	"github.com/onsonr/crypto"
)

var WalletKeyInfo = &KeyInfo{
	Role:      KeyRole_KEY_ROLE_DELEGATION,
	Curve:     KeyCurve_KEY_CURVE_SECP256K1,
	Algorithm: KeyAlgorithm_KEY_ALGORITHM_ECDSA,
	Encoding:  KeyEncoding_KEY_ENCODING_HEX,
	Type:      KeyType_KEY_TYPE_BIP32,
}

var EthKeyInfo = &KeyInfo{
	Role:      KeyRole_KEY_ROLE_DELEGATION,
	Curve:     KeyCurve_KEY_CURVE_KECCAK256,
	Algorithm: KeyAlgorithm_KEY_ALGORITHM_ECDSA,
	Encoding:  KeyEncoding_KEY_ENCODING_HEX,
	Type:      KeyType_KEY_TYPE_BIP32,
}

var SonrKeyInfo = &KeyInfo{
	Role:      KeyRole_KEY_ROLE_INVOCATION,
	Curve:     KeyCurve_KEY_CURVE_P256,
	Algorithm: KeyAlgorithm_KEY_ALGORITHM_ECDSA,
	Encoding:  KeyEncoding_KEY_ENCODING_HEX,
	Type:      KeyType_KEY_TYPE_MPC,
}

var ChainCodeKeyInfos = map[ChainCode]*KeyInfo{
	ChainCodeBTC: WalletKeyInfo,
	ChainCodeETH: EthKeyInfo,
	ChainCodeSNR: SonrKeyInfo,
	ChainCodeIBC: WalletKeyInfo,
}

// NewEthPublicKey returns a new ethereum public key
func NewPublicKey(data []byte, keyInfo *KeyInfo) (*PubKey, error) {
	encKey, err := keyInfo.Encoding.EncodeRaw(data)
	if err != nil {
		return nil, err
	}

	return &PubKey{
		Raw:       encKey,
		Role:      keyInfo.Role,
		Encoding:  keyInfo.Encoding,
		Algorithm: keyInfo.Algorithm,
		Curve:     keyInfo.Curve,
		KeyType:   keyInfo.Type,
	}, nil
}

// Address returns the address of the public key
func (k *PubKey) Address() cryptotypes.Address {
	return nil
}

// Bytes returns the raw bytes of the public key
func (k *PubKey) Bytes() []byte {
	bz, _ := k.GetEncoding().DecodeRaw(k.GetRaw())
	return bz
}

// Clone returns a copy of the public key
func (k *PubKey) Clone() cryptotypes.PubKey {
	return &PubKey{
		Raw:       k.GetRaw(),
		Role:      k.GetRole(),
		Encoding:  k.GetEncoding(),
		Algorithm: k.GetAlgorithm(),
		Curve:     k.GetCurve(),
		KeyType:   k.GetKeyType(),
	}
}

// VerifySignature verifies a signature over the given message
func (k *PubKey) VerifySignature(msg []byte, sig []byte) bool {
	pk, err := crypto.ComputeEcdsaPublicKey(k.Bytes())
	sigMpc, err := crypto.DeserializeMPCSignature(sig)
	if err != nil {
		return false
	}
	return crypto.VerifyMPCSignature(sigMpc, msg, pk)
}

// Equals returns true if two public keys are equal
func (k *PubKey) Equals(k2 cryptotypes.PubKey) bool {
	if k == nil && k2 == nil {
		return true
	}
	return false
}

// Type returns the type of the public key
func (k *PubKey) Type() string {
	return k.KeyType.String()
}

// DecodePublicKey extracts the public key from the given data
func (k *KeyInfo) DecodePublicKey(data interface{}) ([]byte, error) {
	var bz []byte
	switch v := data.(type) {
	case string:
		bz = []byte(v)
	case []byte:
		bz = v
	default:
		return nil, ErrUnsupportedKeyEncoding
	}

	if k.Encoding == KeyEncoding_KEY_ENCODING_RAW {
		return bz, nil
	}
	if k.Encoding == KeyEncoding_KEY_ENCODING_HEX {
		return hex.DecodeString(string(bz))
	}
	if k.Encoding == KeyEncoding_KEY_ENCODING_MULTIBASE {
		return base58.Decode(string(bz))
	}
	return nil, ErrUnsupportedKeyEncoding
}

// EncodePublicKey encodes the public key according to the KeyInfo's encoding
func (k *KeyInfo) EncodePublicKey(data []byte) (string, error) {
	if k.Encoding == KeyEncoding_KEY_ENCODING_RAW {
		return string(data), nil
	}
	if k.Encoding == KeyEncoding_KEY_ENCODING_HEX {
		return hex.EncodeToString(data), nil
	}
	if k.Encoding == KeyEncoding_KEY_ENCODING_MULTIBASE {
		return base58.Encode(data), nil
	}
	return "", ErrUnsupportedKeyEncoding
}
