package mpc

import (
	"crypto/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func randNonce() []byte {
	nonce := make([]byte, 12)
	rand.Read(nonce)
	return nonce
}

func TestKeyShareGeneration(t *testing.T) {
	t.Run("Generate Valid Enclave", func(t *testing.T) {
		nonce := randNonce()
		// Generate enclave
		enclave, err := GenEnclave(nonce)
		require.NoError(t, err)
		require.NotNil(t, enclave)

		// Validate enclave contents
		assert.True(t, enclave.IsValid())
	})

	t.Run("Export and Import", func(t *testing.T) {
		nonce := randNonce()
		// Generate original enclave
		original, err := GenEnclave(nonce)
		require.NoError(t, err)

		// Test key for encryption/decryption
		testKey := []byte("test-key-12345678")

		// Test Export/Import for Validator role
		t.Run("Validator Role", func(t *testing.T) {
			// Export validator share
			valData, err := original.Export(RoleVal, testKey)
			require.NoError(t, err)
			require.NotEmpty(t, valData)

			// Create new empty enclave
			newEnclave, err := GenEnclave(nonce)
			require.NoError(t, err)

			// Import validator share
			err = newEnclave.Import(RoleVal, valData, testKey)
			require.NoError(t, err)

			// Verify the imported share works by signing
			testData := []byte("test message")
			sig, err := newEnclave.Sign(testData)
			require.NoError(t, err)
			valid, err := newEnclave.Verify(testData, sig)
			require.NoError(t, err)
			assert.True(t, valid)
		})

		// Test Export/Import for User role
		t.Run("User Role", func(t *testing.T) {
			// Export user share
			userData, err := original.Export(RoleUser, testKey)
			require.NoError(t, err)
			require.NotEmpty(t, userData)

			// Create new empty enclave
			newEnclave, err := GenEnclave(nonce)
			require.NoError(t, err)

			// Import user share
			err = newEnclave.Import(RoleUser, userData, testKey)
			require.NoError(t, err)

			// Verify the imported share works by signing
			testData := []byte("test message")
			sig, err := newEnclave.Sign(testData)
			require.NoError(t, err)
			valid, err := newEnclave.Verify(testData, sig)
			require.NoError(t, err)
			assert.True(t, valid)
		})

		// Test Invalid Role
		t.Run("Invalid Role", func(t *testing.T) {
			invalidRole := Role("invalid")
			_, err := original.Export(invalidRole, testKey)
			assert.Error(t, err)

			err = original.Import(invalidRole, []byte("test"), testKey)
			assert.Error(t, err)
		})

		// Test Invalid Key
		t.Run("Invalid Key", func(t *testing.T) {
			valData, err := original.Export(RoleVal, testKey)
			require.NoError(t, err)

			wrongKey := []byte("wrong-key-12345678")
			err = original.Import(RoleVal, valData, wrongKey)
			assert.Error(t, err)
		})
	})
}

func TestEnclaveOperations(t *testing.T) {
	t.Run("Signing and Verification", func(t *testing.T) {
		nonce := randNonce()
		// Generate valid enclave
		enclave, err := GenEnclave(nonce)
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
		nonce := randNonce()
		enclave, err := GenEnclave(nonce)
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
		nonce := randNonce()
		enclave, err := GenEnclave(nonce)
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
		nonce := randNonce()
		// Generate original enclave
		original, err := GenEnclave(nonce)
		require.NoError(t, err)
		require.NotNil(t, original)

		// Marshal
		keyclave, ok := original.(*keyEnclave)
		require.True(t, ok)

		data, err := keyclave.Serialize()
		require.NoError(t, err)
		require.NotEmpty(t, data)

		// Unmarshal
		restored := &keyEnclave{}
		err = restored.Unmarshal(data)
		require.NoError(t, err)

		// Verify restored enclave
		assert.Equal(t, keyclave.Addr, restored.Addr)
		assert.True(t, keyclave.PubPoint.Equal(restored.PubPoint))
		assert.True(t, restored.IsValid())
	})
}
