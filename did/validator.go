package did

import (
	"errors"
	"fmt"
	"strings"
)

// ErrDIDDocumentInvalid indicates DID Document validation failed
var ErrDIDDocumentInvalid = validationError{}

// ErrInvalidContext indicates the DID Document's `@context` is invalid
var ErrInvalidContext = errors.New("invalid context")

// ErrInvalidID indicates the DID Document's `id` is invalid
var ErrInvalidID = errors.New("invalid ID")

// ErrInvalidController indicates the DID Document's `controller` is invalid
var ErrInvalidController = errors.New("invalid controller")

// ErrInvalidVerificationMethod indicates the verificationMethod is invalid (e.g. invalid `id` or `type`)
var ErrInvalidVerificationMethod = errors.New("invalid verificationMethod")

// ErrInvalidAuthentication indicates the authentication is invalid (e.g. invalid `id` or `type`)
var ErrInvalidAuthentication = errors.New("invalid authentication")

// ErrInvalidAssertionMethod indicates the assertion method is invalid (e.g. invalid `id` or `type`)
var ErrInvalidAssertionMethod = errors.New("invalid assertionMethod")

// ErrInvalidKeyAgreement indicates the keyAgreement is invalid (e.g. invalid `id` or `type`)
var ErrInvalidKeyAgreement = errors.New("invalid keyAgreement")

// ErrInvalidCapabilityInvocation indicates the capabilityInvocation is invalid (e.g. invalid `id` or `type`)
var ErrInvalidCapabilityInvocation = errors.New("invalid capabilityInvocation")

// ErrInvalidCapabilityDelegation indicates the capabilityDelegation is invalid (e.g. invalid `id` or `type`)
var ErrInvalidCapabilityDelegation = errors.New("invalid capabilityDelegation")

// ErrInvalidService indicates the service is invalid (e.g. invalid `id` or `type`)
var ErrInvalidService = errors.New("invalid service")

// Validator defines functions for validating a DID document.
type Validator interface {
	// Validate validates a DID document. It returns the first validation error is finds wrapped in ErrDIDDocumentInvalid.
	Validate(document Document) error
}

// MultiValidator is a validator that executes zero or more validators. It returns the first validation error it encounters.
type MultiValidator struct {
	Validators []Validator
}

func (m MultiValidator) Validate(document Document) error {
	for _, validator := range m.Validators {
		if err := validator.Validate(document); err != nil {
			return err
		}
	}
	return nil
}

// W3CSpecValidator validates a DID document according to the W3C DID Core Data Model specification (https://www.w3.org/TR/did-core/).
type W3CSpecValidator struct {
}

func (w W3CSpecValidator) Validate(document Document) error {
	return MultiValidator{[]Validator{
		baseValidator{},
		verificationMethodValidator{},
		verificationMethodRelationshipValidator{
			getter: func(document Document) VerificationRelationships {
				return document.Authentication
			},
			err: ErrInvalidAuthentication,
		},
		verificationMethodRelationshipValidator{
			getter: func(document Document) VerificationRelationships {
				return document.AssertionMethod
			},
			err: ErrInvalidAssertionMethod,
		},
		verificationMethodRelationshipValidator{
			getter: func(document Document) VerificationRelationships {
				return document.KeyAgreement
			},
			err: ErrInvalidKeyAgreement,
		},
		verificationMethodRelationshipValidator{
			getter: func(document Document) VerificationRelationships {
				return document.CapabilityInvocation
			},
			err: ErrInvalidCapabilityInvocation,
		},
		verificationMethodRelationshipValidator{
			getter: func(document Document) VerificationRelationships {
				return document.CapabilityDelegation
			},
			err: ErrInvalidCapabilityDelegation,
		},
		serviceValidator{},
	}}.Validate(document)
}

// baseValidator validates simple top-level DID document properties (@context, ID, controller)
type baseValidator struct{}

func (w baseValidator) Validate(document Document) error {
	// Verify `@context`
	if !containsContext(document, DIDContextV1) {
		return makeValidationError(ErrInvalidContext)
	}
	// Verify `id`
	if document.ID.Empty() || document.ID.IsURL() {
		return makeValidationError(ErrInvalidID)
	}
	// Verify `controller`
	for _, controller := range document.Controller {
		if controller.Empty() || controller.IsURL() {
			return makeValidationError(ErrInvalidController)
		}
	}
	return nil
}

type verificationMethodValidator struct{}

func (v verificationMethodValidator) Validate(document Document) error {
	for _, vm := range document.VerificationMethod {
		if !validateVM(vm) {
			return makeValidationError(ErrInvalidVerificationMethod)
		}
	}
	return nil
}

type verificationMethodRelationshipValidator struct {
	getter func(document Document) VerificationRelationships
	err    error
}

func (v verificationMethodRelationshipValidator) Validate(document Document) error {
	for _, vm := range v.getter(document) {
		if !validateVM(vm.VerificationMethod) {
			return makeValidationError(v.err)
		}
	}
	return nil
}

func validateVM(vm *VerificationMethod) bool {
	if vm.ID.Empty() {
		return false
	}
	if len(strings.TrimSpace(string(vm.Type))) == 0 {
		return false
	}
	if vm.Controller.Empty() {
		return false
	}
	return true
}

type serviceValidator struct{}

func (s serviceValidator) Validate(document Document) error {
	for _, service := range document.Service {
		if len(strings.TrimSpace(service.ID.String())) == 0 {
			return makeValidationError(ErrInvalidService)
		}
		if len(strings.TrimSpace(service.Type)) == 0 {
			return makeValidationError(ErrInvalidService)
		}
		if service.ServiceEndpoint == nil {
			return makeValidationError(ErrInvalidService)
		}
		switch x := service.ServiceEndpoint.(type) {
		case string:
		case map[string]interface{}:
		case []interface{}:
			_ = x // We need to 'do' something with 'x'
			break
		default:
			return makeValidationError(ErrInvalidService)
		}
	}
	return nil
}

func containsContext(document Document, ctx string) bool {
	for _, curr := range document.Context {
		if curr.String() == ctx {
			return true
		}
	}
	return false
}

func makeValidationError(validationErr error) error {
	return validationError{cause: validationErr}
}

type validationError struct {
	cause error
}

func (v validationError) Unwrap() error {
	return v.cause
}

func (v validationError) Is(err error) bool {
	_, is := err.(validationError)
	return is
}

func (v validationError) Error() string {
	return fmt.Sprintf("DID Document validation failed: %v", v.cause)
}
