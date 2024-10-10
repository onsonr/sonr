// Code generated from Pkl module `orm`. DO NOT EDIT.
package orm

type Keyshare struct {
	Id string `pkl:"id" json:"id,omitempty" query:"id"`

	Data string `pkl:"data" json:"data,omitempty"`

	Role int `pkl:"role" json:"role,omitempty"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty"`

	LastRefreshed *string `pkl:"lastRefreshed" json:"lastRefreshed,omitempty"`
}
