package did

import (
	"errors"

	"github.com/sonr-io/sonr/pkg/did/ssi"
	jose "gopkg.in/square/go-jose.v2"
)

// EncryptJWE creates a JWE object
func (d *DocumentImpl) EncryptJWE(id DID, buf []byte) (string, error) {
	var err error
	vm := d.Authentication.FindByID(id)
	if vm == nil {
		return "", errors.New("Document VerificationMethod not found")
	}

	// check type of key
	var key interface{}
	switch vm.Type {
	case ssi.JsonWebKey2020:
		key, err = vm.JWK()
		if err != nil {
			return "", err
		}
	default:
		key, err = vm.PublicKey()
		if err != nil {
			return "", err
		}
	}
	// Instantiate an encrypter using RSA-OAEP with AES128-GCM. An error would
	// indicate that the selected algorithm(s) are not currently supported.
	encrypter, err := jose.NewEncrypter(jose.A128GCM, jose.Recipient{Algorithm: jose.DIRECT, Key: key}, nil)
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
func (d *DocumentImpl) DecryptJWE(id DID, serial string) ([]byte, error) {
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
