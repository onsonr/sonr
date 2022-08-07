package did

import (
	"crypto"
	"crypto/ed25519"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/shengdoushi/base58"
	"github.com/sonr-io/sonr/pkg/crypto/jwx"
	"github.com/sonr-io/sonr/pkg/did/internal/marshal"
	"github.com/sonr-io/sonr/pkg/did/ssi"
)

// DocumentImpl represents a DID Document as specified by the DID Core specification (https://www.w3.org/TR/did-core/).
type DocumentImpl struct {
	Context              []ssi.URI                 `json:"@context"`
	ID                   DID                       `json:"id"`
	Controller           []DID                     `json:"controller,omitempty"`
	VerificationMethod   VerificationMethods       `json:"verificationMethod,omitempty"`
	Authentication       VerificationRelationships `json:"authentication,omitempty"`
	AssertionMethod      VerificationRelationships `json:"assertionMethod,omitempty"`
	KeyAgreement         VerificationRelationships `json:"keyAgreement,omitempty"`
	CapabilityInvocation VerificationRelationships `json:"capabilityInvocation,omitempty"`
	CapabilityDelegation VerificationRelationships `json:"capabilityDelegation,omitempty"`
	Service              Services                  `json:"service,omitempty"`
	AlsoKnownAs          []string                  `json:"alsoKnownAs,omitempty"`
}

// BlankDocument creates a Blank Default DID Document
func BlankDocument() Document {
	return &DocumentImpl{
		Context:              make([]ssi.URI, 0),
		VerificationMethod:   make(VerificationMethods, 0),
		Authentication:       make(VerificationRelationships, 0),
		AssertionMethod:      make(VerificationRelationships, 0),
		KeyAgreement:         make(VerificationRelationships, 0),
		CapabilityInvocation: make(VerificationRelationships, 0),
		CapabilityDelegation: make(VerificationRelationships, 0),
		Service:              make([]Service, 0),
		AlsoKnownAs:          make([]string, 0),
	}
}

// NewDocument generates a new DID Document for the provided ID string
func NewDocument(idStr string) (Document, error) {
	fmt.Println(idStr)
	id, err := ParseDID(idStr)
	if err != nil {
		return nil, err
	}

	ctxUri, err := ssi.ParseURI("https://www.w3.org/ns/did/v1")
	if err != nil {
		return nil, err
	}

	return &DocumentImpl{
		ID:                   *id,
		Context:              []ssi.URI{*ctxUri},
		VerificationMethod:   make(VerificationMethods, 0),
		Authentication:       make(VerificationRelationships, 0),
		AssertionMethod:      make(VerificationRelationships, 0),
		KeyAgreement:         make(VerificationRelationships, 0),
		CapabilityInvocation: make(VerificationRelationships, 0),
		CapabilityDelegation: make(VerificationRelationships, 0),
		Service:              make([]Service, 0),
		AlsoKnownAs:          make([]string, 0),
	}, nil
}

func (d *DocumentImpl) ControllerCount() int {
	return len(d.Controller)
}

// FindAuthenticationMethod finds a VerificationMethod by its ID
func (d *DocumentImpl) FindAuthenticationMethod(id DID) *VerificationMethod {
	return d.Authentication.FindByID(id)
}

// FindAssertionMethod finds a VerificationMethod by its ID
func (d *DocumentImpl) FindAssertionMethod(id DID) *VerificationMethod {
	return d.AssertionMethod.FindByID(id)
}

// FindCapabilityDelegation finds a VerificationMethod by its ID
func (d *DocumentImpl) FindCapabilityDelegation(id DID) *VerificationMethod {
	return d.CapabilityDelegation.FindByID(id)
}

// FindCapabilityInvocation finds a VerificationMethod by its ID
func (d *DocumentImpl) FindCapabilityInvocation(id DID) *VerificationMethod {
	return d.CapabilityInvocation.FindByID(id)
}

func (d *DocumentImpl) GetController(did DID) (DID, error) {
	for _, c := range d.Controller {
		if c.Equals(did) {
			return c, nil
		}
	}
	return DID{}, errors.New("did not found")
}

func (d *DocumentImpl) GetContext() []ssi.URI {
	return d.Context
}

// GetAssertionMethods returns the list of assertion methods
func (d *DocumentImpl) GetAssertionMethods() VerificationRelationships {
	return d.AssertionMethod
}

// GetAuthenticationMethods returns the list of authentication methods
func (d *DocumentImpl) GetAuthenticationMethods() VerificationRelationships {
	return d.Authentication
}

// GetCapabilityDelegations returns the list of capability delegations
func (d *DocumentImpl) GetCapabilityDelegations() VerificationRelationships {
	return d.CapabilityDelegation
}

// GetCapabilityInvocations returns the list of capability invocations
func (d *DocumentImpl) GetCapabilityInvocations() VerificationRelationships {
	return d.CapabilityInvocation
}

func (d *DocumentImpl) GetServices() Services {
	return d.Service
}

func (d *DocumentImpl) GetID() DID {
	return d.ID
}

func (d *DocumentImpl) GetAlsoKnownAs() []string {
	return d.AlsoKnownAs
}

// CopyFromBytes unmarshals a JSON document from a byte slice and copies the data into the receiver.
func (d *DocumentImpl) CopyFromBytes(b []byte) error {
	var newDoc DocumentImpl
	err := newDoc.UnmarshalJSON(b)
	if err != nil {
		return err
	}
	return d.copyDocument(&newDoc)
}

// AddController adds a DID as a controller
func (d *DocumentImpl) AddController(id DID) {
	if d.Controller == nil {
		d.Controller = make([]DID, 0)
	}
	d.Controller = append(d.Controller, id)
}

type VerificationMethods []*VerificationMethod

func (d *DocumentImpl) GetVerificationMethods() VerificationMethods {
	return d.VerificationMethod
}

// FindByID find the first VerificationMethod which matches the provided DID.
// Returns nil when not found
func (vms VerificationMethods) FindByID(id DID) *VerificationMethod {
	for _, vm := range vms {
		if vm.ID.Equals(id) {
			return vm
		}
	}
	return nil
}

// Remove removes a VerificationMethod from the slice.
// If a verificationMethod was removed with the given DID, it will be returned
func (vms *VerificationMethods) Remove(id DID) *VerificationMethod {
	var (
		filteredVMS []*VerificationMethod
		foundVM     *VerificationMethod
	)
	for _, vm := range *vms {
		if !vm.ID.Equals(id) {
			filteredVMS = append(filteredVMS, vm)
		} else {
			foundVM = vm
		}
	}
	*vms = filteredVMS
	return foundVM
}

// Add adds a verificationMethod to the verificationMethods if it not already present.
func (vms *VerificationMethods) Add(v *VerificationMethod) {
	for _, ptr := range *vms {
		// check if the pointer is already in the list
		if ptr == v {
			return
		}
		// check if the actual ids match?
		if ptr.ID.Equals(v.ID) {
			return
		}
	}
	*vms = append(*vms, v)
}

type VerificationRelationships []VerificationRelationship

// Count returns the number of VerificationRelationships in the slice
func (vmr VerificationRelationships) Count() int {
	return len(vmr)
}

// FindByID returns the first VerificationRelationship that matches with the id.
// For comparison both the ID of the embedded VerificationMethod and reference is used.
func (vmr VerificationRelationships) FindByID(id DID) *VerificationMethod {
	for _, r := range vmr {
		if r.VerificationMethod != nil {
			if r.VerificationMethod.ID.Equals(id) {
				return r.VerificationMethod
			}
		}
	}
	return nil
}

// Remove removes a VerificationRelationship from the slice.
// If a VerificationRelationship was removed with the given DID, it will be returned
func (vmr *VerificationRelationships) Remove(id DID) *VerificationRelationship {
	var (
		filteredVMRels []VerificationRelationship
		removedRel     *VerificationRelationship
	)
	for _, r := range *vmr {
		if !r.ID.Equals(id) {
			filteredVMRels = append(filteredVMRels, r)
		} else {
			removedRel = &r
		}
	}
	*vmr = filteredVMRels
	return removedRel
}

// Add adds a verificationMethod to a relationship collection.
// When the collection already contains the method it will not be added again.
func (vmr *VerificationRelationships) Add(vm *VerificationMethod) {
	for _, rel := range *vmr {
		if rel.ID.Equals(vm.ID) {
			return
		}
	}
	*vmr = append(*vmr, VerificationRelationship{vm, vm.ID})
}

// AddAuthenticationMethod adds a VerificationMethod as AuthenticationMethod
// If the controller is not set, it will be set to the document's ID
func (d *DocumentImpl) AddAuthenticationMethod(v *VerificationMethod) {
	if v.Controller.Empty() {
		v.Controller = d.ID
	}
	d.VerificationMethod.Add(v)
	d.Authentication.Add(v)
}

// AddAssertionMethod adds a VerificationMethod as AssertionMethod
// If the controller is not set, it will be set to the documents ID
func (d *DocumentImpl) AddAssertionMethod(v *VerificationMethod) {
	if v.Controller.Empty() {
		v.Controller = d.ID
	}
	d.VerificationMethod.Add(v)
	d.AssertionMethod.Add(v)
}

func (d *DocumentImpl) GetKeyAgreements() VerificationRelationships {
	return d.KeyAgreement
}

// AddKeyAgreement adds a VerificationMethod as KeyAgreement
// If the controller is not set, it will be set to the document's ID
func (d *DocumentImpl) AddKeyAgreement(v *VerificationMethod) {
	if v.Controller.Empty() {
		v.Controller = d.ID
	}
	d.VerificationMethod.Add(v)
	d.KeyAgreement.Add(v)
}

// AddCapabilityInvocation adds a VerificationMethod as CapabilityInvocation
// If the controller is not set, it will be set to the document's ID
func (d *DocumentImpl) AddCapabilityInvocation(v *VerificationMethod) {
	if v.Controller.Empty() {
		v.Controller = d.ID
	}
	d.VerificationMethod.Add(v)
	d.CapabilityInvocation.Add(v)
}

// AddCapabilityDelegation adds a VerificationMethod as CapabilityDelegation
// If the controller is not set, it will be set to the document's ID
func (d *DocumentImpl) AddCapabilityDelegation(v *VerificationMethod) {
	if v.Controller.Empty() {
		v.Controller = d.ID
	}
	d.VerificationMethod.Add(v)
	d.CapabilityDelegation.Add(v)
}

func (d *DocumentImpl) AddService(s Service) {
	d.Service = append(d.Service, s)
}

func (d DocumentImpl) MarshalJSON() ([]byte, error) {
	type alias DocumentImpl
	tmp := alias(d)
	if data, err := json.Marshal(tmp); err != nil {
		return nil, err
	} else {
		return marshal.NormalizeDocument(data, marshal.Unplural(contextKey), marshal.Unplural(controllerKey))
	}
}

func (d *DocumentImpl) UnmarshalJSON(b []byte) error {
	type Alias DocumentImpl
	normalizedDoc, err := marshal.NormalizeDocument(b, pluralContext, marshal.Plural(controllerKey))
	if err != nil {
		return err
	}
	doc := Alias{}
	err = json.Unmarshal(normalizedDoc, &doc)
	if err != nil {
		return err
	}
	*d = (DocumentImpl)(doc)

	const errMsg = "unable to resolve all '%s' references: %w"
	if err = resolveVerificationRelationships(d.Authentication, d.VerificationMethod); err != nil {
		return fmt.Errorf(errMsg, authenticationKey, err)
	}
	if err = resolveVerificationRelationships(d.AssertionMethod, d.VerificationMethod); err != nil {
		return fmt.Errorf(errMsg, assertionMethodKey, err)
	}
	if err = resolveVerificationRelationships(d.KeyAgreement, d.VerificationMethod); err != nil {
		return fmt.Errorf(errMsg, keyAgreementKey, err)
	}
	if err = resolveVerificationRelationships(d.CapabilityInvocation, d.VerificationMethod); err != nil {
		return fmt.Errorf(errMsg, capabilityInvocationKey, err)
	}
	if err = resolveVerificationRelationships(d.CapabilityDelegation, d.VerificationMethod); err != nil {
		return fmt.Errorf(errMsg, capabilityDelegationKey, err)
	}
	return nil
}

// IsController returns whether the given DID is a controller of the DID document.
func (d DocumentImpl) IsController(controller DID) bool {
	if controller.Empty() {
		return false
	}
	for _, curr := range d.Controller {
		if curr.Equals(controller) {
			return true
		}
	}

	return false
}

// AddAlias adds a string alias to the document for a .snr domain name into the AlsoKnownAs field
// in the document.
func (d *DocumentImpl) AddAlias(alias string) {
	if d.AlsoKnownAs == nil {
		d.AlsoKnownAs = make([]string, 0)
	}
	d.AlsoKnownAs = append(d.AlsoKnownAs, alias)
}

// ResolveEndpointURL finds the endpoint with the given type and unmarshalls it as single URL.
// It returns the endpoint ID and URL, or an error if anything went wrong;
// - holder document can't be resolved,
// - service with given type doesn't exist,
// - multiple services match,
// - serviceEndpoint isn't a string.
func (d *DocumentImpl) ResolveEndpointURL(serviceType string) (endpointID ssi.URI, endpointURL string, err error) {
	var services []Service
	for _, service := range d.Service {
		if service.Type == serviceType {
			services = append(services, service)
		}
	}
	if len(services) == 0 {
		return ssi.URI{}, "", fmt.Errorf("service not found (did=%s, type=%s)", d.ID, serviceType)
	}
	if len(services) > 1 {
		return ssi.URI{}, "", fmt.Errorf("multiple services found (did=%s, type=%s)", d.ID, serviceType)
	}
	err = services[0].UnmarshalServiceEndpoint(&endpointURL)
	if err != nil {
		return ssi.URI{}, "", fmt.Errorf("unable to unmarshal single URL from service (id=%s): %w", services[0].ID.String(), err)
	}
	return services[0].ID, endpointURL, nil
}

// ControllersAsString returns all DID controllers as a string array
func (d *DocumentImpl) ControllersAsString() []string {
	var controllers []string
	for _, controller := range d.Controller {
		controllers = append(controllers, controller.String())
	}
	return controllers
}

// Service represents a DID Service as specified by the DID Core specification (https://www.w3.org/TR/did-core/#service-endpoints).
type Service struct {
	ID              ssi.URI           `json:"id"`
	Type            string            `json:"type,omitempty"`
	ServiceEndpoint map[string]string `json:"serviceEndpoint,omitempty"`
}

func (s Service) MarshalJSON() ([]byte, error) {
	type alias Service
	tmp := alias(s)
	if data, err := json.Marshal(tmp); err != nil {
		return nil, err
	} else {
		return marshal.NormalizeDocument(data, marshal.Unplural(serviceEndpointKey))
	}
}

func (s *Service) UnmarshalJSON(data []byte) error {
	normalizedData, err := marshal.NormalizeDocument(data, pluralContext, marshal.PluralValueOrMap(serviceEndpointKey))
	if err != nil {
		return err
	}
	type alias Service
	var result alias
	if err := json.Unmarshal(normalizedData, &result); err != nil {
		return err
	}
	*s = (Service)(result)
	return nil
}

// Unmarshal unmarshalls the service endpoint into a domain-specific type.
func (s Service) UnmarshalServiceEndpoint(target interface{}) error {
	if asJSON, err := json.Marshal(s.ServiceEndpoint); err != nil {
		return err
	} else {
		return json.Unmarshal(asJSON, target)
	}
}

type Services []Service

// FindByID returns the first VerificationRelationship that matches with the id.
// For comparison both the ID of the embedded VerificationMethod and reference is used.
func (srs Services) FindByID(id ssi.URI) *Service {
	for _, r := range srs {
		if r.ID == id {
			return &r
		}
	}
	return nil
}

// VerificationMethod represents a DID Verification Method as specified by the DID Core specification (https://www.w3.org/TR/did-core/#verification-methods).
type VerificationMethod struct {
	ID              DID                    `json:"id"`
	Type            ssi.KeyType            `json:"type,omitempty"`
	Controller      DID                    `json:"controller,omitempty"`
	PublicKeyBase58 string                 `json:"publicKeyBase58,omitempty"`
	PublicKeyJwk    map[string]interface{} `json:"publicKeyJwk,omitempty"`
	Credential      *Credential            `json:"credential,omitempty"`
}

// NewVerificationMethod is a convenience method to easily create verificationMethods based on a set of given params.
// It automatically encodes the provided public key based on the keyType.
func NewVerificationMethod(id DID, keyType ssi.KeyType, controller DID, key crypto.PublicKey) (*VerificationMethod, error) {
	vm := &VerificationMethod{
		ID:         id,
		Type:       keyType,
		Controller: controller,
	}

	if keyType == ssi.JsonWebKey2020 {
		keyAsJWK, err := jwx.New(key).CreateEncJWK()
		if err != nil {
			return nil, err
		}
		// Convert to JSON and back to fix encoding of key material to make sure
		// an unmarshalled and newly created VerificationMethod are equal on object level.
		// The format of PublicKeyJwk in verificationMethod is a map[string]interface{}.
		// We can't use the Key.AsMap since the values of the map will all be internal jwk lib structs.
		// After unmarshalling all the fields will be map[string]string.
		keyAsJSON, err := json.Marshal(keyAsJWK)
		if err != nil {
			return nil, err
		}
		keyAsMap := map[string]interface{}{}
		json.Unmarshal(keyAsJSON, &keyAsMap)

		vm.PublicKeyJwk = keyAsMap
	}
	if keyType == ssi.ED25519VerificationKey2018 {
		ed25519Key, ok := key.(ed25519.PublicKey)
		if !ok {
			return nil, errors.New("wrong key type")
		}
		encodedKey := base58.Encode(ed25519Key, base58.BitcoinAlphabet)
		vm.PublicKeyBase58 = encodedKey
	}

	return vm, nil
}

// VerificationRelationship represents the usage of a VerificationMethod e.g. in authentication, assertionMethod, or keyAgreement.
type VerificationRelationship struct {
	*VerificationMethod
	reference DID
}

func (v VerificationRelationship) MarshalJSON() ([]byte, error) {
	if v.reference.Empty() {
		return json.Marshal(*v.VerificationMethod)
	} else {
		return json.Marshal(v.reference)
	}
}

func (v *VerificationRelationship) UnmarshalJSON(b []byte) error {
	// try to figure out if the item is an object of a string
	type Alias VerificationRelationship
	switch b[0] {
	case '{':
		tmp := Alias{VerificationMethod: &VerificationMethod{}}
		err := json.Unmarshal(b, &tmp)
		if err != nil {
			return fmt.Errorf("could not parse verificationRelation method: %w", err)
		}
		*v = (VerificationRelationship)(tmp)
	case '"':
		err := json.Unmarshal(b, &v.reference)
		if err != nil {
			return fmt.Errorf("could not parse verificationRelation key relation DID: %w", err)
		}
	default:
		return errors.New("verificationRelation is invalid")
	}
	return nil
}

func (v *VerificationMethod) UnmarshalJSON(bytes []byte) error {
	type Alias VerificationMethod
	tmp := Alias{}
	err := json.Unmarshal(bytes, &tmp)
	if err != nil {
		return err
	}
	*v = (VerificationMethod)(tmp)
	return nil
}

func resolveVerificationRelationships(relationships []VerificationRelationship, methods []*VerificationMethod) error {
	for i, relationship := range relationships {
		if relationship.reference.Empty() {
			continue
		}
		if resolved := resolveVerificationRelationship(relationship.reference, methods); resolved == nil {
			return fmt.Errorf("unable to resolve %s: %s", verificationMethodKey, relationship.reference.String())
		} else {
			relationships[i] = *resolved
			relationships[i].reference = relationship.reference
		}
	}
	return nil
}

func resolveVerificationRelationship(reference DID, methods []*VerificationMethod) *VerificationRelationship {
	for _, method := range methods {
		if method.ID.Equals(reference) {
			return &VerificationRelationship{VerificationMethod: method}
		}
	}
	return nil
}
