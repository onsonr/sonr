// Code generated from Pkl module `models`. DO NOT EDIT.
package orm

import (
	"github.com/onsonr/sonr/x/did/types/orm/keyalgorithm"
	"github.com/onsonr/sonr/x/did/types/orm/keycurve"
	"github.com/onsonr/sonr/x/did/types/orm/keyencoding"
	"github.com/onsonr/sonr/x/did/types/orm/keyrole"
	"github.com/onsonr/sonr/x/did/types/orm/keytype"
)

type PublicKey struct {
	Role keyrole.KeyRole `pkl:"role" json:"role,omitempty" query:"role"`

	Algorithm keyalgorithm.KeyAlgorithm `pkl:"algorithm"`

	Encoding keyencoding.KeyEncoding `pkl:"encoding"`

	Curve keycurve.KeyCurve `pkl:"curve"`

	KeyType keytype.KeyType `pkl:"key_type"`

	Raw string `pkl:"raw"`

	Jwk *JWK `pkl:"jwk"`
}
