package mpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeyShareGeneration(t *testing.T) {
	t.Run("Generate Valid Enclave", func(t *testing.T) {
		enclave, err := GenEnclave()
		require.NoError(t, err)
		require.NotNil(t, enclave)

		// Validate address and public key
		assert.NotEmpty(t, enclave.Address())
		assert.NotNil(t, enclave.PubKey())
	})
}

func TestKeyShareRoles(t *testing.T) {
	t.Run("Role Determination", func(t *testing.T) {
		tests := []struct {
			input    string
			expected Role
		}{
			{"user.testdata", RoleUser},
			{"validator.testdata", RoleValidator},
			{"invalid.testdata", RoleUnknown},
		}

		for _, tt := range tests {
			t.Run(tt.input, func(t *testing.T) {
				role := determineRole(tt.input)
				assert.Equal(t, tt.expected, role)
			})
		}
	})
}

func TestKeyShareEncoding(t *testing.T) {
	t.Run("Encode/Decode Keyshare", func(t *testing.T) {
		// Generate a valid enclave first
		enclave, err := GenEnclave()
		require.NoError(t, err)

		// Test decoding an encoded share
		for _, role := range []Role{RoleUser, RoleValidator} {
			t.Run(string(role), func(t *testing.T) {
				var share KeyShare
				if role == RoleUser {
					share = enclave[kUserEnclaveKey].(KeyShare)
				} else {
					share = enclave[kValEnclaveKey].(KeyShare)
				}

				// Verify the role
				assert.Equal(t, role, share.Role())

				// Test decoding
				decoded, err := DecodeKeyshare(share.String())
				require.NoError(t, err)
				assert.Equal(t, share, decoded)
			})
		}
	})

	t.Run("Invalid Keyshare Decoding", func(t *testing.T) {
		invalidShares := []string{
			"invalid",
			"invalid.format.extra",
			"unknown.format",
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
		enclave, err := GenEnclave()
		require.NoError(t, err)

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
