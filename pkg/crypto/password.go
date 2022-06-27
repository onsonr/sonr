package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"golang.org/x/crypto/scrypt"
)

func EncryptWithPassword(password, plaintext []byte) (string, error) {
	key, err := deriveKey(password)
	if err != nil {
		return "", err
	}

	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	return string(ciphertext), nil
}

func DecryptWithPassword(password, ciphertext []byte) (string, error) {
	key, err := deriveKey(password)
	if err != nil {
		return "", err
	}

	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return "", err
	}

	nonce, ct := ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():]

	plaintext, err := gcm.Open(nil, nonce, ct, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func deriveKey(password []byte) ([]byte, error) {
	// including a salt would make it impossible to reliably login from other devices
	key, err := scrypt.Key(password, []byte(""), 1<<20, 8, 1, 32)
	if err != nil {
		return nil, err
	}

	return key, nil
}
