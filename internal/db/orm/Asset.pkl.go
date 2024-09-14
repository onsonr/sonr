// Code generated from Pkl module `orm`. DO NOT EDIT.
package orm

type Asset struct {
	Id uint `pkl:"id" gorm:"primaryKey,autoIncrement" json:"id,omitempty" query:"id"`

	Name string `pkl:"name" json:"name,omitempty" param:"name"`

	Symbol string `pkl:"symbol" json:"symbol,omitempty" param:"symbol"`

	Decimals int `pkl:"decimals" json:"decimals,omitempty" param:"decimals"`

	ChainId *int `pkl:"chainId" json:"chainId,omitempty" param:"chainId"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty" param:"createdAt"`
}
