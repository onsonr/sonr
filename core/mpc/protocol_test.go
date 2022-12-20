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

	p1, err := Initialize(n1)
	assert.NoError(t, err)

	n2, err := node.New(context.TODO())
	assert.NoError(t, err)

	p2, err := Initialize(n2)
	assert.NoError(t, err)

	w1, err := p1.HostDoernerKeygen(n2.ID())
	assert.NoError(t, err)

	go p2.JoinDoernerKeygen(n1.ID())
	assert.NoError(t, err)

	t.Log(w1.Public)
}
