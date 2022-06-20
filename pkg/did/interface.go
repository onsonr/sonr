package did

import "github.com/sonr-io/sonr/pkg/did/ssi"

type Document interface {
	// AddCapabilityDelegation adds a VerificationMethod as CapabilityDelegation
	// If the controller is not set, it will be set to the document's ID
	AddCapabilityDelegation(v *VerificationMethod)
	// AddCapabilityDelegation adds a VerificationMethod as CapabilityDelegation
	// If the controller is not set, it will be set to the document's ID
	AddAuthenticationMethod(v *VerificationMethod)
	AddAssertionMethod(v *VerificationMethod)
	// AddCapabilityInvocation adds a VerificationMethod as CapabilityInvocation
	// If the controller is not set, it will be set to the document's ID
	AddCapabilityInvocation(v *VerificationMethod)
	// AddService adds a Service
	AddService(s Service)

	CopyFromBytes(b []byte) error

	// IsController returns whether the given DID is a controller of the DID document.
	IsController(controller DID) bool
	ControllersAsString() []string
	ControllerCount() int

	GetID() DID

	GetAlsoKnownAs() []string

	MarshalJSON() ([]byte, error)
	UnmarshalJSON(b []byte) error

	// AddAlias adds a string alias to the document for a .snr domain name into the AlsoKnownAs field
	// in the document.
	AddAlias(alias string)

	AddController(id DID)

	// ResolveEndpointURL finds the endpoint with the given type and unmarshalls it as single URL.
	// It returns the endpoint ID and URL, or an error if anything went wrong;
	// - holder document can't be resolved,
	// - service with given type doesn't exist,
	// - multiple services match,
	// - serviceEndpoint isn't a string.
	ResolveEndpointURL(serviceType string) (endpointID ssi.URI, endpointURL string, err error)

	// EncryptJWE(id DID, buf []byte) (string, error)
	// DecryptJWE(id DID, serial string) ([]byte, error)
	GetController(id DID) (DID, error)

	// FindAssertionMethod finds the first AssertionMethod with the given DID
	FindAssertionMethod(id DID) *VerificationMethod

	// FindAuthenticationMethod finds the first AuthenticationMethod with the given DID
	FindAuthenticationMethod(id DID) *VerificationMethod

	// FindCapabilityDelegation finds the first CapabilityDelegation with the given DID
	FindCapabilityDelegation(id DID) *VerificationMethod

	// FindCapabilityInvocation finds the first CapabilityInvocation with the given DID
	FindCapabilityInvocation(id DID) *VerificationMethod

	// GetAssertionMethods returns all AssertionMethods
	GetAssertionMethods() VerificationRelationships

	// GetAuthenticationMethods returns all AuthenticationMethods
	GetAuthenticationMethods() VerificationRelationships

	// GetCapabilityDelegations returns all CapabilityDelegations
	GetCapabilityDelegations() VerificationRelationships

	// GetCapabilityInvocations returns all CapabilityInvocations
	GetCapabilityInvocations() VerificationRelationships

	// GetServices gets all services on the DID Document
	GetServices() Services
}
