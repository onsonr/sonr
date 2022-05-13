package did

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JWE(t *testing.T) {
	doc := document()
	jwe, err := doc.CreateJWS([]byte(doc.VerificationMethod[0].PublicKeyBase58))

	if err != nil {
		t.Errorf("Error while creating JWT")
	}

	t.Run("Create JWT Should return defined", func(t *testing.T) {
		assert.NotNil(t, jwe)
	})
}
