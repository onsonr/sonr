package motor

import (
	"github.com/sonr-io/sonr/pkg/crypto/jwx"
)

/*
	Encrypt Content generates a symetric encryption key from the motor's `encryptionKey`. and creates an encryption key
	in the format of `JSON Web Keys (JWK)`. The private key can always be derived from the address to create the key pair.

	This will be done when wishing to decrypt content from this wallet address.

	Returns

	- []byte -> encrypted jwk encrypted by the wallet address.

	- []byte -> encrypted content by the jwk which is encrypted and returned

	- error => returned if there is an error in generating keys and encrypting content, other values will be nil
*/
func (mtr *motorNodeImpl) EncryptContent(data []byte) ([]byte, []byte, error) {
	x := jwx.New()
	x.SetKey(mtr.encryptionKey)

	_, err := x.CreateEncJWK()

	if err != nil {
		return nil, nil, err
	}
	// encrypt our content with the generate key
	encContent, err := x.EncryptJWE(data)

	if err != nil {
		return nil, nil, err
	}

	// here we encrypt the created jwe
	pk, err := mtr.Wallet.CreateEcdsaFromPublicKey()

	if err != nil {
		return nil, nil, err
	}

	j := jwx.New()
	j.SetKey(pk.PublicKey)
	_, err = j.CreateEncJWK()

	if err != nil {
		return nil, nil, err
	}

	keyBytes, err := x.MarshallJSON()

	if err != nil {
		return nil, nil, err
	}

	encyptedKey, err := j.EncryptJWE(keyBytes)

	if err != nil {
		return nil, nil, err
	}

	return encyptedKey, encContent, nil
}
