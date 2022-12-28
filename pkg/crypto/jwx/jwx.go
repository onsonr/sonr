package jwx

import (
	"encoding/json"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwe"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jws"
)

// JWX is a type that can marshal and unmarshal itself to and from JSON, create a JWK for encryption,
// create a JWK for signing, create a JWS, verify a JWS, encrypt a JWE, decrypt a JWE, and sign a JWS.
// @property CreateEncJWK - Creates a JWK for encryption
// @property CreateSignJWK - Creates a JWK for signing.
// @property CreateJWS - Creates a JWS (JSON Web Signature)
// @property VerifyJWS - Verify the JWS signature and return the payload.
// @property EncryptJWE - Encrypts the payload using the JWK.
// @property DecryptJWE - Decrypts a JWE payload.
// @property Sign - This is the method that will be used to sign the JWT.
type JWX interface {
	json.Marshaler
	json.Unmarshaler

	CreateEncJWK() (*jwk.Key, error)
	CreateSignJWK() (*jwk.Key, error)
	CreateJWS() ([]byte, error)
	VerifyJWS(payload []byte, opts ...jws.VerifyOption) ([]byte, error)
	EncryptJWE(payload []byte, options ...jwe.EncryptOption) ([]byte, error)
	DecryptJWE(payload []byte, opts ...jwe.DecryptOption) ([]byte, error)
	Sign([]byte, ...jws.SignOption) ([]byte, error)
}

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
