package vault

import (
	"context"
	"testing"

	v1 "github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/node"
	"github.com/zeebo/assert"
)

func TestNewMultiWallet(t *testing.T) {
	n1, err := node.New(context.Background(), node.WithPeerType(v1.Peer_MOTOR))
	assert.NoError(t, err)

	err = Initialize()
	assert.NoError(t, err)

	w, err := NewMultiWallet(n1)
	assert.NoError(t, err)
	t.Logf("Wallet Address %s", w.Address())
}
