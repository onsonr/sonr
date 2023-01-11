package fs

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/sonr-hq/sonr/pkg/common"
	"golang.org/x/crypto/scrypt"
)

// Storing the share of the wallet.
func (c *Config) StoreShare(share []byte, partyId string, password string) error {
	// Verify WalletConfigShare
	shareConfig := &common.WalletShareConfig{}
	err := shareConfig.Unmarshal(share)
	if err != nil {
		return err
	}

	encShare, err := AesEncryptWithPassword(password, share)
	if err != nil {
		return err
	}

	// Create path for file to be stored and write file
	path := filepath.Join(c.localPath, "_auth", partyId)
	err = os.WriteFile(path, encShare, 0644)
	if err != nil {
		return err
	}

	// Add file to IPFS
	cid, err := c.ipfs.Unixfs().Add(c.ctx, c.rootNode, options.Unixfs.Pin(true))
	if err != nil {
		return err
	}
	c.ipfsPath = cid
	return nil
}

// Loading the shares from the local file system.
func (c *Config) LoadShares(password string) ([]*common.WalletShareConfig, error) {
	shares := []*common.WalletShareConfig{}
	// List all files in the _auth directory
	files, err := os.ReadDir(filepath.Join(c.localPath, "_auth"))
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, errors.New("No shares found")
	}

	// Iterate over all files
	for _, file := range files {
		// Read file
		bz, err := os.ReadFile(filepath.Join(c.localPath, "_auth", file.Name()))
		if err != nil {
			return nil, err
		}

		decBz, err := AesDecryptWithPassword(password, bz)
		if err != nil {
			return nil, err
		}

		// Unmarshal share
		share := &common.WalletShareConfig{}
		err = share.Unmarshal(decBz)
		if err != nil {
			continue
		}
		// Add share to list
		shares = append(shares, share)
	}
	return shares, nil
}

// aesDecryptWithKey uses the give 32-bit key to decrypt plaintext.
func aesDecryptWithKey(aesKey, ciphertext []byte) ([]byte, error) {
	if len(aesKey) != 32 {
		fmt.Printf("aesKey len: %d\n", len(aesKey))
		return nil, errors.New("AES key must be 32 bytes")
	}

	blockCipher, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}

	nonce, ct := ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():]

	plaintext, err := gcm.Open(nil, nonce, ct, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// AesEncryptWithPassword uses the give password to generate an aes key and decrypt plaintext.
func AesEncryptWithPassword(password string, plaintext []byte) ([]byte, error) {
	key, err := deriveKey(password)
	if err != nil {
		return nil, err
	}

	return aesEncryptWithKey(key, plaintext)
}

// AesDecryptWithPassword uses the give password to generate an aes key and encrypt plaintext.
func AesDecryptWithPassword(password string, ciphertext []byte) ([]byte, error) {
	key, err := deriveKey(password)
	if err != nil {
		return nil, err
	}

	return aesDecryptWithKey(key, ciphertext)
}

func deriveKey(password string) ([]byte, error) {
	// including a salt would make it impossible to reliably login from other devices
	key, err := scrypt.Key([]byte(password), []byte(""), 1<<20, 8, 1, 32)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// NewAesKey generates a new 32-bit key.
func NewAesKey() ([]byte, error) {
	key := make([]byte, 32)
	if n, err := rand.Read(key); err != nil {
		return nil, err
	} else if n != 32 {
		return nil, errors.New("could not create key at 32 bytes")
	}

	return key, nil
}

// aesEncryptWithKey uses the give 32-bit key to encrypt plaintext.
func aesEncryptWithKey(aesKey, plaintext []byte) ([]byte, error) {
	if len(aesKey) != 32 {
		return nil, errors.New("AES key must be 32 bytes")
	}

	blockCipher, err := aes.NewCipher(aesKey)
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

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	return ciphertext, nil
}
