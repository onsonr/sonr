package mpc

import (
	"context"
	"testing"

	"github.com/sonr-hq/sonr/pkg/wallet"
	"github.com/stretchr/testify/assert"
)

func Test_MPCCreate(t *testing.T) {
	w := NewProtocol(context.Background())
	c, err := w.Keygen("vault")
	assert.NoError(t, err, "wallet generation succeeds")
	_, err = c.PublicKey()
	assert.NoError(t, err, "public key creation succeeds")
}

func Test_MPCDID(t *testing.T) {
	p := NewProtocol(context.Background())
	_, err := p.Keygen("current")
	assert.NoError(t, err, "keygen succeeds")
}

func Test_MPCSignMessage(t *testing.T) {
	m := []byte("sign this message")
	w := NewProtocol(context.Background())
	ws, err := w.Keygen("current")
	assert.NoError(t, err, "wallet generation succeeds")

	sig, err := w.Sign(m)
	assert.NoError(t, err, "signing succeeds")
	bz, err := wallet.SerializeSignature(sig)
	assert.NoError(t, err, "signature serialization succeeds")
	deserializedSigVerified := ws.Verify(m, bz)
	assert.True(t, deserializedSigVerified, "deserialized signature is verified")
}
