package bip32_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/di-dao/sonr/crypto/bip32"
	"github.com/di-dao/sonr/crypto/mpc"
	"github.com/di-dao/sonr/x/did/types"
	"github.com/stretchr/testify/require"
)

// TestComputePublicKey tests the ComputePublicKey function.
func TestComputePublicKey(t *testing.T) {
	// Create Controller
	keyset, err := mpc.GenerateKss()
	require.NoError(t, err)
	pub := keyset.PublicKey()
	pbz := pub.Bytes()
	coins := types.DefaultCoins()
	for _, c := range coins {
		pub, err := bip32.ComputePublicKey(pbz, c.GetPath(), 0)
		require.NoError(t, err)
		phex := hex.EncodeToString(pub)
		fmt.Printf("Coin: %s, Path: %d, Index: %d, PublicKey: %s \n", c.Name, c.Path, c.Index, phex)
	}
}
