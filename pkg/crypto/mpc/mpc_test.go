package mpc_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan"

	"github.com/yourusername/yourproject/pkg/crypto/mpc"
)

func TestKeyshareGeneration(t *testing.T) {
	// Test data
	testData := []byte("hello world")
	
	// Generate keyshares
	shares, err := GenerateKeyshares()
	require.NoError(t, err)
	require.Len(t, shares, 2)

	// Test signing with keyshares
	sig, err := SignData(shares, testData)
	require.NoError(t, err)
	require.NotNil(t, sig)

	// Verify signature
	valid, err := VerifySignature(shares[0].GetPublicKey(), testData, sig)
	require.NoError(t, err)
	assert.True(t, valid)
}

func TestKeyshareSource(t *testing.T) {
	// Generate keyshares
	shares, err := GenerateKeyshares()
	require.NoError(t, err)

	// Create keyshare source
	source, err := mpc.KeyshareSourceFromArray(shares)
	require.NoError(t, err)

	// Test source properties
	assert.NotEmpty(t, source.Address())
	assert.NotEmpty(t, source.Issuer())

	// Create default token
	token, err := source.DefaultOriginToken()
	require.NoError(t, err)
	require.NotNil(t, token)

	// Verify token properties
	parser := source.TokenParser()
	parsed, err := parser.Parse(token.Raw)
	require.NoError(t, err)
	assert.Equal(t, source.Issuer(), parsed.Issuer)
}

func TestAttenuatedTokens(t *testing.T) {
	shares, err := GenerateKeyshares()
	require.NoError(t, err)

	source, err := mpc.KeyshareSourceFromArray(shares)
	require.NoError(t, err)

	// Create parent token with full capabilities
	caps := mpc.NewSmartAccountCapabilities()
	parentAtt := mpc.CreateSmartAccountAttenuations(caps, source.Address())
	
	now := time.Now()
	parentToken, err := source.NewOriginToken(
		source.Issuer(),
		parentAtt,
		nil,
		now,
		now.Add(24*time.Hour),
	)
	require.NoError(t, err)

	// Create attenuated token with reduced capabilities
	reducedAtt := mpc.CreatePolicyAttenuation(
		caps,
		source.Address(),
		mpc.POLICY_THRESHOLD,
	)

	childToken, err := source.NewAttenuatedToken(
		parentToken,
		source.Issuer(),
		reducedAtt,
		nil,
		now,
		now.Add(12*time.Hour),
	)
	require.NoError(t, err)

	// Verify attenuated token
	parser := source.TokenParser()
	parsed, err := parser.Parse(childToken.Raw)
	require.NoError(t, err)

	// Verify reduced capabilities
	assert.Len(t, parsed.Attenuations, len(reducedAtt))
	assert.True(t, parsed.Attenuations.Contains(reducedAtt))
}

// Helper functions

func GenerateKeyshares() ([]mpc.Share, error) {
	// Implementation would depend on your actual keyshare generation logic
	// This is just a placeholder
	return nil, nil
}

func SignData(shares []mpc.Share, data []byte) ([]byte, error) {
	// Implementation would depend on your actual signing logic
	// This is just a placeholder
	return nil, nil
}

func VerifySignature(pubKey []byte, data []byte, signature []byte) (bool, error) {
	// Implementation would depend on your actual verification logic
	// This is just a placeholder
	return false, nil
}
