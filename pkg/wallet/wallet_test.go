package wallet_test

import (
	"log"
	"testing"

	"github.com/di-dao/sonr/crypto/mpc"
	wallet "github.com/di-dao/sonr/pkg/wallet"
	"github.com/stretchr/testify/require"
)

// TestNewWallet tests the ComputePublicKey function.
func TestNewWallet(t *testing.T) {
	// Create Controller
	keyset, err := mpc.GenerateKss()
	require.NoError(t, err)

	w, err := wallet.NewWallet(keyset)
	require.NoError(t, err)
	require.NotNil(t, w)
	t.Log(w)

	log.Printf("BTC Address: %v", w.Accounts[0][0].Address)
	log.Printf("ETH Address: %v", w.Accounts[60][0].Address)
	log.Printf("SNR Address: %v", w.Accounts[703][0].Address)
}
