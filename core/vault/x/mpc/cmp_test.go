package mpc_test

import (
	"context"
	"testing"

	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-hq/sonr/core/vault/x/mpc"
	"github.com/sonr-hq/sonr/internal/node"
)

const (
	path     = "/Users/pradn/.sonr/wallet/a_sonr_id.json"
	id       = "a"
	password = "goodpassword"
)

var (
	ids = party.IDSlice{"a", "b", "c", "d", "e", "f"}
	msg = []byte("hello world")
	sig []byte
)

func TestCreateSaveLoadWallet(t *testing.T) {
	// Create 2 nodes
	n1, err := node.New(context.Background())
	check(t, err)
	n2, err := node.New(context.Background())
	check(t, err)

	// Make Node 2 join Test channel
	_, err = n2.Join("test")
	check(t, err)

	// Fetch their IDs
	id1, err := n1.ID()
	check(t, err)
	id2, err := n2.ID()
	check(t, err)

	// Create a wallet on node 1
	_, err = mpc.NewWallet(n1, node.IDSlice{id1, id2})
	check(t, err)
}

func check(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
