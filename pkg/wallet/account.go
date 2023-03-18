package wallet

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/crypto/mpc"
	"github.com/sonrhq/core/x/identity/types"
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

	// GetAuthInfo creates an AuthInfo for a transaction
	GetAuthInfo(gas sdk.Coins) (*txtypes.AuthInfo, error)

	// ListKeyshares returns a list of keyshares for the account
	ListKeyshares() ([]KeyShare, error)

	// Name returns the name of the account
	Name() string

	// Path returns the path of the account
	Path() string

	// PartyIDs returns the party IDs of the account
	PartyIDs() []crypto.PartyID

	Nonce() uint64
	IncrementNonce()

	// PubKey returns secp256k1 public key
	PubKey() *crypto.PubKey

	// Rename renames the account
	Rename(name string) error

	// Signs a message
	Sign(bz []byte) ([]byte, error)

	// Type returns the type of the account
	Type() string

	// VerificationMethod returns the verification method for the account
	VerificationMethod(controller string) *types.VerificationMethod

	// Verifies a signature
	Verify(bz []byte, sig []byte) (bool, error)

	// Lock locks the account
	Lock(c *crypto.WebauthnCredential) error

	// Unlock unlocks the account
	Unlock(c *crypto.WebauthnCredential) error
}

type walletAccount struct {
	p string
	n uint64
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                     General                                    ||
// ! ||--------------------------------------------------------------------------------||

// NewWalletAccount loads an accound directory and returns a WalletAccount
func NewWalletAccount(p string) (Account, error) {
	// Check if the path is a directory
	if !isDir(p) {
		return nil, fmt.Errorf("path %s is not a directory", p)
	}
	return &walletAccount{p: p}, nil
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

// Verifies a signature using first unlocked keyshare
func (wa *walletAccount) Verify(bz []byte, sig []byte) (bool, error) {
	kss, err := wa.ListKeyshares()
	if err != nil {
		return false, err
	}

	// Find first unlocked keyshare
	var uks KeyShare
	for _, ks := range kss {
		if ks.IsEncrypted() {
			continue
		}
		uks = ks
		break
	}
	if uks == nil {
		return false, fmt.Errorf("no unlocked keyshares")
	}
	return mpc.VerifyCMP(uks.Config(), bz, sig)
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                  Sonr Specific                                 ||
// ! ||--------------------------------------------------------------------------------||

// Address returns the address of the account based on the coin type
func (wa *walletAccount) Address() string {
	return wa.CoinType().FormatAddress(wa.PubKey())
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

// Type returns the type of the account
func (wa *walletAccount) Type() string {
	return fmt.Sprintf("%s/ecdsa-secp256k1", wa.CoinType().Name())
}

// VerificationMethod returns the verification method of the account
func (wa *walletAccount) VerificationMethod(controller string) *types.VerificationMethod {
	return &types.VerificationMethod{
		Id:                  wa.DID(),
		Type:                crypto.Secp256k1KeyType.FormatString(),
		Controller:          controller,
		PublicKeyMultibase:  wa.PubKey().Multibase(),
		BlockchainAccountId: wa.Address(),
	}
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                            Ethereum specific methods                           ||
// ! ||--------------------------------------------------------------------------------||

// Nonce returns the nonce of the account
func (wa *walletAccount) Nonce() uint64 {
	return wa.n
}

// IncrementNonce increments the nonce of the account
func (wa *walletAccount) IncrementNonce() {
	wa.n++
}

//
// ! ||--------------------------------------------------------------------------------||
// ! ||                              Cosmos specific methods                           ||
// ! ||--------------------------------------------------------------------------------||
//

// GetAuthInfo creates an AuthInfo instance for this account with the specified gas amount.
func (wa *walletAccount) GetAuthInfo(gas sdk.Coins) (*txtypes.AuthInfo, error) {
	// Build signerInfo parameters
	anyPubKey, err := codectypes.NewAnyWithValue(wa.PubKey())
	if err != nil {
		return nil, err
	}

	// Create AuthInfo
	authInfo := txtypes.AuthInfo{
		SignerInfos: []*txtypes.SignerInfo{
			{
				PublicKey: anyPubKey,
				ModeInfo: &txtypes.ModeInfo{
					Sum: &txtypes.ModeInfo_Single_{
						Single: &txtypes.ModeInfo_Single{
							Mode: 1,
						},
					},
				},
				Sequence: 0,
			},
		},
		Fee: &txtypes.Fee{
			Amount:   gas,
			GasLimit: uint64(300000),
			Payer:    wa.Address(),
		},
	}
	return &authInfo, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                             Multi-Party Computation                            ||
// ! ||--------------------------------------------------------------------------------||

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

// ! ||--------------------------------------------------------------------------------||
// ! ||                                Filesystem & I/O                                ||
// ! ||--------------------------------------------------------------------------------||

// Name returns the name of the account
func (wa *walletAccount) Name() string {
	return filepath.Base(wa.p)
}

// Path returns the path of the account
func (wa *walletAccount) Path() string {
	return wa.p
}

// Rename renames the account
func (wa *walletAccount) Rename(name string) error {
	parentDir := filepath.Dir(wa.p)
	newPath := filepath.Join(parentDir, name)

	// Rename the directory to the new name
	return os.Rename(wa.p, newPath)
}

// Lock locks the account
func (wa *walletAccount) Lock(c *crypto.WebauthnCredential) error {
	ks, err := wa.ListKeyshares()
	if err != nil {
		return err
	}

	// Encrypt all keyshares for user
	for _, k := range ks {
		if err := k.Encrypt(c); err != nil {
			return err
		}
	}
	return nil
}

// Unlock unlocks the account
func (wa *walletAccount) Unlock(c *crypto.WebauthnCredential) error {
	ks, err := wa.ListKeyshares()
	if err != nil {
		return err
	}

	// Decrypt all keyshares for user
	for _, k := range ks {
		if err := k.Decrypt(c); err != nil {
			return err
		}
	}
	return nil
}

//
// ! ||--------------------------------------------------------------------------------||
// ! ||                                  Helper functions                              ||
// ! ||--------------------------------------------------------------------------------||
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
