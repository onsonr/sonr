package jwx

import (
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

type JWX interface {
	CreateEncJWK() (*jwk.Key, error)
	CreateSignJWK() (*jwk.Key, error)
	CreateJWS() ([]byte, error)
	VerifyJWS()
	EncryptJWE()
	DecryptJWE()
	Sign()
}

type jwxImpl struct {
	key     interface{}
	keyAlg  jwa.KeyEncryptionAlgorithm
	contAlg jwa.ContentEncryptionAlgorithm
	sigAlg  jwa.SignatureAlgorithm

	Sign signer
}

func New(key interface{}, keyAlg jwa.KeyEncryptionAlgorithm, contAlg jwa.ContentEncryptionAlgorithm) *jwxImpl {
	return &jwxImpl{
		key:     key,
		keyAlg:  keyAlg,
		contAlg: contAlg,

		Sign: createSigner(),
	}
}
