// Code generated from Pkl module `sonr.motr.ORM`. DO NOT EDIT.
package models

type Grant struct {
	Id uint `pkl:"id" json:"id,omitempty" query:"id"`

	Subject string `pkl:"subject" json:"subject,omitempty"`

	Controller string `pkl:"controller" json:"controller,omitempty"`

	Origin string `pkl:"origin" json:"origin,omitempty"`

	Token string `pkl:"token" json:"token,omitempty"`

	Scopes []string `pkl:"scopes" json:"scopes,omitempty"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty"`

	UpdatedAt *string `pkl:"updatedAt" json:"updatedAt,omitempty"`
}
