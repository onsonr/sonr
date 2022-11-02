package jwx

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwe"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/third_party/types/common"
	"github.com/stretchr/testify/assert"
)

func Test_JWK(t *testing.T) {
	t.Run("can create JWK from wallet keys", func(t *testing.T) {
		w, err := mpc.GenerateWallet(common.DefaultCallback())
		assert.NoError(t, err)
		p, err := w.PublicKey()

		// need to pad the key to 40 bytes for
		// ecdsa key generation which sizes its buffers from
		for len(p) < 40 {
			p = append(p, 0)
		}
		assert.NoError(t, err)
		key, err := ecdsa.GenerateKey(elliptic.P256(), bytes.NewReader(p))
		assert.NoError(t, err)

		x := New()
		x.SetKey(key)
		x.SetKeyAlgo(jwa.ECDH_ES_A256KW)

		_, err = x.CreateSignJWK()

		assert.NoError(t, err)

		msg := "my message to sign"

		b, err := x.Sign([]byte(msg), key)

		assert.NoError(t, err)
		assert.NotNil(t, b)

		m, err := x.MarshallJSON()
		assert.NoError(t, err)
		assert.NotNil(t, m)
	})

	t.Run("can create symetric key from aes FOR ENC", func(t *testing.T) {
		aes, err := mpc.NewAesKey()
		assert.NoError(t, err)

		x := New()
		x.SetKey(aes)
		x.CreateEncJWK()

		message := "hello world"

		enc, err := x.EncryptJWE([]byte(message))

		assert.NoError(t, err)
		assert.NotNil(t, enc)

		decrypt, err := x.DecryptJWE(enc, aes)
		assert.NoError(t, err)

		b, err := x.MarshallJSON()
		assert.NoError(t, err)
		x.UnmarshallJSON(b)

		assert.Equal(t, message, string(decrypt))
	})

	t.Run("can create symetric key from aes FOR SIGN", func(t *testing.T) {
		aes, err := mpc.NewAesKey()
		assert.NoError(t, err)
		sk, err := mpc.NewEcdsaFromAes(aes)
		assert.NoError(t, err)

		x := New()
		x.SetKey(sk.PublicKey)
		x.CreateSignJWK()

		message := "hello world"

		enc, err := x.Sign([]byte(message), sk)
		assert.NoError(t, err)
		assert.NotNil(t, enc)

		d, err := x.VerifyJWS(enc)
		assert.NoError(t, err)
		assert.NotNil(t, d)

		decrypt, err := x.VerifySecret(enc)
		assert.NoError(t, err)

		b, err := x.MarshallJSON()
		assert.NoError(t, err)
		x.UnmarshallJSON(b)

		assert.Equal(t, message, string(decrypt))
	})

	t.Run("can create JWK for encryption with public key", func(t *testing.T) {
		key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			t.Errorf("Error while convering credential to public key %s", err)
		}

		x := New()
		x.SetKey(key)
		x.SetKeyAlgo(jwa.ECDH_ES_A256KW)

		jwk, err := x.CreateEncJWK()
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		assert.NotNil(t, jwk)
		assert.Equal(t, jwk.KeyUsage(), "enc")
	})

	t.Run("can create JWK for signature with public key", func(t *testing.T) {
		key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			t.Errorf("Error while convering credential to public key %s", err)
		}

		x := New()
		x.SetKey(key.PublicKey)
		jwk, err := x.CreateSignJWK()
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		assert.NotNil(t, jwk)
		assert.Equal(t, jwk.KeyUsage(), "sig")
	})

	t.Run("can marshall jwk", func(t *testing.T) {
		key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			t.Errorf("Error while convering credential to public key %s", err)
		}

		x := New()
		x.SetKey(key.PublicKey)
		x.SetKeyAlgo(jwa.ECDH_ES_A256KW)

		jwk, err := x.CreateSignJWK()

		assert.NoError(t, err)
		assert.NotNil(t, jwk)
		data, err := x.MarshallJSON()
		assert.NoError(t, err, "marshall succeeds")

		assert.NotNil(t, data)
		assert.True(t, len(data) > 0)
	})

	t.Run("can unmarshall jwk", func(t *testing.T) {
		key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			t.Errorf("Error while convering credential to public key %s", err)
		}

		x := New()
		x.SetKey(key.PublicKey)
		x.SetKeyAlgo(jwa.ECDH_ES_A256KW)

		jwk, err := x.CreateSignJWK()

		assert.NoError(t, err)
		assert.NotNil(t, jwk)
		data, err := x.MarshallJSON()
		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.True(t, len(data) > 0)

		keyAsJson, err := x.UnmarshallJSON(data)

		assert.NoError(t, err)
		assert.NotNil(t, keyAsJson)
	})

	t.Run("can encrypt jwe", func(t *testing.T) {
		key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			t.Errorf("Error while convering credential to public key %s", err)
		}

		x := New()
		x.SetKey(key.PublicKey)
		x.SetKeyAlgo(jwa.ECDH_ES_A256KW)
		jwk, err := x.CreateSignJWK()

		assert.NoError(t, err)
		assert.NotNil(t, jwk)

		test := "this is a message"
		test_encoded := []byte(test)
		opts := []jwe.EncryptOption{}

		payload, err := x.EncryptJWE(test_encoded, opts...)

		assert.NoError(t, err)
		assert.NotNil(t, payload)
	})

	t.Run("can decrypt jwe", func(t *testing.T) {
		key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			t.Errorf("Error while convering credential to public key %s", err)
		}

		x := New()
		x.SetKey(key.PublicKey)
		x.SetKeyAlgo(jwa.ECDH_ES_A256KW)

		jwk, err := x.CreateSignJWK()

		assert.NoError(t, err)
		assert.NotNil(t, jwk)

		test := "this is a message"
		test_encoded := []byte(test)
		e_opts := []jwe.EncryptOption{}

		payload, err := x.EncryptJWE(test_encoded, e_opts...)

		assert.NoError(t, err)
		assert.NotNil(t, payload)

		d_opts := []jwe.DecryptOption{}
		d_payload, err := x.DecryptJWE(payload, key, d_opts...)
		assert.NoError(t, err)
		assert.NotNil(t, payload)

		assert.Equal(t, test, string(d_payload))
	})
}
