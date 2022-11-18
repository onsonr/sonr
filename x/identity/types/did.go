package types

import (
	"crypto"
	"encoding/json"
	"errors"
	fmt "fmt"

	"github.com/sonr-io/sonr/x/identity/types/internal/marshal"
	"github.com/sonr-io/sonr/x/identity/types/ssi"
)

// NewDocument generates a new DID Document for the provided ID string
func NewDocument(idStr string) (*DidDocument, error) {
	fmt.Println(idStr)
	id, err := ParseDID(idStr)
	if err != nil {
		return nil, err
	}

	ctxUri, err := ssi.ParseURI("https://www.w3.org/ns/did/v1")
	if err != nil {
		return nil, err
	}

	return &DidDocument{
		ID:                   id.String(),
		Context:              []string{ctxUri.String()},
		Controller:           []string{id.String()},
		VerificationMethod:   new(VerificationMethods),
		Authentication:       new(VerificationRelationships),
		AssertionMethod:      new(VerificationRelationships),
		CapabilityInvocation: new(VerificationRelationships),
		CapabilityDelegation: new(VerificationRelationships),
		KeyAgreement:         new(VerificationRelationships),
		Service:              new(Services),
		AlsoKnownAs:          make([]string, 0),
	}, nil
}

func NewDocumentFromJson(b []byte) (*DidDocument, error) {
	var doc DidDocument
	err := doc.UnmarshalJSON(b)
	if err != nil {
		return nil, err
	}
	return &doc, nil
}

func (d *DidDocument) ControllerCount() int {
	return len(d.Controller)
}

// FindAuthenticationMethod finds a VerificationMethod by its ID
func (d *DidDocument) FindAuthenticationMethod(id string) *VerificationMethod {
	return d.Authentication.FindByID(id)
}

// FindAssertionMethod finds a VerificationMethod by its ID
func (d *DidDocument) FindAssertionMethod(id string) *VerificationMethod {
	return d.AssertionMethod.FindByID(id)
}

// FindCapabilityDelegation finds a VerificationMethod by its ID
func (d *DidDocument) FindCapabilityDelegation(id string) *VerificationMethod {
	return d.CapabilityDelegation.FindByID(id)
}

// FindCapabilityInvocation finds a VerificationMethod by its ID
func (d *DidDocument) FindCapabilityInvocation(id string) *VerificationMethod {
	return d.CapabilityInvocation.FindByID(id)
}

func (d *DidDocument) FindController(did string) (string, error) {
	for _, c := range d.Controller {
		if c == did {
			return c, nil
		}
	}
	return "", errors.New("did not found")
}

func (d *DidDocument) FindContext() []string {
	return d.Context
}

// AddController adds a DID as a controller
func (d *DidDocument) AddController(id string) {
	if d.Controller == nil {
		d.Controller = make([]string, 0)
	}
	d.Controller = append(d.Controller, id)
}

func (d *DidDocument) GetVerificationMethods() *VerificationMethods {
	return d.VerificationMethod
}

// FindByID find the first VerificationMethod which matches the provided DID.
// Returns nil when not found
func (vms VerificationMethods) FindByID(id string) *VerificationMethod {
	for _, vm := range vms.Data {
		if vm.ID == id {
			return vm
		}
	}
	return nil
}

// Remove removes a VerificationMethod from the slice.
// If a verificationMethod was removed with the given DID, it will be returned
func (vms *VerificationMethods) Remove(id string) *VerificationMethod {
	var (
		filteredVMS []*VerificationMethod
		foundVM     *VerificationMethod
	)
	for _, vm := range vms.Data {
		if vm.ID != id {
			filteredVMS = append(filteredVMS, vm)
		} else {
			foundVM = vm
		}
	}
	vms.Data = filteredVMS
	return foundVM
}

// Add adds a verificationMethod to the verificationMethods if it not already present.
func (vms *VerificationMethods) Add(v *VerificationMethod) {
	for _, ptr := range vms.Data {
		// check if the pointer is already in the list
		if ptr == v {
			return
		}
		// check if the actual ids match?
		if ptr.ID == v.ID {
			return
		}
	}
	vms.Data = append(vms.Data, v)
}

// Count returns the number of VerificationRelationships in the slice
func (vmr *VerificationRelationships) Count() int {
	return len(vmr.GetData())
}

// FindByID returns the first VerificationRelationship that matches with the id.
// For comparison both the ID of the embedded VerificationMethod and reference is used.
func (vmr *VerificationRelationships) FindByID(id string) *VerificationMethod {
	for _, r := range vmr.GetData() {
		if r.VerificationMethod != nil {
			if r.VerificationMethod.ID == id {
				return r.VerificationMethod
			}
		}
	}
	return nil
}

// Remove removes a VerificationRelationship from the slice.
// If a VerificationRelationship was removed with the given DID, it will be returned
func (vmr VerificationRelationships) Remove(id string) *VerificationRelationship {
	var (
		filteredVMRels []*VerificationRelationship
		removedRel     *VerificationRelationship
	)
	for _, r := range vmr.GetData() {
		if r.Reference == id {
			filteredVMRels = append(filteredVMRels, r)
		} else {
			removedRel = r
		}
	}
	vmr.Data = filteredVMRels
	return removedRel
}

// Add adds a verificationMethod to a relationship collection.
// When the collection already contains the method it will not be added again.
func (vmr *VerificationRelationships) Add(vm *VerificationMethod) {
	for _, rel := range vmr.GetData() {
		if rel.Reference == vm.ID {
			return
		}
	}
	vmr.Data = append(vmr.GetData(), &VerificationRelationship{vm, vm.ID})
}

// AddAuthenticationMethod adds a VerificationMethod as AuthenticationMethod
// If the controller is not set, it will be set to the document's ID
func (d *DidDocument) AddAuthenticationMethod(v *VerificationMethod) {
	if v.Controller == "" {
		v.Controller = d.ID
	}
	d.VerificationMethod.Add(v)
	d.Authentication.Add(v)
}

// AddAssertionMethod adds a VerificationMethod as AssertionMethod
// If the controller is not set, it will be set to the documents ID
func (d *DidDocument) AddAssertionMethod(v *VerificationMethod) {
	if v.Controller == "" {
		v.Controller = d.ID
	}
	d.VerificationMethod.Add(v)
	d.AssertionMethod.Add(v)
}

func (d *DidDocument) GetKeyAgreements() VerificationRelationships {
	return *d.KeyAgreement
}

// AddKeyAgreement adds a VerificationMethod as KeyAgreement
// If the controller is not set, it will be set to the document's ID
func (d *DidDocument) AddKeyAgreement(v *VerificationMethod) {
	if v.Controller == "" {
		v.Controller = d.ID
	}
	d.VerificationMethod.Add(v)
	d.KeyAgreement.Add(v)
}

// AddCapabilityInvocation adds a VerificationMethod as CapabilityInvocation
// If the controller is not set, it will be set to the document's ID
func (d *DidDocument) AddCapabilityInvocation(v *VerificationMethod) {
	if v.Controller == "" {
		v.Controller = d.ID
	}
	d.VerificationMethod.Add(v)
	d.CapabilityInvocation.Add(v)
}

// AddCapabilityDelegation adds a VerificationMethod as CapabilityDelegation
// If the controller is not set, it will be set to the document's ID
func (d *DidDocument) AddCapabilityDelegation(v *VerificationMethod) {
	if v.Controller == "" {
		v.Controller = d.ID
	}
	d.VerificationMethod.Add(v)
	d.CapabilityDelegation.Add(v)
}

func (d *DidDocument) AddService(s *Service) {
	d.Service.Data = append(d.Service.Data, s)
}

func (d *DidDocument) RemoveServiceByID(id string) bool {
	for i, s := range d.Service.Data {
		if s.ID == id {
			d.Service.Data = append(d.Service.Data[:i], d.Service.Data[i+1:]...)
			return true
		}
	}
	return false
}

func (d *DidDocument) MarshalJSON() ([]byte, error) {
	type alias *DidDocument
	tmp := alias(d)
	if data, err := json.Marshal(tmp); err != nil {
		return nil, err
	} else {
		return marshal.NormalizeDocument(data, marshal.Unplural(contextKey), marshal.Unplural(controllerKey))
	}
}

func (d *DidDocument) UnmarshalJSON(b []byte) error {
	type alias DidDocument
	normalizedDoc, err := marshal.NormalizeDocument(b, pluralContext, marshal.Plural(controllerKey))
	if err != nil {
		return err
	}
	doc := alias{}
	err = json.Unmarshal(normalizedDoc, &doc)
	if err != nil {
		return err
	}
	*d = (DidDocument)(doc)

	const errMsg = "unable to resolve all '%s' references: %w"
	if err = resolveVerificationRelationships(d.Authentication.Data, d.VerificationMethod.Data); err != nil {
		return fmt.Errorf(errMsg, authenticationKey, err)
	}
	if err = resolveVerificationRelationships(d.AssertionMethod.Data, d.VerificationMethod.Data); err != nil {
		return fmt.Errorf(errMsg, assertionMethodKey, err)
	}
	if err = resolveVerificationRelationships(d.KeyAgreement.Data, d.VerificationMethod.Data); err != nil {
		return fmt.Errorf(errMsg, keyAgreementKey, err)
	}
	if err = resolveVerificationRelationships(d.CapabilityInvocation.Data, d.VerificationMethod.Data); err != nil {
		return fmt.Errorf(errMsg, capabilityInvocationKey, err)
	}
	if err = resolveVerificationRelationships(d.CapabilityDelegation.Data, d.VerificationMethod.Data); err != nil {
		return fmt.Errorf(errMsg, capabilityDelegationKey, err)
	}
	return nil
}

// IsController returns whether the given DID is a controller of the DID document.
func (d *DidDocument) IsController(controller string) bool {
	if controller == "" {
		return false
	}
	for _, curr := range d.Controller {
		if curr == controller {
			return true
		}
	}

	return false
}

// AddAlias adds a string alias to the document for a .snr domain name into the AlsoKnownAs field
// in the document.
func (d *DidDocument) AddAlias(alias string) {
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
func (d *DidDocument) ResolveEndpointURL(serviceType string) (endpointID string, endpointURL string, err error) {
	var services []*Service
	for _, service := range d.Service.Data {
		if service.Type == serviceType {
			services = append(services, service)
		}
	}
	if len(services) == 0 {
		return "", "", fmt.Errorf("service not found (did=%s, type=%s)", d.ID, serviceType)
	}
	if len(services) > 1 {
		return "", "", fmt.Errorf("multiple services found (did=%s, type=%s)", d.ID, serviceType)
	}
	err = services[0].UnmarshalServiceEndpoint(&endpointURL)
	if err != nil {
		return "", "", fmt.Errorf("unable to unmarshal single URL from service (id=%s): %w", services[0].ID, err)
	}
	return services[0].ID, endpointURL, nil
}

// ControllersAsString returns all DID controllers as a string array
func (d *DidDocument) ControllersAsString() []string {
	var controllers []string
	for _, controller := range d.Controller {
		controllers = append(controllers, controller)
	}
	return controllers
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

// FindByID returns the first VerificationRelationship that matches with the id.
// For comparison both the ID of the embedded VerificationMethod and reference is used.
func (srs *Services) FindByID(id string) *Service {
	for _, r := range srs.Data {
		if r.ID == id {
			return r
		}
	}
	return nil
}

// NewVerificationMethod is a convenience method to easily create verificationMethods based on a set of given params.
// It automatically encodes the provided public key based on the keyType.
func NewVerificationMethod(id string, keyType KeyType, controller string, key crypto.PublicKey) (*VerificationMethod, error) {
	vm := &VerificationMethod{
		ID:         id,
		Type:       keyType,
		Controller: controller,
	}

	// if keyType == KeyType_KeyType_JSON_WEB_KEY_2020 {
	// 	keyAsJWK, err := jwx.New(key).CreateEncJWK()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	// Convert to JSON and back to fix encoding of key material to make sure
	// 	// an unmarshalled and newly created VerificationMethod are equal on object level.
	// 	// The format of PublicKeyJwk in verificationMethod is a map[string]interface{}.
	// 	// We can't use the Key.AsMap since the values of the map will all be internal jwk lib structs.
	// 	// After unmarshalling all the fields will be map[string]string.
	// 	keyAsJSON, err := json.Marshal(keyAsJWK)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	keyAsMap := map[string]interface{}{}
	// 	json.Unmarshal(keyAsJSON, &keyAsMap)

	// 	vm.PublicKeyJwk = keyAsMap
	// }
	// if keyType == ssi.ED25519VerificationKey2018 {
	// 	ed25519Key, ok := key.(ed25519.PublicKey)
	// 	if !ok {
	// 		return nil, errors.New("wrong key type")
	// 	}
	// 	encodedKey := base58.Encode(ed25519Key, base58.BitcoinAlphabet)
	// 	vm.PublicKeyBase58 = encodedKey
	// }

	return vm, nil
}

func (v VerificationRelationship) MarshalJSON() ([]byte, error) {
	if v.Reference != "" {
		return json.Marshal(*v.VerificationMethod)
	} else {
		return json.Marshal(v.Reference)
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
		err := json.Unmarshal(b, &v.Reference)
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

func resolveVerificationRelationships(relationships []*VerificationRelationship, methods []*VerificationMethod) error {
	for i, relationship := range relationships {
		if relationship.Reference != "" {
			continue
		}
		if resolved := resolveVerificationRelationship(relationship.Reference, methods); resolved == nil {
			return fmt.Errorf("unable to resolve %s: %s", verificationMethodKey, relationship.Reference)
		} else {
			relationships[i] = resolved
			relationships[i].Reference = relationship.Reference
		}
	}
	return nil
}

func resolveVerificationRelationship(reference string, methods []*VerificationMethod) *VerificationRelationship {
	for _, method := range methods {
		if method.ID == reference {
			return &VerificationRelationship{VerificationMethod: method}
		}
	}
	return nil
}
