package mpc

import (
	"strings"
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

	t.Run("Address and Public Key", func(t *testing.T) {
		enclave, err := GenEnclave()
		require.NoError(t, err)

		// Test Address
		addr := enclave.Address()
		assert.NotEmpty(t, addr)
		assert.True(t, strings.HasPrefix(addr, "idx"))

		// Test Public Key
		pubKey := enclave.PubKey()
		assert.NotNil(t, pubKey)
		assert.NotEmpty(t, pubKey.Bytes())
	})

	t.Run("Refresh Operation", func(t *testing.T) {
		enclave, err := GenEnclave()
		require.NoError(t, err)

		// Test refresh
		refreshedEnclave, err := enclave.Refresh()
		require.NoError(t, err)
		require.NotNil(t, refreshedEnclave)
		
		// Verify refreshed enclave is valid
		assert.True(t, refreshedEnclave.IsValid())
		
		// Verify it maintains the same address
		assert.Equal(t, enclave.Address(), refreshedEnclave.Address())
	})
}

func TestEnclaveSerialization(t *testing.T) {
	t.Run("Marshal and Unmarshal", func(t *testing.T) {
		// Generate original enclave
		original, err := GenEnclave()
		require.NoError(t, err)
		require.NotNil(t, original)

		// Marshal
		keyEnclave, ok := original.(*KeyEnclave)
		require.True(t, ok)
		
		data, err := keyEnclave.Marshal()
		require.NoError(t, err)
		require.NotEmpty(t, data)

		// Unmarshal
		restored := &KeyEnclave{}
		err = restored.Unmarshal(data)
		require.NoError(t, err)

		// Verify restored enclave
		assert.Equal(t, keyEnclave.Addr, restored.Addr)
		assert.True(t, keyEnclave.PubPoint.Equal(restored.PubPoint))
		assert.Equal(t, keyEnclave.VaultCID, restored.VaultCID)
		assert.True(t, restored.IsValid())
	})
}
