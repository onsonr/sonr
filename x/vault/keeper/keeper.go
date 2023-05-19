package keeper

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	servicetypes "github.com/sonrhq/core/x/service/types"
	"github.com/sonrhq/core/x/vault/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
		vaultI *vaultImpl
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
	err := setupVault(k)
	if err != nil {
		k.Logger(sdk.Context{}).Error("Error setting up vault", "error", err)
	}
	return k
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}


// The function inserts an account and its associated key shares into a vault.
func (k Keeper) InsertAccount(acc types.Account) error {
	ksAccListVal := strings.Join(acc.ListKeyShares(), ",")
	_, err := k.vaultI.KsTable.Put(k.vaultI.ctx, accountPrefix(acc.Did()), []byte(ksAccListVal))
	if err != nil {
		return err
	}
	acc.MapKeyShare(func(ks types.KeyShare) types.KeyShare {
		err := k.InsertKeyshare(ks)
		if err != nil {
			return nil
		}
		return ks
	})
	return nil
}

// The function inserts a keyshare into a table and returns an error if there is one.
func (k Keeper) InsertKeyshare(ks types.KeyShare) error {
	_, err := k.vaultI.KsTable.Put(k.vaultI.ctx, keysharePrefix(ks.Did()), ks.Bytes())
	if err != nil {
		return err
	}
	return nil
}

// The function retrieves an account from a key store table using the account's DID and returns it as a
// model.
func (k Keeper) GetAccount(accDid string) (types.Account, error) {
	ksr, err := types.ParseAccountDID(accDid)
	if err != nil {
		return nil, err
	}

	vBiz, err := k.vaultI.KsTable.Get(k.vaultI.ctx, accountPrefix(accDid))
	if err != nil {
		return nil, err
	}

	ksAccListVal := strings.Split(string(vBiz), ",")
	var ksList []types.KeyShare
	for _, ksDid := range ksAccListVal {
		ks, err := k.GetKeyshare(ksDid)
		if err != nil {
			return nil, err
		}
		ksList = append(ksList, ks)
	}
	acc := types.NewAccount(ksList, ksr.CoinType)
	return acc, nil
}

// The function retrieves a keyshare from a vault based on a given key DID.
func (k Keeper) GetKeyshare(keyDid string) (types.KeyShare, error) {

	ksr, err := types.ParseKeyShareDID(keyDid)
	if err != nil {
		return nil, err
	}
	vBiz, err := k.vaultI.KsTable.Get(k.vaultI.ctx, keysharePrefix(keyDid))
	if err != nil {
		return nil, err
	}
	ks, err := types.NewKeyshare(keyDid, vBiz, ksr.CoinType)
	if err != nil {
		return nil, err
	}
	return ks, nil
}

// DeleteAccount deletes an account from the vault.
func (k Keeper) DeleteAccount(accDid string) error {
	// Delete the keyshares
	vBiz, err := k.vaultI.KsTable.Get(k.vaultI.ctx, accountPrefix(accDid))
	if err != nil {
		return err
	}

	ksAccListVal := strings.Split(string(vBiz), ",")
	for _, ksDid := range ksAccListVal {
		_, err = k.vaultI.KsTable.Delete(k.vaultI.ctx, keysharePrefix(ksDid))
		if err != nil {
			return err
		}
	}

	// Delete the account
	_, err = k.vaultI.KsTable.Delete(k.vaultI.ctx, accountPrefix(accDid))
	if err != nil {
		return err
	}
	return nil
}

// FetchCredential retrieves a credential from the vault.
func (k Keeper) FetchCredential(keyDid string) (servicetypes.Credential, error) {
	// Delete the keyshares
	vBiz, err := k.vaultI.KsTable.Get(k.vaultI.ctx, webauthnPrefix(keyDid))
	if err != nil {
		return nil, err
	}

	cred := servicetypes.WebauthnCredential{}
	err = json.Unmarshal(vBiz, &cred)
	if err != nil {
		return nil, err
	}
	return servicetypes.LoadCredential(&cred)
}

// StoreCredential stores a credential in the vault.
func (k Keeper) StoreCredential(cred servicetypes.Credential) error {
	bz, err := cred.Serialize()
	if err != nil {
		return err
	}
	_, err = k.vaultI.KsTable.Put(k.vaultI.ctx, webauthnPrefix(cred.Did()), bz)
	if err != nil {
		return err
	}
	return nil
}


// ReadInbox reads the inbox for the account
func (k Keeper) ReadInbox(accDid string) ([]*types.InboxMessage, error) {
	inbox, err := k.vaultI.LoadInbox(accDid)
	if err != nil {
		return nil, err
	}
	return inbox.Messages, nil
}

// WriteInbox writes a message to the inbox for the account
func (k Keeper) WriteInbox(toDid string, msg *types.InboxMessage) error {
	// Get the inbox
	inbox, err := k.vaultI.LoadInbox(toDid)
	if err != nil {
		return err
	}
	// Add the message to the inbox
	inboxMap, err := inbox.AddMessageToMap(msg)
	if err != nil {
		return err
	}
	// Update the inbox
	_, err = k.vaultI.InTable.Put(k.vaultI.ctx, inboxMap)
	if err != nil {
		return err
	}
	return nil
}
