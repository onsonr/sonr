// Code generated from Pkl module `orm`. DO NOT EDIT.
package orm

import (
	"github.com/onsonr/sonr/internal/orm/keyalgorithm"
	"github.com/onsonr/sonr/internal/orm/keycurve"
	"github.com/onsonr/sonr/internal/orm/keyencoding"
	"github.com/onsonr/sonr/internal/orm/keyrole"
	"github.com/onsonr/sonr/internal/orm/keytype"
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
