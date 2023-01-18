package network

import (
	"context"
	"testing"

	"github.com/sonr-hq/sonr/pkg/vault/internal/mpc"
	"github.com/sonr-hq/sonr/x/identity/types"
	"github.com/stretchr/testify/assert"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
)

// The default shards that are added to the MPC wallet
var defaultTestParticipants = party.IDSlice{"vault", "current"}

func TestWalletSharesFull(t *testing.T) {
	net := NewOfflineNetwork(defaultTestParticipants)
	wsl, err := mpc.Keygen("current", 1, net, "snr")
	assert.NoError(t, err, "wallet generation succeeds")
	ws := OfflineWallet(wsl)
	ws2, err := ws.Refresh("current")
	assert.NoError(t, err, "refresh succeeds")
	bz, err := ws2.Sign("current", []byte("sign this message"))
	assert.NoError(t, err, "signing succeeds")
	deserializedSigVerified := ws2.Verify([]byte("sign this message"), bz)
	assert.True(t, deserializedSigVerified, "deserialized signature is verified")
	res, err := BroadcastTx(context.Background(), bz)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Broadcast result: %s", res.String())
	// Derive the first three accounts
	conf1, err := ws.Bip32Derive(0, "snr")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("First account: %s", conf1.Address())

	conf2, err := ws.Bip32Derive(1, "0x")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Second account: %s", conf2.Address())

	conf3, err := ws.Bip32Derive(2, "btc")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Third account: %s", conf3.Address())
}

func TestFullDIDDocument(t *testing.T) {
	net := NewOfflineNetwork(defaultTestParticipants)
	wsl, err := mpc.Keygen("current", 1, net, "snr")
	assert.NoError(t, err, "wallet generation succeeds")
	ws := OfflineWallet(wsl)
	rootVm := ws.AssertionMethod()
	doc := types.BlankDocument(rootVm.ID)
	doc.AddAssertion(rootVm)
	// Derive the first three accounts
	conf1, err := ws.Bip32Derive(0, "snr")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("First account: %s", conf1.Address())
	doc.AddBlockchainAccount(conf1)

	conf2, err := ws.Bip32Derive(1, "0x")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Second account: %s", conf2.Address())
	doc.AddBlockchainAccount(conf2)

	conf3, err := ws.Bip32Derive(2, "btc")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Third account: %s", conf3.Address())
	doc.AddBlockchainAccount(conf3)

	t.Logf("DID Document: %s", doc)
}
