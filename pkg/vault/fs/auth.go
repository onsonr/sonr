package fs

import (
	"crypto/rand"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/sonr-hq/sonr/pkg/common"
	"golang.org/x/crypto/nacl/secretbox"
)

// Storing the share of the wallet.
func (c *Config) StoreShare(share []byte, partyId string, encryptKey []byte) error {
	// Verify WalletConfigShare
	shareConfig := &common.WalletShareConfig{}
	err := shareConfig.Unmarshal(share)
	if err != nil {
		return err
	}

	encShare, err := encryptData(share, encryptKey)
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

func (c *Config) LoadShares(encryptKey []byte) ([]*common.WalletShareConfig, error) {
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

		decBz, err := decryptData(bz, encryptKey)
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

func encryptData(data []byte, secretKeyBytes []byte) ([]byte, error) {
	var secretKey [32]byte
	copy(secretKey[:], secretKeyBytes)

	// Random Nonce for every Message
	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		return nil, err
	}

	encrypted := secretbox.Seal(nonce[:], data, &nonce, &secretKey)
	return encrypted, nil
}

func decryptData(encrypted []byte, secretKeyBytes []byte) ([]byte, error) {
	var secretKey [32]byte
	copy(secretKey[:], secretKeyBytes)

	var decryptNonce [24]byte
	copy(decryptNonce[:], encrypted[:24])
	decrypted, ok := secretbox.Open(nil, encrypted[24:], &decryptNonce, &secretKey)
	if !ok {
		return nil, errors.New("Error Decrypting data")
	}
	return decrypted, nil
}
