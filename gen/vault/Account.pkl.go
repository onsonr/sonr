// Code generated from Pkl module `vault`. DO NOT EDIT.
package vault

import "github.com/apple/pkl-go/pkl"

type Account interface {
	Model

	GetId() uint

	GetName() string

	GetAddress() string

	GetPublicKey() *pkl.Object

	GetCreatedAt() *string
}

var _ Account = (*AccountImpl)(nil)

type AccountImpl struct {
	Table string `pkl:"table"`

	Id uint `pkl:"id" gorm:"primaryKey" json:"id,omitempty"`

	Name string `pkl:"name" json:"name,omitempty"`

	Address string `pkl:"address" json:"address,omitempty"`

	PublicKey *pkl.Object `pkl:"publicKey" json:"publicKey,omitempty"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty"`
}

func (rcv *AccountImpl) GetTable() string {
	return rcv.Table
}

func (rcv *AccountImpl) GetId() uint {
	return rcv.Id
}

func (rcv *AccountImpl) GetName() string {
	return rcv.Name
}

func (rcv *AccountImpl) GetAddress() string {
	return rcv.Address
}

func (rcv *AccountImpl) GetPublicKey() *pkl.Object {
	return rcv.PublicKey
}

func (rcv *AccountImpl) GetCreatedAt() *string {
	return rcv.CreatedAt
}
