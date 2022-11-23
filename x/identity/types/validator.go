package types

import (
	"errors"
	"fmt"
	"net/url"
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
	Validate(document *DidDocument) error
}

// MultiValidator is a validator that executes zero or more validators. It returns the first validation error it encounters.
type MultiValidator struct {
	Validators []Validator
}

func (m MultiValidator) Validate(document *DidDocument) error {
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

func (w W3CSpecValidator) Validate(document *DidDocument) error {
	return MultiValidator{[]Validator{
		baseValidator{},
		verificationMethodValidator{},
		verificationMethodRelationshipValidator{
			getter: func(document *DidDocument) VerificationRelationships {
				return *document.Authentication
			},
			err: ErrInvalidAuthentication,
		},
		verificationMethodRelationshipValidator{
			getter: func(document *DidDocument) VerificationRelationships {
				return *document.AssertionMethod
			},
			err: ErrInvalidAssertionMethod,
		},
		verificationMethodRelationshipValidator{
			getter: func(document *DidDocument) VerificationRelationships {
				return *document.KeyAgreement
			},
			err: ErrInvalidKeyAgreement,
		},
		verificationMethodRelationshipValidator{
			getter: func(document *DidDocument) VerificationRelationships {
				return *document.CapabilityInvocation
			},
			err: ErrInvalidCapabilityInvocation,
		},
		verificationMethodRelationshipValidator{
			getter: func(document *DidDocument) VerificationRelationships {
				return *document.CapabilityDelegation
			},
			err: ErrInvalidCapabilityDelegation,
		},
		serviceValidator{},
	}}.Validate(document)
}

// baseValidator validates simple top-level DID document properties (@context, ID, controller)
type baseValidator struct{}

func (w baseValidator) Validate(document *DidDocument) error {
	// Verify `@context`
	if !containsContext(document, DIDContextV1) {
		return makeValidationError(ErrInvalidContext)
	}
	// Verify `id`
	if u, err := url.Parse(document.ID); document.ID == "" || (err == nil && u.Scheme != "" && u.Host != "") {
		return makeValidationError(ErrInvalidID)
	}
	// Verify `controller`
	for _, controller := range document.Controller {
		if controller == "" {
			return makeValidationError(ErrInvalidController)
		}
	}
	return nil
}

type verificationMethodValidator struct{}

func (v verificationMethodValidator) Validate(document *DidDocument) error {
	for _, vm := range document.VerificationMethod.GetData() {
		if !validateVM(vm) {
			return makeValidationError(ErrInvalidVerificationMethod)
		}
	}
	return nil
}

type verificationMethodRelationshipValidator struct {
	getter func(document *DidDocument) VerificationRelationships
	err    error
}

func (v verificationMethodRelationshipValidator) Validate(document *DidDocument) error {
	for _, vm := range v.getter(document).Data {
		if !validateVM(vm.VerificationMethod) {
			return makeValidationError(v.err)
		}
	}
	return nil
}

func validateVM(vm *VerificationMethod) bool {
	if vm.ID == "" || vm.Type == KeyType_KeyType_UNSPECIFIED {
		return false
	}
	if len(strings.TrimSpace(string(vm.Type))) == 0 {
		return false
	}
	if vm.Controller == "" {
		return false
	}
	return true
}

type serviceValidator struct{}

func (s serviceValidator) Validate(document *DidDocument) error {
	for _, service := range document.Service.GetData() {
		if len(strings.TrimSpace(service.ID)) == 0 {
			return makeValidationError(ErrInvalidService)
		}
		if len(strings.TrimSpace(service.Type)) == 0 {
			return makeValidationError(ErrInvalidService)
		}
		if len(strings.TrimSpace(service.ServiceEndpoint)) == 0{
			return makeValidationError(ErrInvalidService)
		}
	}
	return nil
}

func containsContext(document *DidDocument, ctx string) bool {
	for _, curr := range document.Context {
		if curr == ctx {
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
func (d *DidDocument) MatchesID(doc *DidDocument) bool {
	return d.ID == doc.ID
}

// EqualsVerificationMethod compares doc Verification Method and given docs verification method then returns true if they are equal.
func (d *DidDocument) EqualsVerificationMethod(doc *DidDocument) bool {
	if len(d.VerificationMethod.GetData()) != len(doc.VerificationMethod.GetData()) {
		return false
	}
	for _, v := range d.VerificationMethod.GetData() {
		if doc.VerificationMethod.FindByID(v.ID) == nil {
			return false
		}
	}
	return true
}

// EqualsAuthentication compares doc Authentication and given docs authentication then returns true if they are equal.
func (d *DidDocument) EqualsAuthentication(doc *DidDocument) bool {
	if len(d.Authentication.Data) != len(doc.Authentication.Data) {
		return false
	}
	for _, v := range d.Authentication.Data {
		if doc.Authentication.FindByID(v.Reference) == nil {
			return false
		}
	}
	return true
}

// EqualsAssertionMethod compares doc AssertionMethod and given docs assertion method then returns true if they are equal.
func (d *DidDocument) EqualsAssertionMethod(doc *DidDocument) bool {
	if len(d.AssertionMethod.Data) != len(doc.AssertionMethod.Data) {
		return false
	}
	for _, v := range d.AssertionMethod.Data {
		if doc.AssertionMethod.FindByID(v.Reference) == nil {
			return false
		}
	}
	return true
}

// EqualsKeyAgreement compares doc KeyAgreement and given docs key agreement then returns true if they are equal.
func (d *DidDocument) EqualsKeyAgreement(doc *DidDocument) bool {
	if len(d.KeyAgreement.Data) != len(doc.KeyAgreement.Data) {
		return false
	}
	for _, v := range d.KeyAgreement.Data {
		if doc.KeyAgreement.FindByID(v.Reference) == nil {
			return false
		}
	}
	return true
}

// DocumentImpl compares doc CapabilityInvocation and given docs capability invocation then returns true if they are equal.
func (d *DidDocument) EqualsCapabilityInvocation(doc *DidDocument) bool {
	if len(d.CapabilityInvocation.GetData()) != len(doc.CapabilityInvocation.Data) {
		return false
	}
	for _, v := range d.CapabilityInvocation.GetData() {
		if doc.CapabilityInvocation.FindByID(v.Reference) == nil {
			return false
		}
	}
	return true
}

// EqualsCapabilityDelegation compares doc CapabilityDelegation and given docs capability delegation then returns true if they are equal.
func (d *DidDocument) EqualsCapabilityDelegation(doc *DidDocument) bool {
	if len(d.CapabilityDelegation.Data) != len(doc.CapabilityDelegation.Data) {
		return false
	}
	for _, v := range d.CapabilityDelegation.Data {
		if doc.CapabilityDelegation.FindByID(v.Reference) == nil {
			return false
		}
	}
	return true
}

// EqualsService compares doc Service and given docs service then returns true if they are equal.
func (d *DidDocument) EqualsService(doc *DidDocument) bool {
	if len(d.Service.Data) != len(doc.Service.Data) {
		return false
	}
	for _, v := range d.Service.Data {
		if doc.Service.FindByID(v.ID) == nil {
			return false
		}
	}
	return true
}

// EqualsAlsoKnownAs compares doc AlsoKnownAs and given docs also known as then returns true if they are equal.
func (d *DidDocument) EqualsAlsoKnownAs(doc *DidDocument) bool {
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
func (d *DidDocument) copyDocument(doc *DidDocument) error {
	if !d.MatchesID(doc) && doc.ID != "" {
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
