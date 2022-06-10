package jwx

import (
	"github.com/lestrrat-go/jwx/jwk"
)

type JWX interface {
	CreateJWK() (*jwk.Key, error)
	CreateJWS() ([]byte, error)
	VerifyJWS()
	EncryptJWE()
	DecryptJWE()
}

type jwxImpl struct {
	SigPair JWKSignaturePair
	keyAlg
}
