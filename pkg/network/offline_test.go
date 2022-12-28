package network

import (
	"testing"

	"github.com/sonr-hq/sonr/pkg/crypto/mpc"
	"github.com/stretchr/testify/assert"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
)

// The default shards that are added to the MPC wallet
var defaultTestParticipants = party.IDSlice{"vault", "current"}

func TestWalletSharesFull(t *testing.T) {
	net := newOfflineNetwork(defaultTestParticipants)
	wsl, err := mpc.Keygen("current", 1, net, "snr")
	assert.NoError(t, err, "wallet generation succeeds")
	ws := OfflineWallet(wsl)
	ws2, err := ws.Refresh("current")
	assert.NoError(t, err, "refresh succeeds")
	bz, err := ws2.Sign("current", []byte("sign this message"))
	assert.NoError(t, err, "signing succeeds")
	deserializedSigVerified := ws2.Verify([]byte("sign this message"), bz)
	assert.True(t, deserializedSigVerified, "deserialized signature is verified")
}
