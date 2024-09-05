package types

import (
	"encoding/hex"

	"github.com/mr-tron/base58/base58"
)

//
// # Genesis Structures
//

// Equal returns true if two asset infos are equal
func (a *AssetInfo) Equal(b *AssetInfo) bool {
	if a == nil && b == nil {
		return true
	}
	return false
}

// Equal returns true if two chain infos are equal
func (c *ChainInfo) Equal(b *ChainInfo) bool {
	if c == nil && b == nil {
		return true
	}
	return false
}

// Equal returns true if two OpenID config infos are equal
func (o *OpenIDConfig) Equal(b *OpenIDConfig) bool {
	if o == nil && b == nil {
		return true
	}
	return false
}

// Equal returns true if two key infos are equal
func (k *KeyInfo) Equal(b *KeyInfo) bool {
	if k == nil && b == nil {
		return true
	}
	return false
}

// Equal returns true if two validator infos are equal
func (v *ValidatorInfo) Equal(b *ValidatorInfo) bool {
	if v == nil && b == nil {
		return true
	}
	return false
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
