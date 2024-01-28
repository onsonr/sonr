package daed_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sonrhq/sonr/pkg/crypto/daed"
)

func TestNewKeyset(t *testing.T) {
	kh, err := daed.NewKeyHandle()
	assert.NoError(t, err)
	assert.NotNil(t, kh)
}

func TestEncryptDecrypt(t *testing.T) {
	msg := []byte("hello world")
	associatedData := []byte("associated data")
	kh, err := daed.NewKeyHandle()
	assert.NoError(t, err)
	assert.NotNil(t, kh)

	ciphertext, err := daed.Encrypt(kh, msg, associatedData)
	assert.NoError(t, err)
	assert.NotNil(t, ciphertext)

	plaintext, err := daed.Decrypt(kh, ciphertext, associatedData)
	assert.NoError(t, err)
	assert.NotNil(t, plaintext)

	assert.Equal(t, msg, plaintext)
}
