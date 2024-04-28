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
