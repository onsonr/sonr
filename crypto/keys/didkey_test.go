package keys

import (
	"crypto/ed25519"
	"crypto/rand"
	"testing"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDID_Creation(t *testing.T) {
	_, pub, err := crypto.GenerateEd25519Key(rand.Reader)
	require.NoError(t, err)

	did, err := NewDID(pub)
	require.NoError(t, err)
	assert.NotEmpty(t, did.String())
	assert.True(t, len(did.String()) > len(KeyPrefix))
}

func TestDID_Parse(t *testing.T) {
	// Generate a test key
	_, pub, err := crypto.GenerateEd25519Key(rand.Reader)
	require.NoError(t, err)

	did, err := NewDID(pub)
	require.NoError(t, err)

	// Convert to string and parse back
	didStr := did.String()
	parsed, err := Parse(didStr)
	require.NoError(t, err)

	// Verify the parsed key matches original
	originalRaw, err := did.Raw()
	require.NoError(t, err)
	parsedRaw, err := parsed.Raw()
	require.NoError(t, err)
	assert.Equal(t, originalRaw, parsedRaw)
}

func TestDID_VerifyKey(t *testing.T) {
	// Test Ed25519
	_, pub, err := crypto.GenerateEd25519Key(rand.Reader)
	require.NoError(t, err)

	did, err := NewDID(pub)
	require.NoError(t, err)

	key, err := did.VerifyKey()
	require.NoError(t, err)
	_, ok := key.(ed25519.PublicKey)
	assert.True(t, ok)
}

func TestDID_UnsupportedKey(t *testing.T) {
	_, err := NewDID(nil)
	assert.Error(t, err)
}
