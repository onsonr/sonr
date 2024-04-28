package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/di-dao/core/x/did/keeper"
	"github.com/di-dao/core/x/did/types"
)

func TestController(t *testing.T) {
	// Create user and validator keyshares
	kss, err := keeper.GenerateKSS()
	require.NoError(t, err)

	pubKey := kss.PublicKey()
	require.NotNil(t, pubKey)

	// Create controller
	ctrl, err := keeper.CreateController(kss)
	require.NoError(t, err)
	require.NotNil(t, ctrl)

	// Test signing and verifying a message
	msg := []byte("test message")
	sig, err := ctrl.Sign(msg)
	require.NoError(t, err)
	require.NotEmpty(t, sig)

	valid := pubKey.VerifySignature(msg, sig)
	require.True(t, valid)

	// Test refreshing the keyshares
	err = ctrl.Refresh()
	require.NoError(t, err)

	// Test getting the public key

	require.NotNil(t, pubKey)
	require.IsType(t, &types.PublicKey{}, pubKey)
}
