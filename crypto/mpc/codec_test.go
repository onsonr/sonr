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
		assert.Contains(t, enclave, kAddrEnclaveKey)
		assert.Contains(t, enclave, kPubKeyEnclaveKey)
		assert.Contains(t, enclave, kValEnclaveKey)
		assert.Contains(t, enclave, kUserEnclaveKey)
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

func TestKeyShareEncoding(t *testing.T) {
	t.Run("Invalid Keyshare Decoding", func(t *testing.T) {
		invalidShares := []string{
			"invalid",
			"invalid.format.extra",
			"unknown.format",
			"notarole.data",
		}

		for _, share := range invalidShares {
			t.Run(share, func(t *testing.T) {
				decoded, err := DecodeKeyshare(share)
				assert.Error(t, err)
				assert.Empty(t, decoded)
			})
		}
	})
}
