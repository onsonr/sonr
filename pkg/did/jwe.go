package did

import (
	"errors"

	jose "gopkg.in/square/go-jose.v2"
)

// EncryptJWE creates a JWE object
func (d *Document) EncryptJWE(id DID, buf []byte) (string, error) {
	vm := d.Authentication.FindByID(id)
	if vm == nil {
		return "", errors.New("Document VerificationMethod not found")
	}

	publicKey, err := vm.PublicKey()
	if err != nil {
		return "", err
	}

	// Instantiate an encrypter using RSA-OAEP with AES128-GCM. An error would
	// indicate that the selected algorithm(s) are not currently supported.
	encrypter, err := jose.NewEncrypter(jose.A128GCM, jose.Recipient{Algorithm: jose.DIRECT, Key: publicKey}, nil)
	if err != nil {
		return "", err
	}

	// Encrypt a sample plaintext. Calling the encrypter returns an encrypted
	// JWE object, which can then be serialized for output afterwards. An error
	// would indicate a problem in an underlying cryptographic primitive.
	object, err := encrypter.Encrypt(buf)
	if err != nil {
		return "", err
	}

	// Serialize the encrypted object using the full serialization format.
	// Alternatively you can also use the compact format here by calling
	// object.CompactSerialize() instead.
	serialized := object.FullSerialize()
	return serialized, nil
}

// DecryptJWE verifies the JWE and returns the buffer
func (d *Document) DecryptJWE(id DID, serial string) ([]byte, error) {
	vm := d.Authentication.FindByID(id)
	if vm == nil {
		return nil, errors.New("Document VerificationMethod not found")
	}

	publicKey, err := vm.PublicKey()
	if err != nil {
		return nil, err
	}

	// Parse the serialized, encrypted JWE object. An error would indicate that
	// the given input did not represent a valid message.
	object, err := jose.ParseEncrypted(serial)
	if err != nil {
		return nil, err
	}

	// Now we can decrypt and get back our original plaintext. An error here
	// would indicate that the message failed to decrypt, e.g. because the auth
	// tag was broken or the message was tampered with.
	decrypted, err := object.Decrypt(publicKey)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}
