package bls_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/didao-org/sonr/pkg/wallet/bls"
)

func TestAccumulatorWitness(t *testing.T) {
	// Create a new accumulator
	secKey, err := bls.NewSecretKey(bls.RandomSeed())
	require.NoError(t, err)
	pubKey, err := secKey.PublicKey()
	require.NoError(t, err)
	acc, err := secKey.CreateAccumulator()
	require.NoError(t, err)

	// Add some values to the accumulator
	values := []string{"value1", "value2", "value3"}
	err = acc.AddValues(secKey, values...)
	require.NoError(t, err)

	// Create Witness for value1
	witness, err := acc.CreateWitness(secKey, "value1")
	require.NoError(t, err)

	// Remove value1 from the accumulator
	err = acc.RemoveValues(secKey, "value1")
	require.NoError(t, err)

	// Verify that the witness is no longer valid
	require.False(t, acc.VerifyElement(pubKey, witness))
}
