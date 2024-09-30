// Code generated from Pkl module `format`. DO NOT EDIT.
package format

type Macaroon struct {
	Location string `pkl:"location" json:"location,omitempty"`

	Originator string `pkl:"originator" json:"originator,omitempty"`

	Identifier string `pkl:"identifier" json:"identifier,omitempty"`

	FirstParty []string `pkl:"first_party" json:"first_party,omitempty"`

	ThirdParty []string `pkl:"third_party" json:"third_party,omitempty"`

	Expiration int `pkl:"expiration" json:"expiration,omitempty"`
}
