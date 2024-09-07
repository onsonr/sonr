// Code generated from Pkl module `vault`. DO NOT EDIT.
package vault

import "github.com/apple/pkl-go/pkl"

type Property interface {
	Model

	GetId() uint

	GetProfileId() string

	GetKey() string

	GetAccumulator() *pkl.Object

	GetPropertyKey() *pkl.Object
}

var _ Property = (*PropertyImpl)(nil)

type PropertyImpl struct {
	Table string `pkl:"table"`

	Id uint `pkl:"id" gorm:"primaryKey" json:"id,omitempty"`

	ProfileId string `pkl:"profileId" json:"profileId,omitempty"`

	Key string `pkl:"key" json:"key,omitempty"`

	Accumulator *pkl.Object `pkl:"accumulator" json:"accumulator,omitempty"`

	PropertyKey *pkl.Object `pkl:"propertyKey" json:"propertyKey,omitempty"`
}

func (rcv *PropertyImpl) GetTable() string {
	return rcv.Table
}

func (rcv *PropertyImpl) GetId() uint {
	return rcv.Id
}

func (rcv *PropertyImpl) GetProfileId() string {
	return rcv.ProfileId
}

func (rcv *PropertyImpl) GetKey() string {
	return rcv.Key
}

func (rcv *PropertyImpl) GetAccumulator() *pkl.Object {
	return rcv.Accumulator
}

func (rcv *PropertyImpl) GetPropertyKey() *pkl.Object {
	return rcv.PropertyKey
}
