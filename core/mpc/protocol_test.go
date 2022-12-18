package mpc

import (
	"context"
	"testing"

	"github.com/sonr-hq/sonr/pkg/node"
	"github.com/zeebo/assert"
)

func TestJoinCMPKeygen(t *testing.T) {
	n1, err := node.New(context.TODO())
	assert.NoError(t, err)

	n2, err := node.New(context.TODO())
	assert.NoError(t, err)

	err = n1.Connect(n2.AddrInfo())
	assert.NoError(t, err)

	p, err := Initialize(n1)
	assert.NoError(t, err)

	_, err = p.JoinCMPKeygen(n2.ID())
}
