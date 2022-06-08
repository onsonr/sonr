package dag_jose

import (
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
)

type signer func(payload []byte, alg jwa.SignatureAlgorithm, key interface{}, options ...jws.SignOption) ([]byte, error)

func CreateSigner() signer {
	return func(payload []byte, alg jwa.SignatureAlgorithm, key interface{}, options ...jws.SignOption) ([]byte, error) {
		return jws.Sign(payload, alg, key)
	}
}
