// Code generated from Pkl module `vault`. DO NOT EDIT.
package vault

import "github.com/apple/pkl-go/pkl"

type PublicKey interface {
	Model

	GetId() uint

	GetRole() int

	GetAlgorithm() int

	GetEncoding() int

	GetRaw() *pkl.Object

	GetHex() string

	GetMultibase() string

	GetJwk() *pkl.Object

	GetCreatedAt() *string
}

var _ PublicKey = (*PublicKeyImpl)(nil)

type PublicKeyImpl struct {
	Table string `pkl:"table"`

	Id uint `pkl:"id" gorm:"primaryKey" json:"id,omitempty"`

	Role int `pkl:"role" json:"role,omitempty"`

	Algorithm int `pkl:"algorithm" json:"algorithm,omitempty"`

	Encoding int `pkl:"encoding" json:"encoding,omitempty"`

	Raw *pkl.Object `pkl:"raw" json:"raw,omitempty"`

	Hex string `pkl:"hex" json:"hex,omitempty"`

	Multibase string `pkl:"multibase" json:"multibase,omitempty"`

	Jwk *pkl.Object `pkl:"jwk" json:"jwk,omitempty"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty"`
}

func (rcv *PublicKeyImpl) GetTable() string {
	return rcv.Table
}

func (rcv *PublicKeyImpl) GetId() uint {
	return rcv.Id
}

func (rcv *PublicKeyImpl) GetRole() int {
	return rcv.Role
}

func (rcv *PublicKeyImpl) GetAlgorithm() int {
	return rcv.Algorithm
}

func (rcv *PublicKeyImpl) GetEncoding() int {
	return rcv.Encoding
}

func (rcv *PublicKeyImpl) GetRaw() *pkl.Object {
	return rcv.Raw
}

func (rcv *PublicKeyImpl) GetHex() string {
	return rcv.Hex
}

func (rcv *PublicKeyImpl) GetMultibase() string {
	return rcv.Multibase
}

func (rcv *PublicKeyImpl) GetJwk() *pkl.Object {
	return rcv.Jwk
}

func (rcv *PublicKeyImpl) GetCreatedAt() *string {
	return rcv.CreatedAt
}
