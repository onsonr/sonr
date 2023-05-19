package types

import (
	"encoding/json"
)

type InboxMessage struct {
	// The message id
	Id string `json:"id,omitempty"`

	// The message type
	Type string `json:"type,omitempty"`

	// The message content
	Content string `json:"content,omitempty"`

	// The message sender
	Sender string `json:"sender,omitempty"`

	// The message receiver
	Receiver string `json:"receiver,omitempty"`

	// CoinType Name is for the blockchain network name
	CoinType string `json:"coinType,omitempty"`
}

type Inbox struct {
	// The inbox id
	Id string `json:"_id,omitempty"`

	// The inbox messages
	Messages []*InboxMessage `json:"messages,omitempty"`
}

func CreateDefaultInboxMap(id string) (map[string]interface{}, error) {
	inbox := &Inbox{
		Id:       id,
		Messages: []*InboxMessage{},
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

func (inbox *Inbox) AddMessageToMap(msg *InboxMessage) (map[string]interface{}, error) {
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

func (im *InboxMessage) InboxMessageToMap() (map[string]interface{}, error) {
	data, err := json.Marshal(im)
	if err != nil {
		return nil, err
	}
	var inboxMessageMap map[string]interface{}
	err = json.Unmarshal(data, &inboxMessageMap)
	if err != nil {
		return nil, err
	}
	return inboxMessageMap, nil
}
