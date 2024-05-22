package local_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	daead "github.com/di-dao/sonr/internal/local"
)

func TestNewKeyset(t *testing.T) {
	kh, err := daead.NewKeyHandle()
	assert.NoError(t, err)
	assert.NotNil(t, kh)
}

func TestEncryptDecrypt(t *testing.T) {
	msg := []byte("hello world")
	associatedData := []byte("associated data")
	kh, err := daead.NewKeyHandle()
	assert.NoError(t, err)
	assert.NotNil(t, kh)

	ciphertext, err := daead.Encrypt(msg, associatedData)
	assert.NoError(t, err)
	assert.NotNil(t, ciphertext)

	plaintext, err := daead.Decrypt(ciphertext, associatedData)
	assert.NoError(t, err)
	assert.NotNil(t, plaintext)

	assert.Equal(t, msg, plaintext)
}
