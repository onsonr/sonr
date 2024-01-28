package daed

import (
	"github.com/tink-crypto/tink-go/v2/daead"
	"github.com/tink-crypto/tink-go/v2/keyset"
)

// NewKeyHandle creates a new keyset and uses it to encrypt and decrypt a message
func NewKeyHandle() (*keyset.Handle, error) {
	return keyset.NewHandle(daead.AESSIVKeyTemplate())
}

// Encrypt takes a keyset handle, plaintext, and associated data and returns ciphertext
func Encrypt(kh *keyset.Handle, plaintext []byte, associatedData []byte) ([]byte, error) {
	d, err := daead.New(kh)
	if err != nil {
		return nil, err
	}

	ciphertext, err := d.EncryptDeterministically(plaintext, associatedData)
	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}

// Decrypt takes a keyset handle, ciphertext, and associated data and returns plaintext
func Decrypt(kh *keyset.Handle, ciphertext []byte, associatedData []byte) ([]byte, error) {
	d, err := daead.New(kh)
	if err != nil {
		return nil, err
	}

	decrypted, err := d.DecryptDeterministically(ciphertext, associatedData)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}
