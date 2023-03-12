package v2

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

// AccountName returns the account name based on the account directory name
func (s *keyShare) AccountName() string {
	return filepath.Dir(s.p)
}
