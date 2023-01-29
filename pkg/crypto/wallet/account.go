package wallet

import (
	"time"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/pkg/common/rosetta"
	v1 "github.com/sonrhq/core/x/identity/types/vault/v1"
	"github.com/ucan-wg/go-ucan"
)

// AccountConfig is a type alias for v1.AccountConfig in x/identity/types/vault/v1.
type AccountConfig = v1.AccountConfig

// Account is the interface that implements the cmp Wallet Share which operates within
// a closed p2p network in order to provide a "private key service".  This package
// provides interfaces for common Wallet operations, and implementations for popular
// networks such as Bitcoin, Ethereum, and Cosmos.

// `Account` is an interface that defines the methods that a wallet account must implement.
// @property AccountConfig - The account configuration
// @property Bip32Derive - This is a method that derives a new account from a BIP32 path.
// @property GetAssertionMethod - returns the verification method for the account.
// @property {bool} IsPrimary - returns true if the account is the primary account
// @property ListConfigs - This is a list of all the configurations that are needed to sign a
// transaction.
// @property Sign - This is the function that signs a transaction.
// @property Verify - Verifies a signature
type Account interface {
	// Bip32Derive derives a new account from a BIP32 path
	Bip32Derive(name string, coinType common.CoinType) (Account, error)

	// CoinType returns the coin type of the account
	CoinType() common.CoinType

	// Config returns the account configuration
	Config() *AccountConfig

	// DID returns the DID of the account
	DID() string

	// Info returns the account information
	Info() map[string]string

	// Marshal returns the local config protobuf bytes
	Marshal() ([]byte, error)

	// NewOriginToken creates a new UCAN token
	NewOriginToken(audienceDID string, att ucan.Attenuations, fct []ucan.Fact, notBefore, expires time.Time) (string, error)

	// NewAttenuatedToken creates a new UCAN token from the parent token
	NewAttenuatedToken(parent *ucan.Token, audienceDID string, att ucan.Attenuations, fct []ucan.Fact, notBefore, expires time.Time) (string, error)

	// PubKey returns secp256k1 public key
	PubKey() common.SNRPubKey

	// Signs a message
	Sign(bz []byte) ([]byte, error)

	// Type returns the type of the account
	Type() string

	// Unmarshal deserializes the local config protobuf bytes
	Unmarshal(bz []byte) error

	// Verifies a signature
	Verify(bz []byte, sig []byte) (bool, error)
}

// BTCAccount is an account that can be used to sign Bitcoin transactions. It
// also implements the rosetta.Client interface.
type BTCAccount interface {
	Account
	rosetta.Client

	// Address returns the address of the account.
	Address() string

	// SignTx hashes the transaction for keccak256 (ETH Hashing Function) and signs it with the MPC Protocol
	SignTx(bz []byte) ([]byte, error)
}

// CosmosAccount is an account that can be used to sign Cosmos transactions. It
// is a wrapper around the Account interface that also implements the rosetta.Client.
type CosmosAccount interface {
	Account
	rosetta.Client

	// Address returns the address of the account.
	Address() string

	// GetSignerData returns the signer data for the account
	GetSignerData() authsigning.SignerData

	// Equals returns true if the account is equal to the other account
	Equals(other cryptotypes.LedgerPrivKey) bool

	// Signs a transaction
	SignTxAux(msgs ...sdk.Msg) (txtypes.AuxSignerData, error)
}

// ETHAccount is an account that can be used to sign Ethereum transactions. It
// also implements the rosetta.Client interface.
type ETHAccount interface {
	Account
	rosetta.Client

	// Address returns the address of the account.
	Address() string

	// SignTx hashes the transaction for keccak256 (ETH Hashing Function) and signs it with the MPC Protocol
	SignTx(bz []byte) ([]byte, error)
}
