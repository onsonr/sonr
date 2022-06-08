package jwx

import "github.com/lestrrat-go/jwx/jwk"

type JWKSet struct {
	JWK    *jwk.Key
	signer signer
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
		set.JWK = &key
		set.signer = nil
	case Type_SIG:
		key, err := CreateJWKForSig(key)
		if err != nil {
			return nil, err
		}
		set.JWK = &key
		set.signer = CreateSigner()
	}

	return set, nil
}
