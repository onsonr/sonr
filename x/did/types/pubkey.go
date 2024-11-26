package types

import (
	"bytes"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/crypto/types"
	"google.golang.org/protobuf/proto"
)

type PubKeyI interface {
	GetRole() string
	GetKeyType() string
	// GetRawKey() *commonv1.RawKey
	// GetJwk() *commonv1.JSONWebKey
}

type PubKeyG[T any] interface {
	*T
	PublicKey
}

type pubKeyImpl struct {
	decode   func(b []byte) (PublicKey, error)
	validate func(key PublicKey) error
}

//	func WithSecp256K1PubKey() Option {
//		return WithPubKeyWithValidationFunc(func(pt *secp256k1.PubKey) error {
//			_, err := dcrd_secp256k1.ParsePubKey(pt.Key)
//			return err
//		})
//	}
//
//	func WithPubKey[T any, PT PubKeyG[T]]() Option {
//		return WithPubKeyWithValidationFunc[T, PT](func(_ PT) error {
//			return nil
//		})
//	}
//
//	func WithPubKeyWithValidationFunc[T any, PT PubKeyG[T]](validateFn func(PT) error) Option {
//		pkImpl := pubKeyImpl{
//			decode: func(b []byte) (PublicKey, error) {
//				key := PT(new(T))
//				err := gogoproto.Unmarshal(b, key)
//				if err != nil {
//					return nil, err
//				}
//				return key, nil
//			},
//			validate: func(k PublicKey) error {
//				concrete, ok := k.(PT)
//				if !ok {
//					return fmt.Errorf(
//						"invalid pubkey type passed for validation, wanted: %T, got: %T",
//						concrete,
//						k,
//					)
//				}
//				return validateFn(concrete)
//			},
//		}
//		return func(a *Account) {
//			a.supportedPubKeys[gogoproto.MessageName(PT(new(T)))] = pkImpl
//		}
//	}
func nameFromTypeURL(url string) string {
	name := url
	if i := strings.LastIndexByte(url, '/'); i >= 0 {
		name = name[i+len("/"):]
	}
	return name
}

// CustomPubKey represents a custom secp256k1 public key.
type CustomPubKey struct {
	proto.Message

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

// NewCustomPubKeyFromRawBytes creates a new CustomPubKey from raw bytes.
func NewCustomPubKeyFromRawBytes(key []byte) (*CustomPubKey, error) {
	// Validate the key length and format
	if len(key) != 33 {
		return nil, fmt.Errorf("invalid key length; expected 33 bytes, got %d", len(key))
	}
	if key[0] != 0x02 && key[0] != 0x03 {
		return nil, fmt.Errorf("invalid key format; expected 0x02 or 0x03 as the first byte, got 0x%02x", key[0])
	}

	return &CustomPubKey{Key: key}, nil
}

// Bytes returns the byte representation of the public key.
func (pk *CustomPubKey) Bytes() []byte {
	return pk.Key
}

// Equals checks if two public keys are equal.
func (pk *CustomPubKey) Equals(other sdk.PubKey) bool {
	return bytes.EqualFold(pk.Bytes(), other.Bytes())
}

// Type returns the type of the public key.
func (pk *CustomPubKey) Type() string {
	return "custom-secp256k1"
}

// Marshal implements the proto.Message interface.
func (pk *CustomPubKey) Marshal() ([]byte, error) {
	return proto.Marshal(pk)
}

// Unmarshal implements the proto.Message interface.
func (pk *CustomPubKey) Unmarshal(data []byte) error {
	return proto.Unmarshal(data, pk)
}

// Address returns the address derived from the public key.
func (pk *CustomPubKey) Address() []byte {
	// Implement address derivation logic here
	// For simplicity, this example uses a placeholder
	return []byte("derived-address")
}

// VerifySignature verifies a signature using the public key.
func (pk *CustomPubKey) VerifySignature(msg []byte, sig []byte) bool {
	// Implement signature verification logic here
	// For simplicity, this example uses a placeholder
	return true
}
