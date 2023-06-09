package keeper

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/sonrhq/core/internal/crypto"
	servicetypes "github.com/sonrhq/core/x/service/types"
	"github.com/sonrhq/core/x/vault/internal/sfs"
	"github.com/sonrhq/core/x/vault/types"
)

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
	claimableWallet.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimableWalletKey))
	appendedValue := k.cdc.MustMarshal(&claimableWallet)
	store.Set(GetClaimableWalletIDBytes(claimableWallet.Id), appendedValue)

	// Update claimableWallet count
	k.SetClaimableWalletCount(ctx, count+1)

	return count
}

// SetClaimableWallet set a specific claimableWallet in the store
func (k Keeper) SetClaimableWallet(ctx sdk.Context, claimableWallet types.ClaimableWallet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimableWalletKey))
	b := k.cdc.MustMarshal(&claimableWallet)
	store.Set(GetClaimableWalletIDBytes(claimableWallet.Id), b)
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

// AssignIdentity verifies that a credential is valid and assigns an Unclaimed Wallet to the credential's owner. This creates the initial
// DID document for the user, containing authentication and capability delegation relationships.
func (k Keeper) AssignVault(ctx sdk.Context, ucwId uint64, cred *servicetypes.WebauthnCredential) ([]types.Account, error) {
	// Get the keyshares for the claimable wallet
	accs := make([]types.Account, 0)
	ucw, found := k.GetClaimableWallet(ctx, ucwId)
	if !found {
		k.Logger(ctx).Error("unclaimed wallet not found", "id", ucwId)
		return nil, fmt.Errorf("unclaimed wallet with ID %d not found", ucwId)
	}
	acc, err := sfs.ClaimAccount(ucw.Keyshares, crypto.SONRCoinType, cred)
	if err != nil {
		k.Logger(ctx).Error("failed to resolve account from keyshares", "error", err)
		return nil, fmt.Errorf("error resolving account from keyshares: %w", err)
	}
	k.RemoveClaimableWallet(ctx, ucwId)
	accs = append(accs, acc)
	// Create btc, eth default accounts
	btcAcc, err := acc.DeriveAccount(crypto.BTCCoinType, 0, "BTC#1")
	if err != nil {
		k.Logger(ctx).Error("(Gateway/service) - error deriving BTC account", err)
		return nil, err
	}
	accs = append(accs, btcAcc)
	ethAcc, err := acc.DeriveAccount(crypto.ETHCoinType, 0, "ETH#1")
	if err != nil {
		k.Logger(ctx).Error("(Gateway/service) - error deriving ETH account", err)
		return nil, err
	}
	accs = append(accs, ethAcc)
	return accs, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                            Helper Utility Functions                            ||
// ! ||--------------------------------------------------------------------------------||
// ChallengeLength - Length of bytes to generate for a challenge.¡¡
const ChallengeLength = 32

// createChallenge creates a new challenge that should be signed and returned by the authenticator. The spec recommends
// using at least 16 bytes with 100 bits of entropy. We use 32 bytes.
func createChallenge() (challenge protocol.URLEncodedBase64, err error) {
	challenge = make([]byte, ChallengeLength)

	if _, err = rand.Read(challenge); err != nil {
		return nil, err
	}

	return challenge, nil
}
