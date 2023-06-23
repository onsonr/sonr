package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/go-webauthn/webauthn/protocol"
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
	return acc.GetAccountInfo(), nil
}

// GetClaimableWalletCount get the total number of claimableWallet
func (k Keeper) GetClaimableWalletCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ClaimableWalletCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetClaimableWalletCount set the total number of claimableWallet
func (k Keeper) SetClaimableWalletCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ClaimableWalletCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendClaimableWallet appends a claimableWallet in the store with a new id and update the count
func (k Keeper) AppendClaimableWallet(
	ctx sdk.Context,
	claimableWallet types.ClaimableWallet,
) uint64 {
	// Create the claimableWallet
	count := k.GetClaimableWalletCount(ctx)

	// Set the ID of the appended value
	claimableWallet.Index = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimableWalletKey))
	appendedValue := k.cdc.MustMarshal(&claimableWallet)
	store.Set(GetClaimableWalletIDBytes(claimableWallet.Index), appendedValue)

	// Update claimableWallet count
	k.SetClaimableWalletCount(ctx, count+1)

	return count
}

// SetClaimableWallet set a specific claimableWallet in the store
func (k Keeper) SetClaimableWallet(ctx sdk.Context, claimableWallet types.ClaimableWallet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimableWalletKey))
	b := k.cdc.MustMarshal(&claimableWallet)
	store.Set(GetClaimableWalletIDBytes(claimableWallet.Index), b)
}

// GetClaimableWallet returns a claimableWallet from its id
func (k Keeper) GetClaimableWallet(ctx sdk.Context, id uint64) (val types.ClaimableWallet, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimableWalletKey))
	b := store.Get(GetClaimableWalletIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveClaimableWallet removes a claimableWallet from the store
func (k Keeper) RemoveClaimableWallet(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimableWalletKey))
	store.Delete(GetClaimableWalletIDBytes(id))
	k.Logger(ctx).Info("(x/vault) Removed claimable wallet", "id", id)
}

// GetAllClaimableWallet returns all claimableWallet
func (k Keeper) GetAllClaimableWallet(ctx sdk.Context) (list []types.ClaimableWallet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimableWalletKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ClaimableWallet
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetClaimableWalletIDBytes returns the byte representation of the ID
func GetClaimableWalletIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetClaimableWalletIDFromBytes returns ID in uint64 format from a byte array
func GetClaimableWalletIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

// NextUnclaimedWallet returns the next unclaimed wallet and its challenge. If no unclaimed wallets exist, an error is
func (k Keeper) NextUnclaimedWallet(ctx sdk.Context) (*types.ClaimableWallet, protocol.URLEncodedBase64, error) {
	// Make sure more than zero unclaimed wallets exist
	if k.GetClaimableWalletCount(ctx) == 0 {
		return nil, nil, fmt.Errorf("no unclaimed wallets exist")
	}

	// Get all unclaimed wallets
	var ucws []types.ClaimableWallet
	store := ctx.KVStore(k.storeKey)
	claimableWalletStore := prefix.NewStore(store, types.KeyPrefix(types.ClaimableWalletKey))
	_, err := query.Paginate(claimableWalletStore, nil, func(key []byte, value []byte) error {
		var claimableWallet types.ClaimableWallet
		if err := k.cdc.Unmarshal(value, &claimableWallet); err != nil {
			k.Logger(ctx).Error("failed to unmarshal claimable wallet", "error", err)
			return err
		}

		ucws = append(ucws, claimableWallet)
		return nil
	})
	if err != nil {
		k.Logger(ctx).Error("failed to get unclaimed wallets", "error", err)
		return nil, nil, err
	}

	// Get the next unclaimed wallet
	ucw := ucws[0]
	chal, err := createChallenge()
	if err != nil {
		k.Logger(ctx).Error("failed to create challenge", "error", err)
		return nil, nil, fmt.Errorf("error creating challenge: %w", err)
	}
	return &ucw, chal, nil
}
