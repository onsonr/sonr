// Code generated from Pkl module `sonr.orm.Models`. DO NOT EDIT.
package models

import (
	"github.com/onsonr/sonr/internal/models/keyalgorithm"
	"github.com/onsonr/sonr/internal/models/keycurve"
	"github.com/onsonr/sonr/internal/models/keyencoding"
	"github.com/onsonr/sonr/internal/models/keyrole"
	"github.com/onsonr/sonr/internal/models/keytype"
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
