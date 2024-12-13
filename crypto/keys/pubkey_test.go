package keys

import (
	"crypto/rand"
	"testing"

	"github.com/onsonr/sonr/crypto/core/curves"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPubKey_Creation(t *testing.T) {
	curve := curves.K256()
	_, pub := curve.GenerateKey(rand.Reader)
	
	pubKey := NewPubKey(pub)
	assert.NotNil(t, pubKey)
	assert.Equal(t, "secp256k1", pubKey.Type())
	assert.NotEmpty(t, pubKey.Hex())
	assert.NotEmpty(t, pubKey.Bytes())
}

func TestPubKey_Verification(t *testing.T) {
	curve := curves.K256()
	priv, pub := curve.GenerateKey(rand.Reader)
	
	pubKey := NewPubKey(pub)
	
	// Test message
	msg := []byte("test message")
	
	// Sign message
	sig, err := priv.Sign(msg)
	require.NoError(t, err)
	
	// Verify signature
	valid, err := pubKey.Verify(msg, sig)
	require.NoError(t, err)
	assert.True(t, valid)
	
	// Test invalid signature
	invalidSig := make([]byte, len(sig))
	copy(invalidSig, sig)
	invalidSig[0] ^= 0xff // Flip some bits
	
	valid, err = pubKey.Verify(msg, invalidSig)
	require.NoError(t, err)
	assert.False(t, valid)
}

func TestPubKey_InvalidSignature(t *testing.T) {
	curve := curves.K256()
	_, pub := curve.GenerateKey(rand.Reader)
	
	pubKey := NewPubKey(pub)
	
	// Test with invalid signature length
	_, err := pubKey.Verify([]byte("test"), []byte("invalid"))
	assert.Error(t, err)
}
