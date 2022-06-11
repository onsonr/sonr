package jwx

import (
	"github.com/lestrrat-go/jwx/v2/jwk"
)

type JWKSignaturePair struct {
	Key        *jwk.Key
	Signer     signer
	Signatures map[string][]byte // map of name to signature, name of signature must be provided
}

type KeyType = string

var (
	Type_ENC KeyType = "enc"
	Type_SIG KeyType = "sig"
)

func NewKeySigSet(kt KeyType, key interface{}) (*JWKSignaturePair, error) {
	set := &JWKSignaturePair{}
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
		set.Signer = createSigner()
	}

	return set, nil
}

func (jsp *JWKSignaturePair) AddSignature(name string, sig []byte) {
	jsp.Signatures[name] = sig
}

func (jsp *JWKSignaturePair) Marshal() {}

func (jsp *JWKSignaturePair) Unmarshal() {}
