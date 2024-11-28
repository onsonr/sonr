// Code generated from Pkl module `sonr.motr.ORM`. DO NOT EDIT.
package models

import (
	"github.com/onsonr/sonr/x/dwn/types/models/keyalgorithm"
	"github.com/onsonr/sonr/x/dwn/types/models/keycurve"
	"github.com/onsonr/sonr/x/dwn/types/models/keyencoding"
	"github.com/onsonr/sonr/x/dwn/types/models/keyrole"
	"github.com/onsonr/sonr/x/dwn/types/models/keytype"
)

type DID struct {
	Id string `pkl:"id" json:"id,omitempty" query:"id"`

	Role keyrole.KeyRole `pkl:"role"`

	Algorithm keyalgorithm.KeyAlgorithm `pkl:"algorithm"`

	Encoding keyencoding.KeyEncoding `pkl:"encoding"`

	Curve keycurve.KeyCurve `pkl:"curve"`

	KeyType keytype.KeyType `pkl:"key_type"`

	Raw string `pkl:"raw"`

	Jwk *JWK `pkl:"jwk"`
}
