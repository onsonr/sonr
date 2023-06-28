package keeper

import (
	"errors"
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/sonrhq/core/internal/sfs"
	"github.com/sonrhq/core/x/vault/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// The function retrieves an account from a key store table using the account's DID and returns it as a
// model.
func (k Keeper) GetAccount(accDid string) (types.KeyShareCollection, error) {
return nil, errors.New("not implemented")
}



// ReadInbox reads the inbox for the account
func (k Keeper) ReadInbox(accDid string) ([]*types.WalletMail, error) {
	return sfs.ReadInbox(accDid)
}

// WriteInbox writes a message to the inbox for the account
func (k Keeper) WriteInbox(toDid string, msg *types.WalletMail) error {
	return sfs.WriteInbox(toDid, msg)
}

// GetAccountInfo returns the account info for the given account DID
func (k Keeper) GetAccountInfo(accDid string) (*types.AccountInfo, error) {
return nil, errors.New("not implemented")
}
