package daed

import (
	"fmt"

	"github.com/tink-crypto/tink-go/v2/daead"
	"github.com/tink-crypto/tink-go/v2/keyset"
)

func NewKeyset() error {
	kh, err := keyset.NewHandle(daead.AESSIVKeyTemplate())
	if err != nil {
		return err
	}

	d, err := daead.New(kh)
	if err != nil {
		return err
	}

	// Use the primitive to encrypt a message. In this case the primary key of the
	// keyset will be used (which is also the only key in this example).
	plaintext := []byte("message")
	associatedData := []byte("associated data")
	ciphertext, err := d.EncryptDeterministically(plaintext, associatedData)
	if err != nil {
		return err
	}

	// Use the primitive to decrypt the message. Decrypt finds the correct key in
	// the keyset and decrypts the ciphertext. If no key is found or decryption
	// fails, it returns an error.
	decrypted, err := d.DecryptDeterministically(ciphertext, associatedData)
	if err != nil {
		return err
	}

	fmt.Println(ciphertext)
	fmt.Println(string(decrypted))
	return nil
	// Output:
	// [1 114 102 56 62 150 98 146 84 99 211 36 127 214 229 231 157 56 143 192 250 132 32 153 124 244 238 112]
	// message
}
