package jwx

import (
	"github.com/lestrrat-go/jwx/jwk"
)

type JWKSet struct {
	Key       *jwk.Key
	Signer    signer
	Signature map[[]byte][]byte
}

type KeyType = string

var (
	Type_ENC KeyType = "enc"
	Type_SIG KeyType = "sig"
)

func NewKeySigSet(kt KeyType, key interface{}) (*JWKSet, error) {
	set := &JWKSet{}
	switch kt {
	case Type_ENC:
		key, err := CreateJWKForEnc(key)
		if err != nil {
			return nil, err
		}
		set.Key = &key
		set.Signer = nil
	case Type_SIG:
		key, err := CreateJWKForSig(key)
		if err != nil {
			return nil, err
		}
		set.Key = &key
		set.Signer = CreateSigner()
	}

	return set, nil
}
