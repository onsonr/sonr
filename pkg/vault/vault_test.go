package vault_test

import (
	"context"
	"testing"

	"github.com/sonrhq/sonr/pkg/vault"
	"github.com/stretchr/testify/require"
)

func TestNewVault(t *testing.T) {
	v, err := vault.New(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, v.Key)
	t.Logf("PeerID: %s", v.PeerID)
	t.Logf("SonrAddress: %s", v.SonrAddress)
}
