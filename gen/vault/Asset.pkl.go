// Code generated from Pkl module `vault`. DO NOT EDIT.
package vault

type Asset interface {
	Model

	GetId() uint

	GetName() string

	GetSymbol() string

	GetDecimals() int

	GetChainId() *int

	GetCreatedAt() *string
}

var _ Asset = (*AssetImpl)(nil)

type AssetImpl struct {
	Table string `pkl:"table"`

	Id uint `pkl:"id" gorm:"primaryKey" json:"id,omitempty"`

	Name string `pkl:"name" json:"name,omitempty"`

	Symbol string `pkl:"symbol" json:"symbol,omitempty"`

	Decimals int `pkl:"decimals" json:"decimals,omitempty"`

	ChainId *int `pkl:"chainId" json:"chainId,omitempty"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty"`
}

func (rcv *AssetImpl) GetTable() string {
	return rcv.Table
}

func (rcv *AssetImpl) GetId() uint {
	return rcv.Id
}

func (rcv *AssetImpl) GetName() string {
	return rcv.Name
}

func (rcv *AssetImpl) GetSymbol() string {
	return rcv.Symbol
}

func (rcv *AssetImpl) GetDecimals() int {
	return rcv.Decimals
}

func (rcv *AssetImpl) GetChainId() *int {
	return rcv.ChainId
}

func (rcv *AssetImpl) GetCreatedAt() *string {
	return rcv.CreatedAt
}
