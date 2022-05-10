package jwt

import (
	"fmt"
	"testing"

	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

func Test_JWT(t *testing.T) {
	root := "did:snr:123456"
	ctx, err := ssi.ParseURI("https://www.w3.org/ns/did/v1")
	id, err := did.ParseDID(root)
	if err != nil {
		t.Errorf("Error while generating did %s", err)
	}

	didController, err := did.ParseDID(fmt.Sprintf("%s#test", root))
	if err != nil {
		t.Errorf("Error while creating test controller %s", err)
	}

	priv := secp256k1.GenPrivKey()
	pub := priv.PubKey()
	vm, _ := did.NewVerificationMethod(*id, ssi.ECDSASECP256K1VerificationKey2019, *didController, pub)

	t.Run("JWT creation should contain options", func(t *testing.T) {
		jwt := DefaultNew()
		assert.NotNil(t, jwt)
		assert.NotNil(t, jwt.options)
	})

	t.Run("Should generate JWT from did uri", func(t *testing.T) {
		doc := &did.Document{
			ID:      *id,
			Context: []ssi.URI{*ctx},
		}
		doc.AddController(*didController)
		doc.AddAuthenticationMethod(vm)
		jwt := DefaultNew()
		token, err := jwt.Generate(doc)

		if err != nil {
			t.Errorf("Error while generating token %s", err)
		}

		assert.NotNil(t, token)
		assert.Greater(t, len(token), 0)
	})
}
