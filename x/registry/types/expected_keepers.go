package types

import (
	crypto "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	// Return a new account with the next account number and the specified address. Does not save the new account to the store.
	NewAccountWithAddress(sdk.Context, sdk.AccAddress) types.AccountI

	// Return a new account with the next account number. Does not save the new account to the store.
	NewAccount(sdk.Context, types.AccountI) types.AccountI

	// Check if an account exists in the store.
	HasAccount(sdk.Context, sdk.AccAddress) bool

	// Retrieve an account from the store.
	GetAccount(sdk.Context, sdk.AccAddress) types.AccountI

	// Set an account in the store.
	SetAccount(sdk.Context, types.AccountI)

	// Remove an account from the store.
	RemoveAccount(sdk.Context, types.AccountI)

	// Iterate over all accounts, calling the provided function. Stop iteration when it returns true.
	IterateAccounts(sdk.Context, func(types.AccountI) bool)

	// Fetch the public key of an account at a specified address
	GetPubKey(sdk.Context, sdk.AccAddress) (crypto.PubKey, error)

	// Fetch the sequence of an account at a specified address.
	GetSequence(sdk.Context, sdk.AccAddress) (uint64, error)

	// Fetch the next account number, and increment the internal counter.
	GetNextAccountNumber(sdk.Context) uint64
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	// Methods imported from bank should be defined here
}
