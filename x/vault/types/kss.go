package types

import (
	"fmt"
	"strings"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/pkg/mpc"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
	"lukechampine.com/blake3"
)

const (
	kVaultFragment = "vault"
	kClaimsFragment = "ucw-"
)

// NewKSS creates a new KeyShareCollection
func NewKSS(kss ...*VaultKeyshare) KeyShareCollection {
	return KeyShareCollection(kss)
}

// Address returns the address of the account based on the coin type
func (a KeyShareCollection) Address() string {
	return a.CoinType().FormatAddress(a.PubKey())
}

// CoinType returns the coin type of the account
func (a KeyShareCollection) CoinType() crypto.CoinType {
	return crypto.CoinTypeFromBipPath(a[0].CoinType)
}

// FindKeyshare returns the keyshare that contains the fragment
func (a KeyShareCollection) FindKeyshare(fragment string) *VaultKeyshare {
	for _, ks := range a {
		if strings.Contains(ks.Id, fragment) {
			return ks
		}
	}
	return nil
}

// GenerateSecretKey generates a new secret phrase of 32 bytes
func (a KeyShareCollection) GenerateSecretKey(fragment string) ([]byte, error) {
	sig, err := a.Sign([]byte(fragment))
	if err != nil {
		return nil, err
	}
	hashDerivKey := blake3.Sum256(sig)
	return hashDerivKey[:], nil
}

// Index returns the keyshare at the index
func (a KeyShareCollection) Index(i int) *VaultKeyshare {
	return a[i]
}

// IsValid returns true if the keyshare collection is valid. A valid keyshare collection has at least 2 keyshares
func (a KeyShareCollection) IsValid() bool {
	return len(a) >= 2
}

// PubKey returns secp256k1 public key
func (wa KeyShareCollection) PubKey() *crypto.PubKey {
	return wa[0].PubKey()
}

// PubKey returns secp256k1 public key
func (wa KeyShareCollection) PubKeyType() string {
	return wa[0].PubKey().KeyType
}

// Signs a message using the account
func (wa KeyShareCollection) Sign(bz []byte) ([]byte, error) {
	var configs []*cmp.Config
	for _, ks := range wa {
		configs = append(configs, ks.CMPConfig())
	}
	return mpc.SignCMP(configs, bz)
}

// Signs a cosmos transaction
func (wa KeyShareCollection) SignCosmosTx(msgs ...sdk.Msg) ([]byte, error) {
	if !wa.CoinType().IsCosmos() && !wa.CoinType().IsSonr() {
		return nil, fmt.Errorf("coin type %s not supported for cosmos tx signing", wa.CoinType())
	}
	return SignAnyTransactions(wa, msgs...)
}



// ! ||--------------------------------------------------------------------------------||
// ! ||                                  Sonr Specific                                 ||
// ! ||--------------------------------------------------------------------------------||

// DID returns the DID of the account
func (wa KeyShareCollection) Did() string {
	return fmt.Sprintf("did:%s:%s", wa.CoinType().DidMethod(), wa.Address())
}

// GetAccountInfo returns the proto representation of the account
func (wa KeyShareCollection) GetAccountInfo() *AccountInfo {
	return &AccountInfo{
		Address:   wa.Address(),
		CoinType:  wa.CoinType().String(),
		Did:       wa.Did(),
		PublicKey: wa.PubKey().Base64(),
		Type:      wa.Type(),
	}
}

// Type returns the type of the account
func (wa KeyShareCollection) Type() string {
	return fmt.Sprintf("%s/ecdsa-secp256k1", wa.CoinType().Name())
}

//
// ! ||--------------------------------------------------------------------------------||
// ! ||                              Cosmos specific methods                           ||
// ! ||--------------------------------------------------------------------------------||
//

// GetAuthInfo creates an AuthInfo instance for this account with the specified gas amount.
func (wa KeyShareCollection) GetAuthInfo(gas sdk.Coins) (*txtypes.AuthInfo, error) {
	// Build signerInfo parameters
	seckpPubKey, err := wa.PubKey().Secp256k1()
	if err != nil {
		return nil, err
	}
	anyPubKey, err := codectypes.NewAnyWithValue(seckpPubKey)
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
		},
	}
	return &authInfo, nil
}
