package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EncryptWithPassword(t *testing.T) {
	t.Run("encrypts and decrypts successfully", func(t *testing.T) {
		ciphertext, err := EncryptWithPassword("password", []byte("mycontent"))
		assert.NoError(t, err, "encryption succeeds")

		plaintext, err := DecryptWithPassword("password", []byte(ciphertext))
		assert.NoError(t, err, "decryption succeeds")

		assert.Equal(t, "mycontent", string(plaintext), "plaintext matches")
	})

	t.Run("fails to decrypt with wrong password", func(t *testing.T) {
		ciphertext, err := EncryptWithPassword("password", []byte("mycontent"))
		assert.NoError(t, err, "encryption succeeds")

		_, err = DecryptWithPassword("wrongpassword", []byte(ciphertext))
		assert.ErrorContains(t, err, "cipher: message authentication failed", "decryption fails")
	})
}
