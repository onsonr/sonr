package v2

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// Define the BIP32 path constants
const (
	rootFolder = "m"
	purpose    = 44
)

// getPrefix builds the Bip32 path constants into a filepath string and returns it.
func getBip32Prefix(ct crypto.CoinType) string {
	return filepath.Join(rootFolder, fmt.Sprintf("%d'", purpose), fmt.Sprintf("%d", ct.BipPath()))
}

type FileStore struct {
	basePath string
}

func NewFileStore(p string) (*FileStore, error) {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	ds := &FileStore{
		basePath: p,
	}
	return ds, nil
}

// ListAccountNames returns a list of all account names in the BIP32 file system.
func (fs *FileStore) ListAccountsForToken(coinType crypto.CoinType) ([]Account, error) {
	dir := filepath.Join(fs.basePath, getBip32Prefix(coinType))
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory %s doesn't exist", dir)
	}

	var accounts []Account
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return nil
		}
		acc, _ := NewWalletAccount(filepath.Join(dir, info.Name()))
		if acc != nil {
			accounts = append(accounts, acc)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// GetAccountNames returns a list of all account names for each supported coin in the BIP32 file system.
func (fs *FileStore) ListAccounts() (map[crypto.CoinType][]Account, error) {
	accountNames := make(map[crypto.CoinType][]Account)
	for _, ct := range crypto.AllCoinTypes() {
		names, err := fs.ListAccountsForToken(ct)
		if err != nil {
			return nil, err
		}
		accountNames[ct] = names
	}
	return accountNames, nil
}

// WriteCmpConfig writes a CmpConfig to the BIP32 file system.
func (fs *FileStore) WriteCmpConfigs(ct crypto.CoinType, cmpConfigs []*cmp.Config) (Account, error) {
	// Get the path for the account directory
	accountDir, err := getNextAccountDir(fs.basePath, ct)
	if err != nil {
		return nil, fmt.Errorf("failed to get account directory for coin type %s: %w", ct.Ticker(), err)
	}

	// Create the account directory if it doesn't exist
	if _, err := os.Stat(accountDir); os.IsNotExist(err) {
		if err := os.Mkdir(accountDir, os.ModePerm); err != nil {
			return nil, fmt.Errorf("failed to create account directory for coin type %s: %w", ct.Ticker(), err)
		}
	}

	// Write each CmpConfig to the account directory
	for _, conf := range cmpConfigs {
		// Create or truncate the file at the specified path
		file, err := os.Create(filepath.Join(accountDir, fmt.Sprintf("%s.key", conf.ID)))
		if err != nil {
			return nil, err
		}
		defer file.Close()

		// Call MarshalBinary on the CmpConfig to get the binary representation
		binaryBytes, err := conf.MarshalBinary()
		if err != nil {
			return nil, err
		}
		_, err = file.Write(binaryBytes)
		if err != nil {
			return nil, err
		}
	}
	return NewWalletAccount(accountDir)
}

// Read CmpConfig reads a CmpConfig from the specified file path in the BIP32 file
// system.
func ReadCmpConfig(path string) (*cmp.Config, error) {
	// Open the file at the specified path
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the binary-encoded CmpConfig from the file
	cmpConfig := cmp.EmptyConfig(curve.Secp256k1{})
	binaryBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Call UnmarshalBinary on the CmpConfig to decode the binary representation
	err = cmpConfig.UnmarshalBinary(binaryBytes)
	if err != nil {
		return nil, err
	}
	return cmpConfig, nil
}

func getNextAccountDir(basePath string, coinType crypto.CoinType) (string, error) {
	// Create the parent directories for the file path
	err := os.MkdirAll(filepath.Join(basePath, getBip32Prefix(coinType)), os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("could not create directory for coin type %d: %w", coinType, err)
	}

	// Check if the directory for the coin type exists, create it if it doesn't
	coinTypeDir := filepath.Join(basePath, getBip32Prefix(coinType))
	if _, err := os.Stat(coinTypeDir); os.IsNotExist(err) {
		if err := os.Mkdir(coinTypeDir, os.ModePerm); err != nil {
			return "", fmt.Errorf("could not create directory for coin type %d: %w", coinType, err)
		}
	}
	// Find the next available account number for the coin type
	nextAccountNum := 1
	err = filepath.Walk(coinTypeDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && filepath.Base(path) == fmt.Sprintf("account%d", nextAccountNum) {
			nextAccountNum++
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("could not find next available account number: %w", err)
	}

	// Return the path for the new account
	accountPath := filepath.Join(coinTypeDir, fmt.Sprintf("account%d", nextAccountNum))
	return accountPath, nil
}
