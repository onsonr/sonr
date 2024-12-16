// Code generated from Pkl module `sonr.orm.Models`. DO NOT EDIT.
package models

type Chain struct {
	Id string `pkl:"id" json:"id,omitempty" query:"id"`

	Name string `pkl:"name" json:"name,omitempty"`

	NetworkId string `pkl:"networkId" json:"networkId,omitempty"`

	ChainCode uint `pkl:"chainCode" json:"chainCode,omitempty"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty"`
}
