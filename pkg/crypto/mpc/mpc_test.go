package mpc_test

import (
	"fmt"
	"testing"

	"github.com/onsonr/sonr/pkg/crypto/mpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeyshareGeneration(t *testing.T) {
	// Test data
	testData := []byte("hello world")

	// Generate keyshares
	src, err := mpc.NewKeyshareSource()
	require.NoError(t, err)
	src.Address()

	// Test signing with keyshares
	sig, err := src.SignData(testData)
	require.NoError(t, err)
	require.NotNil(t, sig)

	// Verify signature
	valid, err := src.VerifyData(testData, sig)
	require.NoError(t, err)
	assert.True(t, valid)
	tk, err := src.DefaultOriginToken()
	require.NoError(t, err)
	cid, err := tk.CID()
	require.NoError(t, err)
	fmt.Println(cid)
}
