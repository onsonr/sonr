package vault_test

import (
	"context"
	"testing"

	"github.com/ipfs/kubo/client/rpc"
	"github.com/stretchr/testify/require"

	"github.com/sonrhq/sonr/pkg/vault"
)

func TestNewController(t *testing.T) {
	if !checkLocalIPFSConn() {
		t.Skip("Skipping test due to no local IPFS connection")
	}
	v, err := vault.Create(context.Background())
	require.NoError(t, err)
	t.Logf("Controller: %s", v.String())
}

func checkLocalIPFSConn() bool {
	_, err := rpc.NewLocalApi()
	return err == nil
}
