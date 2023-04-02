package vault

import (
	"errors"

	"berty.tech/go-orbit-db/iface"
	"github.com/sonrhq/core/internal/models"
)

// CreateInbox sets up a default inbox for the account
func (v *vaultImpl) CreateInbox(accDid string) error {
	inbox, err := models.CreateDefaultInboxMap(accDid)
	if err != nil {
		return err
	}
	_, err = v.InTable.Put(v.ctx, inbox)
	if err != nil {
		return err
	}
	return nil
}

// LoadInbox loads the inbox for the account
func (v *vaultImpl) LoadInbox(accDid string) (*models.Inbox, error) {
	inboxRaw, err := v.InTable.Get(v.ctx, accDid, &iface.DocumentStoreGetOptions{
		CaseInsensitive: true,
		PartialMatches:  false,
	})
	if err != nil {
		err = v.CreateInbox(accDid)
		if err != nil {
			return nil, err
		}
		return v.LoadInbox(accDid)
	}
	if len(inboxRaw) == 0 {
		err = v.CreateInbox(accDid)
		if err != nil {
			return nil, err
		}
		return v.LoadInbox(accDid)
	}
	inboxMap, ok := inboxRaw[0].(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid inbox")
	}
	inbox, err := models.NewInboxFromMap(inboxMap)
	if err != nil {
		return nil, err
	}
	return inbox, nil
}

// ReadInbox reads the inbox for the account
func ReadInbox(accDid string) ([]*models.InboxMessage, error) {
	inbox, err := v.LoadInbox(accDid)
	if err != nil {
		return nil, err
	}
	return inbox.Messages, nil
}

// WriteInbox writes a message to the inbox for the account
func WriteInbox(toDid string, msg *models.InboxMessage) error {
	// Get the inbox
	inbox, err := v.LoadInbox(toDid)
	if err != nil {
		return err
	}
	// Add the message to the inbox
	inboxMap, err := inbox.AddMessageToMap(msg)
	if err != nil {
		return err
	}
	// Update the inbox
	_, err = v.InTable.Put(v.ctx, inboxMap)
	if err != nil {
		return err
	}
	return nil
}
