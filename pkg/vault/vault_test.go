package vault_test

import (
	"context"
	"testing"

	"github.com/ipfs/kubo/client/rpc"
	"github.com/stretchr/testify/require"

	"github.com/sonrhq/sonr/pkg/vault"
)

func TestNewVault(t *testing.T) {
	if !checkLocalIPFSConn() {
		t.Skip("Skipping test due to no local IPFS connection")
	}
	v, err := vault.New(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, v.Key)
	t.Logf("PeerID: %s", v.PeerID)
	t.Logf("SonrAddress: %s", v.SonrAddress)
}

func checkLocalIPFSConn() bool {
	_, err := rpc.NewLocalApi()
	return err == nil
}
