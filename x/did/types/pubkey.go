package types

import (
	didv1 "github.com/onsonr/sonr/api/did/v1"
)

type PubKeyI interface {
	GetRole() string
	GetKeyType() string
	GetRawKey() *didv1.RawKey
	GetJwk() *didv1.JSONWebKey
}
