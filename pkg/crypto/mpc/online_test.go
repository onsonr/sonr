package mpc

import (
	"context"
	"testing"

	"github.com/sonr-hq/sonr/pkg/node"
	"github.com/stretchr/testify/assert"
)

func TestCMPKeygenOnline(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create two nodes
	n1, err := node.New(ctx, node.WithPartyId("vault"))
	assert.NoError(t, err, "node creation succeeds")

	n2, err := node.New(ctx, node.WithPartyId("current"))
	assert.NoError(t, err, "node creation succeeds")

	// Connect the nodes
	err = n2.Connect(n1.MultiAddr())
	assert.NoError(t, err, "node connection succeeds")

	// Create MPC Protocol
	w := Initialize()
	net, err := createOnlineNetwork(ctx, n1, n2)
	assert.NoError(t, err, "network creation succeeds")

	// Run MPC Protocol
	c, err := w.Keygen("current", net)
	assert.NoError(t, err, "wallet generation succeeds")

	// Check that the address is valid
	assert.Contains(t, c.Address(), "snr", "address is valid")
}
