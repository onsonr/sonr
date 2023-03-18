package wallet

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// KeyShare is a type that interacts with a cmp.Config file located on disk.
type KeyShare interface {
	// Path returns the path to the file.
	Path() string

	// Config returns the cmp.Config.
	Config() *cmp.Config

	// CoinType returns the coin type based on the account directories parent
	CoinType() crypto.CoinType

	// AccountName returns the account name based on the account directory name
	AccountName() string

	// Encrypt checks if the file at current path is encrypted and if not, encrypts it.
	Encrypt(credential *crypto.WebauthnCredential) error

	// Encrypt checks if the file at current path is encrypted and if not, encrypts it.
	Decrypt(credential *crypto.WebauthnCredential) error

	IsEncrypted() bool
}

type keyShare struct {
	cnfg *cmp.Config
	p    string
}

// NewKeyshare creates a new KeyShare.
func NewKeyshare(path string) (KeyShare, error) {
	cnfg := cmp.EmptyConfig(curve.Secp256k1{})
	bz, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}
	if err := cnfg.UnmarshalBinary(bz); err != nil {
		return nil, fmt.Errorf("could not unmarshal cmp config: %w", err)
	}
	return &keyShare{
		cnfg: cnfg,
		p:    path,
	}, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                Filesystem & I/O                                ||
// ! ||--------------------------------------------------------------------------------||

// AccountName returns the account name based on the account directory name
func (s *keyShare) AccountName() string {
	return filepath.Dir(s.p)
}

// Name returns the name of the file.
func (s *keyShare) Name() string {
	return filepath.Base(s.p)
}

// Path returns the path to the file.
func (s *keyShare) Path() string {
	return s.p
}

// Config returns the cmp.Config.
func (s *keyShare) Config() *cmp.Config {
	return s.cnfg
}

// CoinType returns the coin type based on the account directories parent
func (s *keyShare) CoinType() crypto.CoinType {
	coinDir := filepath.Base(filepath.Dir(s.p))
	allCoins := crypto.AllCoinTypes()
	for _, coin := range allCoins {
		if strings.Contains(coinDir, fmt.Sprintf("%d", coin.BipPath())) {
			return coin
		}
	}
	return crypto.TestCoinType
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                    Webauthn                                    ||
// ! ||--------------------------------------------------------------------------------||

// Encrypt checks if the file at current path is encrypted and if not, encrypts it.
func (s *keyShare) Encrypt(credential *crypto.WebauthnCredential) error {
	if s.Name() == "vault" {
		return nil
	}
	bz, err := s.cnfg.MarshalBinary()
	if err != nil {
		return err
	}
	encBz, err := credential.Encrypt(bz)
	if err != nil {
		return err
	}
	return os.WriteFile(s.p, encBz, 0600)
}

// Decrypt checks if the file at current path is encrypted and if not, encrypts it.
func (s *keyShare) Decrypt(credential *crypto.WebauthnCredential) error {
	if s.Name() == "vault" {
		return nil
	}
	bz, err := os.ReadFile(s.p)
	if err != nil {
		return err
	}
	decBz, err := credential.Decrypt(bz)
	if err != nil {
		return err
	}
	cnfg := cmp.EmptyConfig(curve.Secp256k1{})
	if err := cnfg.UnmarshalBinary(decBz); err != nil {
		return err
	}
	s.cnfg = cnfg
	return nil
}

// IsEncrypted checks if the file at current path is encrypted.
func (s *keyShare) IsEncrypted() bool {
	if s.Name() == "vault" {
		return false
	}
	bz, err := os.ReadFile(s.p)
	if err != nil {
		return false
	}

	cnfg := cmp.EmptyConfig(curve.Secp256k1{})
	if err := cnfg.UnmarshalBinary(bz); err != nil {
		return true
	}
	return false
}
