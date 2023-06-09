package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/sonrhq/core/x/vault/internal/sfs"
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
	k := &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
	return k
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// The function inserts an account and its associated key shares into a vault.
func (k Keeper) InsertAccount(acc types.Account) error {
	sfs.InsertAccount(acc)
	return nil
}

// The function inserts a keyshare into a table and returns an error if there is one.
func (k Keeper) InsertKeyshare(ks types.KeyShare) error {
	sfs.InsertKeyshare(ks)
	return nil
}

// The function retrieves an account from a key store table using the account's DID and returns it as a
// model.
func (k Keeper) GetAccount(accDid string) (types.Account, error) {
	return sfs.GetAccount(accDid)
}

// The function retrieves a keyshare from a vault based on a given key DID.
func (k Keeper) GetKeyshare(keyDid string) (types.KeyShare, error) {
return sfs.GetKeyshare(keyDid)
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
	acc, err := sfs.GetAccount(accDid)
	if err != nil {
		return nil, err
	}
	return acc.ToProto(), nil
}
