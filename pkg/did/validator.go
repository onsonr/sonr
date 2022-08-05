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
	Validate(document DocumentImpl) error
}

// MultiValidator is a validator that executes zero or more validators. It returns the first validation error it encounters.
type MultiValidator struct {
	Validators []Validator
}

func (m MultiValidator) Validate(document DocumentImpl) error {
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

func (w W3CSpecValidator) Validate(document DocumentImpl) error {
	return MultiValidator{[]Validator{
		baseValidator{},
		verificationMethodValidator{},
		verificationMethodRelationshipValidator{
			getter: func(document DocumentImpl) VerificationRelationships {
				return document.Authentication
			},
			err: ErrInvalidAuthentication,
		},
		verificationMethodRelationshipValidator{
			getter: func(document DocumentImpl) VerificationRelationships {
				return document.AssertionMethod
			},
			err: ErrInvalidAssertionMethod,
		},
		verificationMethodRelationshipValidator{
			getter: func(document DocumentImpl) VerificationRelationships {
				return document.KeyAgreement
			},
			err: ErrInvalidKeyAgreement,
		},
		verificationMethodRelationshipValidator{
			getter: func(document DocumentImpl) VerificationRelationships {
				return document.CapabilityInvocation
			},
			err: ErrInvalidCapabilityInvocation,
		},
		verificationMethodRelationshipValidator{
			getter: func(document DocumentImpl) VerificationRelationships {
				return document.CapabilityDelegation
			},
			err: ErrInvalidCapabilityDelegation,
		},
		serviceValidator{},
	}}.Validate(document)
}

// baseValidator validates simple top-level DID document properties (@context, ID, controller)
type baseValidator struct{}

func (w baseValidator) Validate(document DocumentImpl) error {
	// Verify `@context`
	if !containsContext(document, DIDContextV1) {
		return makeValidationError(ErrInvalidContext)
	}
	// Verify `id`
	if document.ID.Empty() {
		return makeValidationError(ErrInvalidID)
	}
	// Verify `controller`
	for _, controller := range document.Controller {
		if controller.Empty() {
			return makeValidationError(ErrInvalidController)
		}
	}
	return nil
}

type verificationMethodValidator struct{}

func (v verificationMethodValidator) Validate(document DocumentImpl) error {
	for _, vm := range document.VerificationMethod {
		if !validateVM(vm) {
			return makeValidationError(ErrInvalidVerificationMethod)
		}
	}
	return nil
}

type verificationMethodRelationshipValidator struct {
	getter func(document DocumentImpl) VerificationRelationships
	err    error
}

func (v verificationMethodRelationshipValidator) Validate(document DocumentImpl) error {
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

func (s serviceValidator) Validate(document DocumentImpl) error {
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
	}
	return nil
}

func containsContext(document DocumentImpl, ctx string) bool {
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

// MatchesID returns true if the two DIDs are equal.
func (d DocumentImpl) MatchesID(doc *DocumentImpl) bool {
	return d.ID.String() == doc.ID.String()
}

// EqualsVerificationMethod compares doc Verification Method and given docs verification method then returns true if they are equal.
func (d DocumentImpl) EqualsVerificationMethod(doc *DocumentImpl) bool {
	if len(d.VerificationMethod) != len(doc.VerificationMethod) {
		return false
	}
	for _, v := range d.VerificationMethod {
		if doc.VerificationMethod.FindByID(v.ID) == nil {
			return false
		}
	}
	return true
}

// EqualsAuthentication compares doc Authentication and given docs authentication then returns true if they are equal.
func (d DocumentImpl) EqualsAuthentication(doc *DocumentImpl) bool {
	if len(d.Authentication) != len(doc.Authentication) {
		return false
	}
	for _, v := range d.Authentication {
		if doc.Authentication.FindByID(v.ID) == nil {
			return false
		}
	}
	return true
}

// EqualsAssertionMethod compares doc AssertionMethod and given docs assertion method then returns true if they are equal.
func (d DocumentImpl) EqualsAssertionMethod(doc *DocumentImpl) bool {
	if len(d.AssertionMethod) != len(doc.AssertionMethod) {
		return false
	}
	for _, v := range d.AssertionMethod {
		if doc.AssertionMethod.FindByID(v.ID) == nil {
			return false
		}
	}
	return true
}

// EqualsKeyAgreement compares doc KeyAgreement and given docs key agreement then returns true if they are equal.
func (d DocumentImpl) EqualsKeyAgreement(doc *DocumentImpl) bool {
	if len(d.KeyAgreement) != len(doc.KeyAgreement) {
		return false
	}
	for _, v := range d.KeyAgreement {
		if doc.KeyAgreement.FindByID(v.ID) == nil {
			return false
		}
	}
	return true
}

// DocumentImpl compares doc CapabilityInvocation and given docs capability invocation then returns true if they are equal.
func (d DocumentImpl) EqualsCapabilityInvocation(doc *DocumentImpl) bool {
	if len(d.CapabilityInvocation) != len(doc.CapabilityInvocation) {
		return false
	}
	for _, v := range d.CapabilityInvocation {
		if doc.CapabilityInvocation.FindByID(v.ID) == nil {
			return false
		}
	}
	return true
}

// EqualsCapabilityDelegation compares doc CapabilityDelegation and given docs capability delegation then returns true if they are equal.
func (d DocumentImpl) EqualsCapabilityDelegation(doc *DocumentImpl) bool {
	if len(d.CapabilityDelegation) != len(doc.CapabilityDelegation) {
		return false
	}
	for _, v := range d.CapabilityDelegation {
		if doc.CapabilityDelegation.FindByID(v.ID) == nil {
			return false
		}
	}
	return true
}

// EqualsService compares doc Service and given docs service then returns true if they are equal.
func (d DocumentImpl) EqualsService(doc *DocumentImpl) bool {
	if len(d.Service) != len(doc.Service) {
		return false
	}
	for _, v := range d.Service {
		if doc.Service.FindByID(v.ID) == nil {
			return false
		}
	}
	return true
}

// EqualsAlsoKnownAs compares doc AlsoKnownAs and given docs also known as then returns true if they are equal.
func (d DocumentImpl) EqualsAlsoKnownAs(doc *DocumentImpl) bool {
	if len(d.AlsoKnownAs) != len(doc.AlsoKnownAs) {
		return false
	}
	for _, v := range d.AlsoKnownAs {
		if !contains(doc.AlsoKnownAs, v) {
			return false
		}
	}
	return true
}

// Equals is a helper function that compares two documents and returns true if they are equal.
func (d *DocumentImpl) copyDocument(doc *DocumentImpl) error {
	if !d.MatchesID(doc) && doc.ID.String() != "" {
		d.ID = doc.ID
	}
	if !d.EqualsVerificationMethod(doc) && doc.VerificationMethod != nil {
		d.VerificationMethod = doc.VerificationMethod
	}
	if !d.EqualsAuthentication(doc) && doc.Authentication != nil {
		d.Authentication = doc.Authentication
	}
	if !d.EqualsAssertionMethod(doc) && doc.AssertionMethod != nil {
		d.AssertionMethod = doc.AssertionMethod
	}
	if !d.EqualsKeyAgreement(doc) && doc.KeyAgreement != nil {
		d.KeyAgreement = doc.KeyAgreement
	}
	if !d.EqualsCapabilityInvocation(doc) && doc.CapabilityInvocation != nil {
		d.CapabilityInvocation = doc.CapabilityInvocation
	}
	if !d.EqualsCapabilityDelegation(doc) && doc.CapabilityDelegation != nil {
		d.CapabilityDelegation = doc.CapabilityDelegation
	}
	if !d.EqualsService(doc) && doc.Service != nil {
		d.Service = doc.Service
	}
	if !d.EqualsAlsoKnownAs(doc) && doc.AlsoKnownAs != nil {
		d.AlsoKnownAs = doc.AlsoKnownAs
	}
	return nil
}
