package types

import (
	"encoding/json"
)

type Inbox struct {
	// The inbox id
	Id string `json:"_id,omitempty"`

	// The inbox messages
	Messages []*WalletMail `json:"messages,omitempty"`
}

func CreateDefaultInboxMap(id string) (map[string]interface{}, error) {
	inbox := &Inbox{
		Id:       id,
		Messages: []*WalletMail{},
	}
	return inbox.ToMap()
}

func NewInboxFromMap(inboxMap map[string]interface{}) (*Inbox, error) {
	data, err := json.Marshal(inboxMap)
	if err != nil {
		return nil, err
	}
	var inbox *Inbox
	err = json.Unmarshal(data, &inbox)
	if err != nil {
		return nil, err
	}
	return inbox, nil
}

func (inbox *Inbox) AddMessageToMap(msg *WalletMail) (map[string]interface{}, error) {
	inbox.Messages = append(inbox.Messages, msg)
	return inbox.ToMap()
}

func (inbox *Inbox) ToMap() (map[string]interface{}, error) {
	data, err := json.Marshal(inbox)
	if err != nil {
		return nil, err
	}
	var inboxMap map[string]interface{}
	err = json.Unmarshal(data, &inboxMap)
	if err != nil {
		return nil, err
	}
	return inboxMap, nil
}

func (im *WalletMail) WalletMailToMap() (map[string]interface{}, error) {
	data, err := json.Marshal(im)
	if err != nil {
		return nil, err
	}
	var WalletMailMap map[string]interface{}
	err = json.Unmarshal(data, &WalletMailMap)
	if err != nil {
		return nil, err
	}
	return WalletMailMap, nil
}
