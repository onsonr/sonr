package keychain_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sonrhq/sonr/internal/keychain"
)

func TestNewKeychain(t *testing.T) {
	kc, err := keychain.New(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, kc.Address)
}
