package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/di-dao/core/x/did/keeper"
	"github.com/di-dao/core/x/did/types"
)

func TestControllerSigning(t *testing.T) {
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
}

func TestAddressConversion(t *testing.T) {
	// Create user and validator keyshares
	kss, err := keeper.GenerateKSS()
	require.NoError(t, err)

	pub := kss.PublicKey()
	require.NotNil(t, pub)

	// Test address conversion
	btcAddr, err := types.CreateBitcoinAddress(pub)
	require.NoError(t, err)
	t.Logf("Bitcoin address: %s", btcAddr)
	err = btcAddr.Validate()
	require.NoError(t, err)

	ethAddr, err := types.CreateEthereumAddress(pub)
	require.NoError(t, err)
	t.Logf("Ethereum address: %s", ethAddr)
	err = ethAddr.Validate()
	require.NoError(t, err)

	// Test address validation
	snrAddr, err := types.CreateSonrAddress(pub)
	require.NoError(t, err)
	t.Logf("Sonr address: %s", snrAddr)
}

func TestLinkUnlinkProperty(t *testing.T) {
	kss, err := keeper.GenerateKSS()
	require.NoError(t, err)
	ctrl, err := keeper.CreateController(kss)
	require.NoError(t, err)

	// Link a property
	propertyKey := "email"
	propertyValue := "user@example.com"
	witness, err := ctrl.Link(propertyKey, propertyValue)
	require.NoError(t, err)
	require.NotEmpty(t, witness)

	// Validate the linked property
	valid := ctrl.Validate(propertyKey, witness)
	require.True(t, valid)

	// Unlink the property
	err = ctrl.Unlink(propertyKey, propertyValue)
	require.NoError(t, err)

	// Validate the unlinked property
	valid = ctrl.Validate(propertyKey, witness)
	require.False(t, valid)
}

func TestUnlinkNonExistentProperty(t *testing.T) {
	kss, err := keeper.GenerateKSS()
	require.NoError(t, err)
	ctrl, err := keeper.CreateController(kss)
	require.NoError(t, err)

	// Unlink a non-existent property
	propertyKey := "non_existent"
	propertyValue := "value"
	err = ctrl.Unlink(propertyKey, propertyValue)
	require.Error(t, err)
}
