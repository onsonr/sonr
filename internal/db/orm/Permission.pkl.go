// Code generated from Pkl module `orm`. DO NOT EDIT.
package orm

type Permission struct {
	Id uint `pkl:"id" gorm:"primaryKey,autoIncrement" json:"id,omitempty" query:"id"`

	ServiceId string `pkl:"serviceId" json:"serviceId,omitempty" param:"serviceId"`

	Grants string `pkl:"grants" json:"grants,omitempty" param:"grants"`

	Scopes string `pkl:"scopes" json:"scopes,omitempty" param:"scopes"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty" param:"createdAt"`

	UpdatedAt *string `pkl:"updatedAt" json:"updatedAt,omitempty" param:"updatedAt"`
}
