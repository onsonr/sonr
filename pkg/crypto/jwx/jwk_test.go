package jwx

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/lestrrat-go/jwx/v2/jwe"

	"github.com/stretchr/testify/assert"
)

func Test_JWK(t *testing.T) {
	t.Run("can create JWK for encryption with public key", func(t *testing.T) {
		key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			t.Errorf("Error while convering credential to public key %s", err)
		}

		x := New(key.Public())
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

		x := New(key.Public())
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

		x := New(key.Public())
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

		x := New(key.Public())
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

		x := New(key.Public())
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

		x := New(key.Public())
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
