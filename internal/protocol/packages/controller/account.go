package controller

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/getsentry/sentry-go"
	_ "github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/crypto/mpc"
	v1 "github.com/sonrhq/core/types/highway/v1"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// Account is an interface for an account in the wallet
type Account interface {
	// Address returns the address of the account based on the coin type
	Address() string

	// CoinType returns the coin type of the account
	CoinType() crypto.CoinType

	// DID returns the DID of the account
	Did() string

	// Get the controller's DID document
	DidDocument(controller string) *types.DidDocument

	// GetAuthInfo creates an AuthInfo for a transaction
	GetAuthInfo(gas sdk.Coins) (*txtypes.AuthInfo, error)

	// ListKeyshares returns a list of keyshares for the account
	ListKeyshares() ([]KeyShare, error)

	// MapKeyshares performs a function on each keyshare of the account
	MapKeyshares(f func(KeyShare) error) error

	// Name returns the name of the account
	Name() string

	// PartyIDs returns the party IDs of the account
	PartyIDs() []crypto.PartyID

	// Nonce returns the nonce of the account
	Nonce() uint64

	// IncrementNonce increments the nonce of the account
	IncrementNonce()

	// PubKey returns secp256k1 public key
	PubKey() *crypto.PubKey

	// Signs a message
	Sign(bz []byte) ([]byte, error)

	// ToProto returns the proto representation of the account
	ToProto() (*v1.Account)

	// ToStore returns the store representation of the account
	ToStore() (string, []string)

	// Type returns the type of the account
	Type() string

	// Verifies a signature
	Verify(bz []byte, sig []byte) (bool, error)

	// Lock locks the account
	Lock(c *crypto.WebauthnCredential, rootDir string) error

	// Unlock unlocks the account
	Unlock(c *crypto.WebauthnCredential, rootDir string) error
}

type account struct {
	kss []KeyShare
	n   uint64
	p   string
	ct crypto.CoinType
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                     General                                    ||
// ! ||--------------------------------------------------------------------------------||

// NewAccount creates a new account
func NewAccount(kss []KeyShare, ct crypto.CoinType) Account {
	return &account{kss: kss, n: 0, p: "", ct: ct}
}


// PubKey returns secp256k1 public key
func (wa *account) PubKey() *crypto.PubKey {
	tks, err := getFirstDecryptedKeyshare(wa.kss)
	if err != nil {
		return nil
	}
	return tks.PubKey()
}

// Signs a message using the account
func (wa *account) Sign(bz []byte) ([]byte, error) {
	kss, err := wa.ListKeyshares()
	if err != nil {
		return nil, err
	}
	var configs []*cmp.Config
	for _, ks := range kss {
		configs = append(configs, ks.Config())
	}
	return mpc.SignCMP(configs, bz)
}

// ToProto returns the proto representation of the account
func (wa *account) ToProto() (*v1.Account) {
	return &v1.Account{
		Name: wa.Name(),
		Address: wa.Address(),
		CoinType: wa.CoinType().Name(),
		ChainId: "sonr-testnet-0",
		PublicKey: wa.PubKey().Base64(),
		Type: wa.Type(),
	}
}

func (wa *account) ToStore() (string, []string) {
	selfDid := wa.Did()
	ksDids := make([]string, 0)
	for _, ks := range wa.kss {
		ksDids = append(ksDids, ks.Did())
	}
	return selfDid, ksDids
}

// Verifies a signature using first unlocked keyshare
func (wa *account) Verify(bz []byte, sig []byte) (bool, error) {
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
func (a *account) Address() string {
	return a.CoinType().FormatAddress(a.PubKey())
}

// CoinType returns the coin type of the account
func (a *account) CoinType() crypto.CoinType {
	return a.ct
}

// DID returns the DID of the account
func (wa *account) Did() string {
	tks, err := getFirstDecryptedKeyshare(wa.kss)
	if err != nil {
		sentry.CaptureException(err)
		return ""
	}
	return fmt.Sprintf("did:%s:%s", tks.CoinType().DidMethod(), wa.Address())
}

// DidDocument returns the DID document of the account
func (wa *account) DidDocument(controller string) *types.DidDocument {
	doc := types.NewBlockchainIdentity(controller, wa.CoinType(), wa.PubKey())
	return doc
}

// Type returns the type of the account
func (wa *account) Type() string {
	return fmt.Sprintf("%s/ecdsa-secp256k1", wa.CoinType().Name())
}

// VerificationMethod returns the verification method of the account
func (wa *account) VerificationMethod(controller string) *types.VerificationMethod {
	return &types.VerificationMethod{
		Id:                  wa.Did(),
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
func (wa *account) Nonce() uint64 {
	return wa.n
}

// IncrementNonce increments the nonce of the account
func (wa *account) IncrementNonce() {
	wa.n++
}

//
// ! ||--------------------------------------------------------------------------------||
// ! ||                              Cosmos specific methods                           ||
// ! ||--------------------------------------------------------------------------------||
//

// GetAuthInfo creates an AuthInfo instance for this account with the specified gas amount.
func (wa *account) GetAuthInfo(gas sdk.Coins) (*txtypes.AuthInfo, error) {
	// Build signerInfo parameters
	anyPubKey, err := codectypes.NewAnyWithValue(wa.PubKey())
	if err != nil {
		sentry.CaptureException(err)
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
func (wa *account) PartyIDs() []crypto.PartyID {
	var ids []crypto.PartyID
	for _, ks := range wa.kss {
		ids = append(ids, ks.PartyID())
	}
	return ids
}

// ListKeyshares returns a list of keyshares for the account
func (wa *account) ListKeyshares() ([]KeyShare, error) {
	return wa.kss, nil
}

// MapKeyshares performs a mapping function on the keyshares of the account
func (wa *account) MapKeyshares(f func(KeyShare) error) error {
	for _, ks := range wa.kss {
		err := f(ks)
		if err != nil {
			sentry.CaptureException(err)
			return err
		}
	}
	return nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                Filesystem & I/O                                ||
// ! ||--------------------------------------------------------------------------------||

// Name returns the name of the account
func (wa *account) Name() string {
	ks, err := getFirstDecryptedKeyshare(wa.kss)
	if err != nil {
		sentry.CaptureException(err)
		return ""
	}
	kspr, err := ParseKeyShareDid(ks.Did())
	if err != nil {
		sentry.CaptureException(err)
		return ""
	}
	return kspr.AccountName
}

// Lock locks the account
func (wa *account) Lock(c *crypto.WebauthnCredential, rootDir string) error {
	ks, err := wa.ListKeyshares()
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	// Encrypt all keyshares for user
	for _, k := range ks {
		if err := k.Encrypt(c); err != nil {
			sentry.CaptureException(err)
			return err
		}
	}
	return nil
}

// Unlock unlocks the account
func (wa *account) Unlock(c *crypto.WebauthnCredential, rootDir string) error {
	ks, err := wa.ListKeyshares()
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	// Decrypt all keyshares for user
	for _, k := range ks {
		if err := k.Decrypt(c); err != nil {
			sentry.CaptureException(err)
			return err
		}
	}
	return nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                  Helper functions                              ||
// ! ||--------------------------------------------------------------------------------||

// getFirstDecryptedKeyshare returns the first decrypted keyshare
func getFirstDecryptedKeyshare(kss []KeyShare) (KeyShare, error) {
	for _, ks := range kss {
		if !ks.IsEncrypted() {
			return ks, nil
		}
	}
	err := fmt.Errorf("no decrypted keyshares found")
	sentry.CaptureException(err)
	return nil, err
}
