package jwx

import "github.com/lestrrat-go/jwx/v2/jwe"

func (x *jwxImpl) EncryptJWE(payload []byte, opts ...jwe.EncryptOption) ([]byte, error) {
	_opts := make([]jwe.EncryptOption, len(opts)+1)
	_opts[len(opts)] = jwe.WithKey(x.sigAlg, x.key)
	jwe.Encrypt(payload, _opts...)
}

func (x *jwxImpl) DecryptJWE(payload []byte, opts ...jwe.DecryptOption) ([]byte, error) {
	_opts := make([]jwe.DecryptOption, len(opts)+1)
	_opts[len(opts)] = jwe.WithKey(x.sigAlg, x.key)
	jwe.Decrypt(payload, _opts...)
}
