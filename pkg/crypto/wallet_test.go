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

	_, err = w.Sign([]byte("sign this message"))
	assert.NoError(t, err, "signing succeeds")
}

func Test_MPCCreateWhoIs(t *testing.T) {
	w, err := Generate()
	assert.NoError(t, err, "wallet generation succeeds")

	err = w.BroadcastCreateWhoIs()
	assert.NoError(t, err, "broadcast succeeds")
}
