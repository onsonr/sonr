// Code generated from Pkl module `orm`. DO NOT EDIT.
package orm

type Session struct {
	Id string `pkl:"id" json:"id,omitempty" query:"id"`

	Subject string `pkl:"subject" json:"subject,omitempty"`

	Controller string `pkl:"controller" json:"controller,omitempty"`

	Origin string `pkl:"origin" json:"origin,omitempty"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty"`

	UpdatedAt *string `pkl:"updatedAt" json:"updatedAt,omitempty"`
}
