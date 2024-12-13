package mpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeyShareGeneration(t *testing.T) {
	t.Run("Generate Valid Enclave", func(t *testing.T) {
		// Generate enclave with both user and validator shares
		enclave, err := GenEnclave()
		require.NoError(t, err)
		require.NotNil(t, enclave)

		// Validate user share exists and has correct role
		userShare, ok := enclave[kUserEnclaveKey].(KeyShare)
		require.True(t, ok, "user share should exist")
		assert.Equal(t, RoleUser, userShare.Role())

		// Validate validator share exists and has correct role
		valShare, ok := enclave[kValEnclaveKey].(KeyShare)
		require.True(t, ok, "validator share should exist")
		assert.Equal(t, RoleValidator, valShare.Role())

		// Validate address and public key
		assert.NotEmpty(t, enclave.Address())
		assert.NotNil(t, enclave.PubKey())
	})
}

func TestKeyShareRoles(t *testing.T) {
	t.Run("Role Determination", func(t *testing.T) {
		// Generate valid shares first
		enclave, err := GenEnclave()
		require.NoError(t, err)

		userShare := enclave[kUserEnclaveKey].(KeyShare)
		valShare := enclave[kValEnclaveKey].(KeyShare)

		tests := []struct {
			name     string
			share    KeyShare
			expected Role
		}{
			{"User Share", userShare, RoleUser},
			{"Validator Share", valShare, RoleValidator},
			{"Invalid Share", KeyShare("invalid.data"), RoleUnknown},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				role := tt.share.Role()
				assert.Equal(t, tt.expected, role)
			})
		}
	})
}

func TestKeyShareEncoding(t *testing.T) {
	t.Run("Encode/Decode Valid Keyshares", func(t *testing.T) {
		// Generate valid enclave first
		enclave, err := GenEnclave()
		require.NoError(t, err)

		// Test both user and validator shares
		shares := map[string]KeyShare{
			"User Share":      enclave[kUserEnclaveKey].(KeyShare),
			"Validator Share": enclave[kValEnclaveKey].(KeyShare),
		}

		for name, share := range shares {
			t.Run(name, func(t *testing.T) {
				// Get original role
				originalRole := share.Role()
				require.NotEqual(t, RoleUnknown, originalRole)

				// Test decoding
				decoded, err := DecodeKeyshare(share.String())
				require.NoError(t, err)
				assert.Equal(t, share, decoded)
				assert.Equal(t, originalRole, decoded.Role())

				// Verify message can be extracted
				msg, err := share.Message()
				require.NoError(t, err)
				require.NotNil(t, msg)
			})
		}
	})

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

func TestEnclaveOperations(t *testing.T) {
	t.Run("Signing Operations", func(t *testing.T) {
		// Generate valid enclave
		enclave, err := GenEnclave()
		require.NoError(t, err)

		// Verify both shares exist
		_, ok := enclave[kUserEnclaveKey].(KeyShare)
		require.True(t, ok, "user share should exist")
		_, ok = enclave[kValEnclaveKey].(KeyShare)
		require.True(t, ok, "validator share should exist")

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
