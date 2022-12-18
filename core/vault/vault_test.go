package vault

import (
	"context"
	"testing"

	"github.com/sonr-hq/sonr/pkg/node"
	"github.com/stretchr/testify/assert"
)

func TestCreateWallet(t *testing.T) {
	// Initialize Vault
	err := InitVault("/Users/pradn/.sonr/node_key.json")
	if err != nil {
		t.Error(err)
	}

	// Create a new Node and start it
	n, err := node.New(context.TODO())
	if err != nil {
		t.Error(err)
	}

	assert.NotEqual(t, n.ID(), v.P2P.ID())

	err = v.P2P.Connect(n.AddrInfo())
	assert.NoError(t, err)

	// Generate a new wallet
	w, err := GenerateWallet(n.ID())
	if err != nil {
		t.Error(err)
	}

	// Print the wallet
	t.Log(w)
}
