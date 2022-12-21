package jwx

import "github.com/lestrrat-go/jwx/v2/jwe"

/*
	Encrypts a JWE with the internal key which is assumed to be a public key type
*/
func (x *jwxImpl) EncryptJWE(payload []byte, opts ...jwe.EncryptOption) ([]byte, error) {
	_opts := make([]jwe.EncryptOption, len(opts)+1)
	_opts[len(opts)] = jwe.WithKey(x.keyAlg, x.key)
	return jwe.Encrypt(payload, _opts...)
}

/*
	Decrypts a JWE with a given private key assumed to be the associated one with the public key held
*/
func (x *jwxImpl) DecryptJWE(payload []byte, key interface{}, opts ...jwe.DecryptOption) ([]byte, error) {
	_opts := make([]jwe.DecryptOption, len(opts)+1)
	_opts[len(opts)] = jwe.WithKey(x.keyAlg, key)
	return jwe.Decrypt(payload, _opts...)
}
