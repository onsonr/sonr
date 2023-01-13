package types

import (
	"crypto/rand"
	"testing"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestWebauthnVM(t *testing.T) {
	_, pub, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	pbBz, err := pub.Raw()
	if err != nil {
		t.Fatal(err)
	}

	// Create example Webauthn Credential
	webauthnCredential := &common.WebauthnCredential{
		Id:              []byte("test"),
		PublicKey:       pbBz,
		AttestationType: "platform",
		Authenticator: &common.WebauthnAuthenticator{
			Aaguid:       []byte("test"),
			CloneWarning: true,
			SignCount:    1,
		},
		Transport: []string{"usb"},
	}

	vm, err := NewWebAuthnVM(webauthnCredential)
	if err != nil {
		t.Fatal(err)
	}

	wcm, err := vm.WebAuthnCredential()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, webauthnCredential, wcm)
}
