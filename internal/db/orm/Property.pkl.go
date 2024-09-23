// Code generated from Pkl module `orm`. DO NOT EDIT.
package orm

type Property struct {
	Id uint `pkl:"id" gorm:"primaryKey,autoIncrement" json:"id,omitempty" query:"id"`

	ProfileId string `pkl:"profileId" json:"profileId,omitempty" param:"profileId"`

	Key string `pkl:"key" json:"key,omitempty" param:"key"`

	Accumulator string `pkl:"accumulator" json:"accumulator,omitempty" param:"accumulator"`

	PropertyKey string `pkl:"propertyKey" json:"propertyKey,omitempty" param:"propertyKey"`
}
