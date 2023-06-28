package stores

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"

	"github.com/sonrhq/core/pkg/mpc"
	"github.com/sonrhq/core/x/vault/types"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
	"lukechampine.com/blake3"
)

var (
	kCategorySetKeyPrefix = "category"
	kPrimaryListKeyPrefix = "primary"
	kKeyshareMapKeyPrefix = "kss"
)

// CategorySetKeyName returns the key name for a category set. (i.e. UnclaimedWallets, ControlledWallets, PendingClaims)
func CategorySetKeyName(category string) string {
	return fmt.Sprintf("%s/%s", kCategorySetKeyPrefix, category)
}

// PrimaryListKeyName returns the key name for a primary account list. Parameter should be an idx... address.
func PrimaryListKeyName(address string) string {
	return fmt.Sprintf("%s/%s", kPrimaryListKeyPrefix, address)
}

// KeyshareMapKeyName returns the key name for a keyshare map. Parameter should be a did.
func KeyshareMapKeyName(did string) string {
	return fmt.Sprintf("%s/%s", kKeyshareMapKeyPrefix, did)
}

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

func generateEncryptionKey(shortId string, kss ...*types.VaultKeyshare) ([]byte, error) {
	sig, err := mpc.SignCMP(convertToCmpConfigs(kss...), []byte(shortId))
	if err != nil {
		return nil, err
	}
	hashDerivKey := blake3.Sum256(sig)
	return hashDerivKey[:], nil
}

func convertToCmpConfigs(kss ...*types.VaultKeyshare) (val []*cmp.Config) {
	for _, ks := range kss {
		val = append(val, ks.CMPConfig())
	}
	return val
}

