package keychain_test

import (
	"context"
	"testing"

	"github.com/ipfs/kubo/client/rpc"
	"github.com/stretchr/testify/require"

	"github.com/sonrhq/sonr/internal/keychain"
)

func TestNewKeychain(t *testing.T) {
	if !checkLocalIPFSConn() {
		t.Skip("Skipping test due to no local IPFS connection")
	}
	kc, err := keychain.New(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, kc.Address)
}

func checkLocalIPFSConn() bool {
	_, err := rpc.NewLocalApi()
	return err == nil
}

func BenchmarkNewKeychain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kc, err := keychain.New(context.Background())
		if err != nil {
			b.Fail()
		}
		if len(kc.Address) == 0 {
			b.Fail()
		}
	}
}
