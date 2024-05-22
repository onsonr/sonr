package controller_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/di-dao/core/crypto/mpc"
	"github.com/di-dao/core/internal/controller"
)

func TestControllerSigning(t *testing.T) {
	// Create user and validator keyshares
	kss, err := mpc.GenerateKss()
	require.NoError(t, err)

	pubKey := kss.PublicKey()
	require.NotNil(t, pubKey)

	// Create controller
	ctrl, err := controller.Create(kss)
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
}

func TestAddressConversion(t *testing.T) {
	// Create user and validator keyshares
	kss, err := mpc.GenerateKss()
	require.NoError(t, err)

	pub := kss.PublicKey()
	require.NotNil(t, pub)
}

func TestLinkUnlinkProperty(t *testing.T) {
	// Property key
	propertyKey := "email"

	// Create controller
	kss, err := mpc.GenerateKss()
	require.NoError(t, err)
	ctrl, err := controller.Create(kss)
	require.NoError(t, err)

	// Link a property
	propertyValue := "user@example.com"
	witness, err := ctrl.Set(propertyKey, propertyValue)
	require.NoError(t, err)
	require.NotEmpty(t, witness)

	// Validate the linked property
	valid := ctrl.Check(propertyKey, witness)
	require.True(t, valid)

	// Unlink the property
	err = ctrl.Remove(propertyKey, propertyValue)
	require.NoError(t, err)

	// Validate the unlinked property
	valid = ctrl.Check(propertyKey, witness)
	require.False(t, valid)
}

func TestUnlinkNonExistentProperty(t *testing.T) {
	kss, err := mpc.GenerateKss()
	require.NoError(t, err)
	ctrl, err := controller.Create(kss)
	require.NoError(t, err)

	// Unlink a non-existent property
	propertyKey := "non_existent"
	propertyValue := "value"
	err = ctrl.Remove(propertyKey, propertyValue)
	require.Error(t, err)
}

func TestSignRefreshSign(t *testing.T) {
	// Create controller
	kss, err := mpc.GenerateKss()
	require.NoError(t, err)
	ctrl, err := controller.Create(kss)
	require.NoError(t, err)

	// Sign a message
	msg := []byte("test message")
	sig1, err := ctrl.Sign(msg)
	require.NoError(t, err)
	require.NotEmpty(t, sig1)

	// Refresh the controller
	err = ctrl.Refresh()
	require.NoError(t, err)

	// Sign the same message again
	sig2, err := ctrl.Sign(msg)
	require.NoError(t, err)
	require.NotEmpty(t, sig2)

	// Signatures should be different
	require.NotEqual(t, sig1, sig2)
}
