package jwx

import (
	"github.com/lestrrat-go/jwx/v2/jws"
)

func (x *jwxImpl) Sign(payload []byte, opts ...jws.SignOption) ([]byte, error) {
	_opts := make([]jws.SignOption, len(opts)+1)
	_opts[len(opts)] = jws.WithKey(x.sigAlg, x.key)
	return jws.Sign(payload, _opts...)
}

// VerifySecret verifies the signature of the given payload using the given
// algorithm and key.
func (x *jwxImpl) VerifySecret(payload []byte, opts ...jws.VerifyOption) ([]byte, error) {
	_opts := make([]jws.VerifyOption, len(opts)+1)
	_opts[len(opts)] = jws.WithKey(x.sigAlg, x.key)
	return jws.Verify(payload, _opts...)
}

// VerifySignature verifies the signature of the given payload using the given
// algorithm and key.
func (x *jwxImpl) VerifyJWS(payload []byte, opts ...jws.VerifyOption) ([]byte, error) {
	_opts := make([]jws.VerifyOption, len(opts)+1)
	_opts[len(opts)] = jws.WithKey(x.sigAlg, x.key)
	return jws.Verify(payload, _opts...)
}
