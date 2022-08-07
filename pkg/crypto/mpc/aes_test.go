package mpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EncryptWithPassword(t *testing.T) {
	t.Run("encrypts and decrypts successfully", func(t *testing.T) {
		ciphertext, err := AesEncryptWithPassword("password", []byte("mycontent"))
		assert.NoError(t, err, "encryption succeeds")

		plaintext, err := AesDecryptWithPassword("password", []byte(ciphertext))
		assert.NoError(t, err, "decryption succeeds")

		assert.Equal(t, "mycontent", string(plaintext), "plaintext matches")
	})

	t.Run("fails to decrypt with wrong password", func(t *testing.T) {
		ciphertext, err := AesEncryptWithPassword("password", []byte("mycontent"))
		assert.NoError(t, err, "encryption succeeds")

		_, err = AesDecryptWithPassword("wrongpassword", []byte(ciphertext))
		assert.ErrorContains(t, err, "cipher: message authentication failed", "decryption fails")
	})
}

func Test_EncryptWithKey(t *testing.T) {
	t.Run("encrypts and decrypts successfully", func(t *testing.T) {
		key, err := NewAesKey()
		assert.NoError(t, err, "generates aes key")

		ciphertext, err := AesEncryptWithKey(key, []byte("mycontent"))
		assert.NoError(t, err, "encryption succeeds")

		plaintext, err := AesDecryptWithKey(key, []byte(ciphertext))
		assert.NoError(t, err, "decryption succeeds")

		assert.Equal(t, "mycontent", string(plaintext), "plaintext matches")
	})

	t.Run("fails to decrypt with wrong key", func(t *testing.T) {
		key, err := NewAesKey()
		assert.NoError(t, err, "generates aes key")

		ciphertext, err := AesEncryptWithKey(key, []byte("mycontent"))
		assert.NoError(t, err, "encryption succeeds")

		key[0] ^= 0xFF
		_, err = AesDecryptWithKey(key, []byte(ciphertext))
		assert.ErrorContains(t, err, "cipher: message authentication failed", "decryption fails")
	})
}
