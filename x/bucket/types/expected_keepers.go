package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	rt "github.com/sonr-io/sonr/x/registry/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// RegistryKeeper defines the expected interface needed to manage account did documents.
type RegistryKeeper interface {
	GetWhoIs(ctx sdk.Context, id string) (val rt.WhoIs, found bool)
	FindWhoIsByAlias(ctx sdk.Context, alias string) (val rt.WhoIs, found bool)
	GetWhoIsFromOwner(ctx sdk.Context, owner string) (val rt.WhoIs, found bool)
	// Methods imported from bank should be defined here
}
