// Code generated from Pkl module `txns`. DO NOT EDIT.
package txns

type SWT struct {
	Origin string `pkl:"origin"`

	Location string `pkl:"location"`

	Identifier string `pkl:"identifier"`

	Scopes []string `pkl:"scopes"`

	Properties map[string]string `pkl:"properties"`

	ExpiryBlock int `pkl:"expiryBlock"`
}
