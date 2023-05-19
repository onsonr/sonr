package keeper

import (
	"context"
	"errors"

	"berty.tech/go-orbit-db/iface"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/x/vault/internal/node"
	"github.com/sonrhq/core/x/vault/types"
)


type vaultImpl struct {
	KsTable node.IPFSKVStore
	InTable node.IPFSDocsStore

	ctx context.Context
}

func setupVault(k *Keeper) error {
	ctx := context.Background()
	snrctx := local.Context()
	kv, err := node.OpenKeyValueStore(ctx, snrctx.GlobalKvKsStore)
	if err != nil {
		return err
	}
	docs, err := node.OpenDocumentStore(ctx, snrctx.GlobalInboxDocsStore, nil)
	if err != nil {
		return err
	}
	vi := &vaultImpl{
		KsTable: kv,
		InTable: docs,
		ctx:     ctx,
	}
	k.vaultI = vi
	return nil
}


// ! ||--------------------------------------------------------------------------------||
// ! ||                         Inbox handler for W2W messages                         ||
// ! ||--------------------------------------------------------------------------------||

// CreateInbox sets up a default inbox for the account
func (v *vaultImpl) CreateInbox(accDid string) error {
	inbox, err := types.CreateDefaultInboxMap(accDid)
	if err != nil {
		return err
	}
	_, err = v.InTable.Put(v.ctx, inbox)
	if err != nil {
		return err
	}
	return nil
}

// HasInbox checks if the account has an inbox
func (v *vaultImpl) HasInbox(accDid string) (bool, error) {
	inboxRaw, err := v.InTable.Get(v.ctx, accDid, &iface.DocumentStoreGetOptions{})
	if err != nil {
		return false, err
	}
	if len(inboxRaw) == 0 {
		return false, nil
	}
	return true, nil
}

// LoadInbox loads the inbox for the account
func (v *vaultImpl) LoadInbox(accDid string) (*types.Inbox, error) {
	// Check if the inbox exists
	hasInbox, err := v.HasInbox(accDid)
	if err != nil {
		return nil, err
	}
	if !hasInbox {
		err := v.CreateInbox(accDid)
		if err != nil {
			return nil, err
		}
	}

	// Load the inbox
	inboxRaw, err := v.InTable.Get(v.ctx, accDid, &iface.DocumentStoreGetOptions{})
	inboxMap, ok := inboxRaw[0].(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid inbox")
	}
	inbox, err := types.NewInboxFromMap(inboxMap)
	if err != nil {
		return nil, err
	}
	return inbox, nil
}


// ! ||--------------------------------------------------------------------------------||
// ! ||                         Helper Methods for Module Setup                        ||
// ! ||--------------------------------------------------------------------------------||

func keysharePrefix(v string) string {
	return "ks/" + v
}

func accountPrefix(v string) string {
	return "acc/" + v
}

func webauthnPrefix(v string) string {
	return "webauthn/" + v
}
