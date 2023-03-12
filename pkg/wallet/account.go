package v2

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/crypto/mpc"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// Account is an interface for an account in the wallet
type Account interface {
	// Address returns the address of the account based on the coin type
	Address() string

	// CoinType returns the coin type of the account
	CoinType() crypto.CoinType

	// DID returns the DID of the account
	DID() string

	// ListKeyshares returns a list of keyshares for the account
	ListKeyshares() ([]KeyShare, error)

	// Name returns the name of the account
	Name() string

	// Path returns the path of the account
	Path() string

	// PartyIDs returns the party IDs of the account
	PartyIDs() []crypto.PartyID

	// PubKey returns secp256k1 public key
	PubKey() *crypto.PubKey

	// Rename renames the account
	Rename(name string) error

	// Signs a message
	Sign(bz []byte) ([]byte, error)

	// Type returns the type of the account
	Type() string

	// Verifies a signature
	Verify(bz []byte, sig []byte) (bool, error)
}

type walletAccount struct {
	p string
}

// NewWalletAccount loads an accound directory and returns a WalletAccount
func NewWalletAccount(p string) (Account, error) {
	// Check if the path is a directory
	if !isDir(p) {
		return nil, fmt.Errorf("path %s is not a directory", p)
	}
	return &walletAccount{p: p}, nil
}

// Address returns the address of the account based on the coin type
func (wa *walletAccount) Address() string {
	// if wa.CoinType().IsEthereum() {
	// 	return wa.PubKey().ETHAddress()
	// }
	// addr, err := wa.PubKey().Bech32(wa.CoinType().AddrPrefix())
	// if err != nil {
	// 	return ""
	// }
	// return addr
	addr, _ := wa.PubKey().Bech32(wa.CoinType().AddrPrefix())
	return addr
}

// CoinType returns the coin type of the account
func (wa *walletAccount) CoinType() crypto.CoinType {
	coinBipStr := filepath.Base(filepath.Dir(wa.p))
	coinBip, err := strconv.Atoi(coinBipStr)
	if err != nil {
		return crypto.TestCoinType
	}
	return crypto.CoinTypeFromBipPath(int32(coinBip))
}

// DID returns the DID of the account
func (wa *walletAccount) DID() string {
	if wa.CoinType().IsSonr() {
		return fmt.Sprintf("did:%s:%s", wa.CoinType().DidMethod(), wa.Address())
	}
	return fmt.Sprintf("did:%s:%s#%s", wa.CoinType().DidMethod(), wa.Address(), wa.Name())
}

// ListKeyshares returns a list of keyshares for the account
func (wa *walletAccount) ListKeyshares() ([]KeyShare, error) {
	files, err := os.ReadDir(wa.p)
	if err != nil {
		return nil, err
	}
	var keyshares []KeyShare
	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".key" {
			ks, err := NewKeyshare(filepath.Join(wa.p, f.Name()))
			if err != nil {
				return nil, err
			}
			keyshares = append(keyshares, ks)
		}
	}
	return keyshares, nil
}

// Name returns the name of the account
func (wa *walletAccount) Name() string {
	return filepath.Base(wa.p)
}

// PartyIDs returns the party IDs of the account
func (wa *walletAccount) PartyIDs() []crypto.PartyID {
	files, err := os.ReadDir(wa.p)
	if err != nil {
		return nil
	}
	var partyIDs []crypto.PartyID
	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".key" {
			id := strings.TrimRight(f.Name(), ".key")
			partyIDs = append(partyIDs, crypto.PartyID(id))
		}
	}
	return partyIDs
}

// Path returns the path of the account
func (wa *walletAccount) Path() string {
	return wa.p
}

// PubKey returns secp256k1 public key
func (wa *walletAccount) PubKey() *crypto.PubKey {
	files, err := os.ReadDir(wa.p)
	if err != nil {
		return nil
	}
	ks, err := NewKeyshare(filepath.Join(wa.p, files[0].Name()))
	if err != nil {
		return nil
	}
	skPP, ok := ks.Config().PublicPoint().(*curve.Secp256k1Point)
	if !ok {
		return nil
	}
	bz, err := skPP.MarshalBinary()
	if err != nil {
		return nil
	}
	return crypto.NewSecp256k1PubKey(bz)
}

// Rename renames the account
func (wa *walletAccount) Rename(name string) error {
	parentDir := filepath.Dir(wa.p)
	newPath := filepath.Join(parentDir, name)

	// Rename the directory to the new name
	return os.Rename(wa.p, newPath)
}

// Signs a message using the account
func (wa *walletAccount) Sign(bz []byte) ([]byte, error) {
	kss, err := wa.ListKeyshares()
	if err != nil {
		return nil, err
	}
	var configs []*cmp.Config
	for _, ks := range kss {
		configs = append(configs, ks.Config())
	}
	return mpc.SignCMP(configs, bz, wa.PartyIDs())
}

// Type returns the type of the account
func (wa *walletAccount) Type() string {
	return fmt.Sprintf("%s/ecdsa-secp256k1", wa.CoinType().Name())
}

// Verifies a signature
func (wa *walletAccount) Verify(bz []byte, sig []byte) (bool, error) {
	kss, err := wa.ListKeyshares()
	if err != nil {
		return false, err
	}
	return mpc.VerifyCMP(kss[0].Config(), bz, sig)
}

//
// Helper functions
//

// isDir checks if the path is a directory and contains at least one MPC shard file
func isDir(p string) bool {
	fi, err := os.Stat(p)
	if err != nil {
		return false
	}
	if !fi.IsDir() {
		return false
	}
	return true
}
