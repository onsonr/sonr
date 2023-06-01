package sfs

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"github.com/sonrhq/core/x/vault/types"
)

// The function encrypts data using the AES encryption algorithm with a given key.
func encryptAES(key, data []byte) ([]byte, error) {
    blockCipher, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(blockCipher)
    if err != nil {
        return nil, err
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err = rand.Read(nonce); err != nil {
        return nil, err
    }

    ciphertext := gcm.Seal(nonce, nonce, data, nil)

    return ciphertext, nil
}

// The function decrypts data using the AES algorithm with a given key.
func decryptAES(key, data []byte) ([]byte, error) {
    blockCipher, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(blockCipher)
    if err != nil {
        return nil, err
    }

    nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]

    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return nil, err
    }
    return plaintext, nil
}

// The function inserts a keyshare into a table and returns an error if there is one.
func insertEncKeyshare(ks types.KeyShare, secret_key []byte) {
	dat := ks.Bytes()
	encDat, err := encryptAES(secret_key, dat)
	if err != nil {
		return
	}
	_, err = ksTable.Put(ctx, types.KeysharePrefix(ks.Did()), encDat)
	if err != nil {
		return
	}
	return 
}

// The function retrieves a keyshare from a table using the keyshare's DID and returns it as a
// model.
func getEncKeyshare(ksDid string, secret_key []byte) (types.KeyShare, error) {
	ksr, err := types.ParseKeyShareDID(ksDid)
	if err != nil {
		return nil, err
	}
	vEnc, err := ksTable.Get(ctx, types.KeysharePrefix(ksDid))
	if err != nil {
		return nil, err
	}
	vBiz, err := decryptAES(secret_key, vEnc)
	if err != nil {
		return nil, err
	}

	ks, err := types.NewKeyshare(ksDid, vBiz, ksr.CoinType)
	if err != nil {
		return nil, err
	}
	return ks, nil
}