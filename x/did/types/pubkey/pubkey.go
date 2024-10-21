package pubkey

import (
	"strings"

	didv1 "github.com/onsonr/sonr/api/did/v1"
)

type PubKeyI interface {
	GetRole() string
	GetKeyType() string
	GetRawKey() *didv1.RawKey
	GetJwk() *didv1.JSONWebKey
}

// PubKey defines a generic pubkey.
type PublicKey interface {
	VerifySignature(msg, sig []byte) bool
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
