package keys

import "github.com/onsonr/sonr/crypto/core/curves"

type PubKey interface {
	Type() string
	Value() string
}

type mpcPubKey struct {
	PublicPoint curves.Point
}
