package jwx

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Document(t *testing.T) {

	t.Run("can create JWK with public key", func(t *testing.T) {
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
	})
}
