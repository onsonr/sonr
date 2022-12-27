package mpc

import (
	"testing"

	"github.com/sonr-hq/sonr/pkg/wallet"
	"github.com/stretchr/testify/assert"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
)

// The default shards that are added to the MPC wallet
var defaultTestParticipants = party.IDSlice{"vault", "current"}

func TestCMPKeygenOffline(t *testing.T) {
	w := Initialize()
	net := createOfflineNetwork(defaultTestParticipants)
	c, err := w.Keygen("vault", net)
	assert.NoError(t, err, "wallet generation succeeds")
	assert.Contains(t, c.Address(), "snr", "address is valid")
}

func TestCMPRefreshOffline(t *testing.T) {
	p := Initialize()
	net := createOfflineNetwork(defaultTestParticipants)
	ws, err := p.Keygen("current", net)
	assert.NoError(t, err, "keygen succeeds")
	ws2, err := p.Refresh("current", net)
	assert.NoError(t, err, "refresh succeeds")
	assert.Equal(t, ws.Address(), ws2.Address(), "refreshed wallet has same address")
}

func TestCMPSignOffline(t *testing.T) {
	m := []byte("sign this message")
	w := Initialize()
	net := createOfflineNetwork(defaultTestParticipants)
	ws, err := w.Keygen("current", net)
	assert.NoError(t, err, "wallet generation succeeds")
	sig, err := w.Sign("current", m, net)
	assert.NoError(t, err, "signing succeeds")
	bz, err := wallet.SerializeSignature(sig)
	assert.NoError(t, err, "signature serialization succeeds")
	deserializedSigVerified := ws.Verify(m, bz)
	assert.True(t, deserializedSigVerified, "deserialized signature is verified")
}
