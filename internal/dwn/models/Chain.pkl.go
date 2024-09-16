// Code generated from Pkl module `models`. DO NOT EDIT.
package models

type Chain struct {
	Id uint `pkl:"id" gorm:"primaryKey,autoIncrement" json:"id,omitempty" query:"id"`

	Name string `pkl:"name" json:"name,omitempty" param:"name"`

	NetworkId string `pkl:"networkId" json:"networkId,omitempty" param:"networkId"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty" param:"createdAt"`
}
