package keys

import (
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"strings"

	"github.com/libp2p/go-libp2p/core/crypto"
	mb "github.com/multiformats/go-multibase"
	varint "github.com/multiformats/go-varint"
)

const (
	// KeyPrefix indicates a decentralized identifier that uses the key method
	KeyPrefix = "did:key"
	// MulticodecKindRSAPubKey rsa-x509-pub https://github.com/multiformats/multicodec/pull/226
	MulticodecKindRSAPubKey = 0x1205
	// MulticodecKindEd25519PubKey ed25519-pub
	MulticodecKindEd25519PubKey = 0xed
	// MulticodecKindSecp256k1PubKey secp256k1-pub
	MulticodecKindSecp256k1PubKey = 0x1206
)

// DID is a DID:key identifier
type DID struct {
	crypto.PubKey
}

// NewDID constructs an Identifier from a public key
func NewDID(pub crypto.PubKey) (DID, error) {
	switch pub.Type() {
	case crypto.Ed25519, crypto.RSA, crypto.Secp256k1:
		return DID{PubKey: pub}, nil
	default:
		return DID{}, fmt.Errorf("unsupported key type: %s", pub.Type())
	}
}

// MulticodecType indicates the type for this multicodec
func (id DID) MulticodecType() uint64 {
	switch id.Type() {
	case crypto.RSA:
		return MulticodecKindRSAPubKey
	case crypto.Ed25519:
		return MulticodecKindEd25519PubKey
	case crypto.Secp256k1:
		return MulticodecKindSecp256k1PubKey
	default:
		panic("unexpected crypto type")
	}
}

// String returns this did:key formatted as a string
func (id DID) String() string {
	raw, err := id.Raw()
	if err != nil {
		return ""
	}

	t := id.MulticodecType()
	size := varint.UvarintSize(t)
	data := make([]byte, size+len(raw))
	n := varint.PutUvarint(data, t)
	copy(data[n:], raw)

	b58BKeyStr, err := mb.Encode(mb.Base58BTC, data)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s:%s", KeyPrefix, b58BKeyStr)
}

// VerifyKey returns the backing implementation for a public key, one of:
// *rsa.PublicKey, ed25519.PublicKey
func (id DID) VerifyKey() (interface{}, error) {
	rawPubBytes, err := id.PubKey.Raw()
	if err != nil {
		return nil, err
	}
	switch id.PubKey.Type() {
	case crypto.RSA:
		verifyKeyiface, err := x509.ParsePKIXPublicKey(rawPubBytes)
		if err != nil {
			return nil, err
		}
		verifyKey, ok := verifyKeyiface.(*rsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("public key is not an RSA key. got type: %T", verifyKeyiface)
		}
		return verifyKey, nil
	case crypto.Ed25519:
		return ed25519.PublicKey(rawPubBytes), nil
	case crypto.Secp256k1:
		// Handle both compressed and uncompressed Secp256k1 public keys
		if len(rawPubBytes) == 65 || len(rawPubBytes) == 33 {
			return rawPubBytes, nil
		}
		return nil, fmt.Errorf("invalid Secp256k1 public key length: %d", len(rawPubBytes))
	default:
		return nil, fmt.Errorf("unrecognized Public Key type: %s", id.Type())
	}
}

// Parse turns a string into a key method ID
func Parse(keystr string) (DID, error) {
	var id DID
	if !strings.HasPrefix(keystr, KeyPrefix) {
		return id, fmt.Errorf("decentralized identifier is not a 'key' type")
	}

	keystr = strings.TrimPrefix(keystr, KeyPrefix+":")

	enc, data, err := mb.Decode(keystr)
	if err != nil {
		return id, fmt.Errorf("decoding multibase: %w", err)
	}

	if enc != mb.Base58BTC {
		return id, fmt.Errorf("unexpected multibase encoding: %s", mb.EncodingToStr[enc])
	}

	keyType, n, err := varint.FromUvarint(data)
	if err != nil {
		return id, err
	}

	switch keyType {
	case MulticodecKindRSAPubKey:
		pub, err := crypto.UnmarshalRsaPublicKey(data[n:])
		if err != nil {
			return id, err
		}
		return DID{pub}, nil
	case MulticodecKindEd25519PubKey:
		pub, err := crypto.UnmarshalEd25519PublicKey(data[n:])
		if err != nil {
			return id, err
		}
		return DID{pub}, nil
	case MulticodecKindSecp256k1PubKey:
		// Handle both compressed and uncompressed formats
		keyData := data[n:]
		if len(keyData) != 33 && len(keyData) != 65 {
			return id, fmt.Errorf("invalid Secp256k1 public key length: %d", len(keyData))
		}
		pub, err := crypto.UnmarshalSecp256k1PublicKey(keyData)
		if err != nil {
			return id, fmt.Errorf("failed to unmarshal Secp256k1 key: %w", err)
		}
		return DID{pub}, nil
	}

	return id, fmt.Errorf("unrecognized key type multicodec prefix: %x", data[0])
}
