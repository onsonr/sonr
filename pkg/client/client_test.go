package client

import (
	"fmt"
	"testing"

	"github.com/sonr-io/sonr/pkg/crypto"

	"github.com/stretchr/testify/assert"
)

func Test_FaucetCheckBalance(t *testing.T) {
	client := NewClient(true)
	w, err := crypto.Generate()
	assert.NoError(t, err, "wallet generation succeeds")
	addr, err := w.Address()
	assert.NoError(t, err, "Bech32Address successfully created")
	fmt.Println("Address:", addr)
	resp, err := client.CheckBalance(addr)
	assert.NoError(t, err, "Check Balance succeeds")
	t.Logf("-- Get Balances (1) --\n%+v\n", resp)

	err = client.RequestFaucet(addr)
	assert.NoError(t, err, "faucet request succeeds")
	resp2, err := client.CheckBalance(addr)
	assert.NoError(t, err, "Check Balance succeeds")
	t.Logf("-- Get Balances (2) --\n%+v\n", resp2)
}
