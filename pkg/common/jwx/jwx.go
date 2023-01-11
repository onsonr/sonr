package jwx

import (
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

type KeyType = string

var (
	Type_ENC KeyType = "enc"
	Type_SIG KeyType = "sig"
)

type jwxImpl struct {
	jwk     jwk.Key
	key     interface{}
	keyAlg  jwa.KeyEncryptionAlgorithm
	contAlg jwa.ContentEncryptionAlgorithm
	sigAlg  jwa.SignatureAlgorithm
}

// It creates a new JWE object, using the given key, and sets the default key and content encryption
// algorithms to ECDH-ES and AES-CBC-HMAC-SHA256, respectively
func New(key interface{}) *jwxImpl {
	return &jwxImpl{
		jwk:     nil,
		key:     key,
		keyAlg:  jwa.ECDH_ES_A256KW,
		contAlg: jwa.A128CBC_HS256,
	}
}
