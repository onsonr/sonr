// Code generated from Pkl module `models`. DO NOT EDIT.
package models

type Keyshare struct {
	Id uint `pkl:"id" gorm:"primaryKey,autoIncrement" json:"id,omitempty" query:"id"`

	Data string `pkl:"data" json:"data,omitempty" param:"data"`

	Role int `pkl:"role" json:"role,omitempty" param:"role"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty" param:"createdAt"`
}
