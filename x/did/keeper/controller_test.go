package keeper_test

import (
	"testing"

	"github.com/di-dao/core/x/did/keeper"
	"github.com/di-dao/core/x/did/types"
	"github.com/stretchr/testify/require"
)

func TestController(t *testing.T) {
	// Create user and validator keyshares
	vks, uks, err := keeper.GenerateKSS()
	require.NoError(t, err)

	// Create controller
	ctrl, err := keeper.CreateController(uks, vks)
	require.NoError(t, err)
	require.NotNil(t, ctrl)

	// Test linking a property
	key := "email"
	value := "test@example.com"
	witness, err := ctrl.Link(key, value)
	require.NoError(t, err)
	require.NotEmpty(t, witness)

	// Test validating the linked property
	valid := ctrl.Validate(key, witness)
	require.True(t, valid)

	// Test unlinking the property
	witness, err = ctrl.Unlink(key, value)
	require.NoError(t, err)
	require.NotEmpty(t, witness)

	// Test validating the unlinked property
	valid = ctrl.Validate(key, witness)
	require.False(t, valid)

	// Test signing and verifying a message
	msg := []byte("test message")
	sig, err := ctrl.Sign(msg)
	require.NoError(t, err)
	require.NotEmpty(t, sig)

	valid = ctrl.Verify(msg, sig)
	require.True(t, valid)

	// Test refreshing the keyshares
	err = ctrl.Refresh()
	require.NoError(t, err)

	// Test getting the public key
	pubKey := ctrl.PublicKey()
	require.NotNil(t, pubKey)
	require.IsType(t, &types.PublicKey{}, pubKey)
}

func TestStartKsProtocol(t *testing.T) {
	// Create user and validator keyshares
	vks, uks, err := keeper.GenerateKSS()
	require.NoError(t, err)

	// Get refresh functions
	valRefresh, err := vks.GetRefreshFunc()
	require.NoError(t, err)
	usrRefresh, err := uks.GetRefreshFunc()
	require.NoError(t, err)

	// Test running the keyshare protocol
	err = keeper.StartKsProtocol(valRefresh, usrRefresh)
	require.NoError(t, err)
}
