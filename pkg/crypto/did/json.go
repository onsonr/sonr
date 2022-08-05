package did

import (
	"github.com/sonr-io/sonr/pkg/crypto/did/internal/marshal"
)

const contextKey = "@context"
const controllerKey = "controller"
const authenticationKey = "authentication"
const assertionMethodKey = "assertionMethod"
const keyAgreementKey = "keyAgreement"
const capabilityInvocationKey = "capabilityInvocation"
const capabilityDelegationKey = "capabilityDelegation"
const verificationMethodKey = "verificationMethod"
const serviceEndpointKey = "serviceEndpoint"

var pluralContext = marshal.Plural(contextKey)
