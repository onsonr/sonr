package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
)

func Test_MPCCreate(t *testing.T) {
	pl := pool.NewPool(0)
	defer pl.TearDown()

	// Generate method currently creates a wallet and tests signing a message. This is a bit of a hack, due to Sign not working independtly of a wallet.
	_, err := Generate()
	if err != nil {
		t.Error(err)
	}
}

func Test_MPCDID(t *testing.T) {
	w, err := Generate()
	if err != nil {
		t.Error(err)
		return
	}

	pub, err := w.Bech32Address()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success", pub)

	doc, err := w.DIDDocument()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success", doc)
}

func Test_MPCSignMessage(t *testing.T) {
	w, err := Generate()
	assert.NoError(t, err, "wallet generation succeeds")

	// pl := pool.NewPool(0)
	// defer pl.TearDown()

	sig, err := w.Sign([]byte("sign this message"))
	assert.NoError(t, err, "signing succeeds")
	fmt.Printf("sig: %+v\n", sig)
}

func Test_MPCCreateWhoIs(t *testing.T) {
	w, err := Generate()
	assert.NoError(t, err, "wallet generation succeeds")

	// pl := pool.NewPool(0)
	// defer pl.TearDown()

	err = w.BroadcastCreateWhoIs()
	assert.NoError(t, err, "signing succeeds")
}
