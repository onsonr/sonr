// Code generated from Pkl module `models`. DO NOT EDIT.
package orm

type Profile struct {
	Id string `pkl:"id" json:"id,omitempty" query:"id"`

	Subject string `pkl:"subject" json:"subject,omitempty"`

	Controller string `pkl:"controller" json:"controller,omitempty"`

	OriginUri *string `pkl:"originUri" json:"originUri,omitempty"`

	PublicMetadata *string `pkl:"publicMetadata" json:"publicMetadata,omitempty"`

	PrivateMetadata *string `pkl:"privateMetadata" json:"privateMetadata,omitempty"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty"`

	UpdatedAt *string `pkl:"updatedAt" json:"updatedAt,omitempty"`
}
