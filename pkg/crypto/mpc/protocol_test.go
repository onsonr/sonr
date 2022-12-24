package mpc

import (
	"testing"

	"github.com/sonr-hq/sonr/pkg/wallet"
	"github.com/stretchr/testify/assert"
)

func TestCMPKeygen(t *testing.T) {
	w := Initialize()
	c, err := w.Keygen("vault")
	assert.NoError(t, err, "wallet generation succeeds")
	assert.Contains(t, c.Address(), "snr", "address is valid")
}

func TestCMPRefresh(t *testing.T) {
	p := Initialize()
	ws, err := p.Keygen("current")
	assert.NoError(t, err, "keygen succeeds")
	ws2, err := p.Refresh("current")
	assert.NoError(t, err, "refresh succeeds")
	assert.Equal(t, ws.Address(), ws2.Address(), "refreshed wallet has same address")
}

func TestCMPSign(t *testing.T) {
	m := []byte("sign this message")
	w := Initialize()
	ws, err := w.Keygen("current")
	assert.NoError(t, err, "wallet generation succeeds")
	sig, err := w.Sign("current", m, ws.PartyIDs())
	assert.NoError(t, err, "signing succeeds")
	bz, err := wallet.SerializeSignature(sig)
	assert.NoError(t, err, "signature serialization succeeds")
	deserializedSigVerified := ws.Verify(m, bz)
	assert.True(t, deserializedSigVerified, "deserialized signature is verified")
}
