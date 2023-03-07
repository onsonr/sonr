package vc

import (
	"github.com/sonrhq/core/pkg/crypto/types/internal/marshal"
)

const (
	contextKey              = "@context"
	typeKey                 = "type"
	credentialSubjectKey    = "credentialSubject"
	proofKey                = "proof"
	verifiableCredentialKey = "verifiableCredential"
)

var pluralContext = marshal.Plural(contextKey)
