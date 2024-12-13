package mpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeyShareGeneration(t *testing.T) {
	t.Run("Generate Valid Enclave", func(t *testing.T) {
		// Generate enclave
		enclave, err := GenEnclave()
		require.NoError(t, err)
		require.NotNil(t, enclave)

		// Validate enclave contents
		assert.True(t, enclave.IsValid())
	})
}

func TestEnclaveOperations(t *testing.T) {
	t.Run("Signing and Verification", func(t *testing.T) {
		// Generate valid enclave
		enclave, err := GenEnclave()
		require.NoError(t, err)

		// Test signing
		testData := []byte("test message")
		signature, err := enclave.Sign(testData)
		require.NoError(t, err)
		require.NotNil(t, signature)

		// Verify the signature
		valid, err := enclave.Verify(testData, signature)
		require.NoError(t, err)
		assert.True(t, valid)

		// Test invalid data verification
		invalidData := []byte("wrong message")
		valid, err = enclave.Verify(invalidData, signature)
		require.NoError(t, err)
		assert.False(t, valid)
	})
}
