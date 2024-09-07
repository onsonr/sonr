// Code generated from Pkl module `vault`. DO NOT EDIT.
package vault

import "github.com/apple/pkl-go/pkl"

type Keyshare interface {
	Model

	GetId() uint

	GetMetadata() string

	GetPayloads() string

	GetProtocol() string

	GetPublicKey() *pkl.Object

	GetRole() int

	GetVersion() int

	GetCreatedAt() *string
}

var _ Keyshare = (*KeyshareImpl)(nil)

type KeyshareImpl struct {
	Table string `pkl:"table"`

	Id uint `pkl:"id" gorm:"primaryKey" json:"id,omitempty"`

	Metadata string `pkl:"metadata" json:"metadata,omitempty"`

	Payloads string `pkl:"payloads" json:"payloads,omitempty"`

	Protocol string `pkl:"protocol" json:"protocol,omitempty"`

	PublicKey *pkl.Object `pkl:"publicKey" json:"publicKey,omitempty"`

	Role int `pkl:"role" json:"role,omitempty"`

	Version int `pkl:"version" json:"version,omitempty"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty"`
}

func (rcv *KeyshareImpl) GetTable() string {
	return rcv.Table
}

func (rcv *KeyshareImpl) GetId() uint {
	return rcv.Id
}

func (rcv *KeyshareImpl) GetMetadata() string {
	return rcv.Metadata
}

func (rcv *KeyshareImpl) GetPayloads() string {
	return rcv.Payloads
}

func (rcv *KeyshareImpl) GetProtocol() string {
	return rcv.Protocol
}

func (rcv *KeyshareImpl) GetPublicKey() *pkl.Object {
	return rcv.PublicKey
}

func (rcv *KeyshareImpl) GetRole() int {
	return rcv.Role
}

func (rcv *KeyshareImpl) GetVersion() int {
	return rcv.Version
}

func (rcv *KeyshareImpl) GetCreatedAt() *string {
	return rcv.CreatedAt
}
