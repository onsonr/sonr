// Code generated from Pkl module `common.types.ORM`. DO NOT EDIT.
package models

type Asset struct {
	Id string `pkl:"id" json:"id,omitempty" query:"id"`

	Name string `pkl:"name" json:"name,omitempty"`

	Symbol string `pkl:"symbol" json:"symbol,omitempty"`

	Decimals int `pkl:"decimals" json:"decimals,omitempty"`

	ChainCode uint `pkl:"chainCode" json:"chainCode,omitempty"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty"`
}
