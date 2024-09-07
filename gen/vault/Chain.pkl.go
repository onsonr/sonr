// Code generated from Pkl module `vault`. DO NOT EDIT.
package vault

type Chain interface {
	Model

	GetId() uint

	GetName() string

	GetNetworkId() string

	GetCreatedAt() *string
}

var _ Chain = (*ChainImpl)(nil)

type ChainImpl struct {
	Table string `pkl:"table"`

	Id uint `pkl:"id" gorm:"primaryKey" json:"id,omitempty"`

	Name string `pkl:"name" json:"name,omitempty"`

	NetworkId string `pkl:"networkId" json:"networkId,omitempty"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty"`
}

func (rcv *ChainImpl) GetTable() string {
	return rcv.Table
}

func (rcv *ChainImpl) GetId() uint {
	return rcv.Id
}

func (rcv *ChainImpl) GetName() string {
	return rcv.Name
}

func (rcv *ChainImpl) GetNetworkId() string {
	return rcv.NetworkId
}

func (rcv *ChainImpl) GetCreatedAt() *string {
	return rcv.CreatedAt
}
