package jwx

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JWK(t *testing.T) {

	t.Run("can create JWK for encryption with public key", func(t *testing.T) {
		key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			t.Errorf("Error while convering credential to public key %s", err)
		}
		pk := key.Public()
		jwk, err := CreateJWKForEnc(pk)
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
		pk := key.Public()
		jwk, err := CreateJWKForSig(pk)
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
		pk := key.Public()
		jwk, err := CreateJWKForSig(pk)
		assert.NoError(t, err)
		data, err := Marshall(&jwk)

		assert.NotNil(t, data)
		assert.True(t, len(data) > 0)
	})

	t.Run("can unmarshall jwk", func(t *testing.T) {
		key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			t.Errorf("Error while convering credential to public key %s", err)
		}
		pk := key.Public()
		jwk, err := CreateJWKForSig(pk)
		assert.NoError(t, err)
		data, err := Marshall(&jwk)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.True(t, len(data) > 0)

		keyAsJson, err := Unmarshall(data)

		assert.NoError(t, err)
		assert.NotNil(t, keyAsJson)
	})
}
