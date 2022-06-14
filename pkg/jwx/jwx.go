package jwx

import (
	"encoding/json"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jws"
)

type JWX interface {
	json.Marshaler
	json.Unmarshaler

	CreateEncJWK() (*jwk.Key, error)
	CreateSignJWK() (*jwk.Key, error)
	CreateJWS() ([]byte, error)
	VerifyJWS()  // TODO: Determine method signature
	EncryptJWE() // TODO: Determine method signature
	DecryptJWE() // TODO: Determine method signature
	Sign([]byte, ...jws.SignOption) ([]byte, error)
}

type KeyType = string

var (
	Type_ENC KeyType = "enc"
	Type_SIG KeyType = "sig"
)

type jwxImpl struct {
	key     interface{}
	keyAlg  jwa.KeyEncryptionAlgorithm
	contAlg jwa.ContentEncryptionAlgorithm
	sigAlg  jwa.SignatureAlgorithm
}

func New(key interface{}) *jwxImpl {
	return &jwxImpl{
		key:     key,
		keyAlg:  jwa.ECDH_ES_A256KW,
		contAlg: jwa.A128CBC_HS256,
	}
}
