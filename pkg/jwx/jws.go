package jwx

import (
	"github.com/lestrrat-go/jwx/v2/jws"
)

func (x *jwxImpl) Sign(payload []byte, opts ...jws.SignOption) ([]byte, error) {
	_opts := make([]jws.SignOption, len(opts)+1)
	_opts[len(opts)] = jws.WithKey(x.sigAlg, x.key)
	return jws.Sign(payload, _opts...)
}
