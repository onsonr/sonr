package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MPCCreate(t *testing.T) {
	_, err := Generate()
	assert.NoError(t, err, "wallet generation succeeds")
}

func Test_MPCDID(t *testing.T) {
	w, err := Generate()
	assert.NoError(t, err, "wallet generation succeeds")

	_, err = w.Bech32Address()
	assert.NoError(t, err, "Bech32Address successfully created")

	_, err = w.DIDDocument()
	assert.NoError(t, err, "DID Document creation succeeds")
}

func Test_MPCSignMessage(t *testing.T) {
	w, err := Generate()
	assert.NoError(t, err, "wallet generation succeeds")

	sig, err := w.Sign([]byte("sign this message"))
	assert.NoError(t, err, "signing succeeds")

	buf := SerializeECDSA(sig)
	assert.NotNil(t, buf, "serialization succeeds")

	sig2, err := ParseDERSignature(buf)
	assert.NoError(t, err, "deserialization succeeds")
	assert.Equal(t, sig, sig2, "valid signature deserialization")
}

// func Test_MPCCreateWhoIs(t *testing.T) {
// 	w, err := Generate()
// 	assert.NoError(t, err, "wallet generation succeeds")
// 	addr, err := w.Bech32Address()
// 	assert.NoError(t, err, "Bech32Address successfully created")
// 	err = client.RequestFaucet(addr)
// 	assert.NoError(t, err, "faucet request succeeds")
// 	resp := w.Balances()
// 	t.Logf("-- Get Balances --\n%+v\n", resp)

// 	err = w.BroadcastCreateWhoIsWithEncoding()
// 	assert.NoError(t, err, "broadcast with encoding succeeds")
// }
