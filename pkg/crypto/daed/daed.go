package daed

import (
	"bytes"
	"fmt"

	"github.com/tink-crypto/tink-go/v2/daead"
	"github.com/tink-crypto/tink-go/v2/insecurecleartextkeyset"
	"github.com/tink-crypto/tink-go/v2/keyset"
)

func NewKeyset() error {
	jsonKeyset := `{
                        "key": [{
                                "keyData": {
                                                "keyMaterialType":
                                                                "SYMMETRIC",
                                                "typeUrl":
                                                                "type.googleapis.com/google.crypto.tink.AesSivKey",
                                                "value":
                                                                "EkAl9HCMmKTN1p3V186uhZpJQ+tivyc4IKyE+opg6SsEbWQ/WesWHzwCRrlgRuxdaggvgMzwWhjPnkk9gptBnGLK"
                                },
                                "keyId": 1919301694,
                                "outputPrefixType": "TINK",
                                "status": "ENABLED"
                }],
                "primaryKeyId": 1919301694
        }`

	// Create a keyset handle from the cleartext keyset in the previous
	// step. The keyset handle provides abstract access to the underlying keyset to
	// limit the exposure of accessing the raw key material. WARNING: In practice,
	// it is unlikely you will want to use a insecurecleartextkeyset, as it implies
	// that your key material is passed in cleartext, which is a security risk.
	// Consider encrypting it with a remote key in Cloud KMS, AWS KMS or HashiCorp Vault.
	// See https://github.com/google/tink/blob/master/docs/GOLANG-HOWTO.md#storing-and-loading-existing-keysets.
	keysetHandle, err := insecurecleartextkeyset.Read(
		keyset.NewJSONReader(bytes.NewBufferString(jsonKeyset)))
	if err != nil {
		return err
	}

	// Retrieve the DAEAD primitive we want to use from the keyset handle.
	primitive, err := daead.New(keysetHandle)
	if err != nil {
		return err
	}

	// Use the primitive to encrypt a message. In this case the primary key of the
	// keyset will be used (which is also the only key in this example).
	plaintext := []byte("message")
	associatedData := []byte("associated data")
	ciphertext, err := primitive.EncryptDeterministically(plaintext, associatedData)
	if err != nil {
		return err
	}

	// Use the primitive to decrypt the message. Decrypt finds the correct key in
	// the keyset and decrypts the ciphertext. If no key is found or decryption
	// fails, it returns an error.
	decrypted, err := primitive.DecryptDeterministically(ciphertext, associatedData)
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
