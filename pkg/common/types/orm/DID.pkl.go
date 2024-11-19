// Code generated from Pkl module `common.types.ORM`. DO NOT EDIT.
package orm

import (
	"github.com/onsonr/sonr/pkg/common/types/orm/keyalgorithm"
	"github.com/onsonr/sonr/pkg/common/types/orm/keycurve"
	"github.com/onsonr/sonr/pkg/common/types/orm/keyencoding"
	"github.com/onsonr/sonr/pkg/common/types/orm/keyrole"
	"github.com/onsonr/sonr/pkg/common/types/orm/keytype"
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
