// Code generated from Pkl module `models`. DO NOT EDIT.
package models

type Account struct {
	Id string `pkl:"id" json:"id,omitempty" query:"id"`

	Name string `pkl:"name" json:"name,omitempty"`

	Address any `pkl:"address" json:"address,omitempty"`

	PublicKey string `pkl:"publicKey" json:"publicKey,omitempty"`

	ChainCode uint `pkl:"chainCode" json:"chainCode,omitempty"`

	Index int `pkl:"index" json:"index,omitempty"`

	Controller string `pkl:"controller" json:"controller,omitempty"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty"`
}
