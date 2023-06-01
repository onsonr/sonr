package sfs

import (
	"errors"

	"berty.tech/go-orbit-db/iface"
	"github.com/sonrhq/core/x/vault/types"
)

// CreateInbox sets up a default inbox for the account
func CreateInbox(accDid string) error {
	inbox, err := types.CreateDefaultInboxMap(accDid)
	if err != nil {
		return err
	}
	_, err = mailTable.Put(ctx, inbox)
	if err != nil {
		return err
	}
	return nil
}

// HasInbox checks if the account has an inbox
func HasInbox(accDid string) (bool, error) {
	inboxRaw, err := mailTable.Get(ctx, accDid, &iface.DocumentStoreGetOptions{})
	if err != nil {
		return false, err
	}
	if len(inboxRaw) == 0 {
		return false, nil
	}
	return true, nil
}

// LoadInbox loads the inbox for the account
func LoadInbox( accDid string) (*types.Inbox, error) {
	// Check if the inbox exists
	hasInbox, err := HasInbox( accDid)
	if err != nil {
		return nil, err
	}
	if !hasInbox {
		err := CreateInbox(accDid)
		if err != nil {
			return nil, err
		}
	}

	// Load the inbox
	inboxRaw, err := mailTable.Get(ctx, accDid, &iface.DocumentStoreGetOptions{})
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
