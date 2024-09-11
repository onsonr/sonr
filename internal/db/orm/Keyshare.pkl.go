// Code generated from Pkl module `orm`. DO NOT EDIT.
package orm

type Keyshare struct {
	Id uint `pkl:"id" gorm:"primaryKey,autoIncrement" json:"id,omitempty" query:"id"`

	Metadata string `pkl:"metadata" json:"metadata,omitempty" param:"metadata"`

	Payloads string `pkl:"payloads" json:"payloads,omitempty" param:"payloads"`

	Protocol string `pkl:"protocol" json:"protocol,omitempty" param:"protocol"`

	PublicKey string `pkl:"publicKey" json:"publicKey,omitempty" param:"publicKey"`

	Role int `pkl:"role" json:"role,omitempty" param:"role"`

	Version int `pkl:"version" json:"version,omitempty" param:"version"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty" param:"createdAt"`
}
