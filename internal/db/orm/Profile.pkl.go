// Code generated from Pkl module `orm`. DO NOT EDIT.
package orm

type Profile struct {
	Id string `pkl:"id" gorm:"primaryKey,autoIncrement" json:"id,omitempty" query:"id"`

	Subject string `pkl:"subject" json:"subject,omitempty" param:"subject"`

	Controller string `pkl:"controller" json:"controller,omitempty" param:"controller"`

	OriginUri *string `pkl:"originUri" json:"originUri,omitempty" param:"originUri"`

	PublicMetadata *string `pkl:"publicMetadata" json:"publicMetadata,omitempty" param:"publicMetadata"`

	PrivateMetadata *string `pkl:"privateMetadata" json:"privateMetadata,omitempty" param:"privateMetadata"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty" param:"createdAt"`

	UpdatedAt *string `pkl:"updatedAt" json:"updatedAt,omitempty" param:"updatedAt"`
}
