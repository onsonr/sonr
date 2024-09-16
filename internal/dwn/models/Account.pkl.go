// Code generated from Pkl module `models`. DO NOT EDIT.
package models

type Account struct {
	Id uint `pkl:"id" gorm:"primaryKey,autoIncrement" json:"id,omitempty" query:"id"`

	Name string `pkl:"name" json:"name,omitempty" param:"name"`

	Address string `pkl:"address" json:"address,omitempty" param:"address"`

	PublicKey string `pkl:"publicKey" json:"publicKey,omitempty" param:"publicKey"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty" param:"createdAt"`
}
