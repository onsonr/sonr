package types

import (
	"encoding/json"
	"fmt"

	"github.com/sonr-io/sonr/pkg/did"
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

var pluralContext = Plural(contextKey)

func NewDIDDocumentFromBytes(buf []byte) (*DIDDocument, error) {
	var doc DIDDocument
	err := doc.UnmarshalJSON(buf)
	if err != nil {
		return nil, err
	}
	return &doc, nil
}

func NewDIDDocumentFromPkg(doc did.Document) (*DIDDocument, error) {
	var err error
	res := &DIDDocument{}

	res.Context = make([]string, 0)
	for _, c := range doc.GetContext() {
		res.Context = append(res.Context, c.String())
	}

	res.Id = doc.GetID().String()

	res.Controller = doc.ControllersAsString()

	res.VerificationMethod = verificationMethodsFromPkg(doc.GetVerificationMethods())

	res.Authentication, err = stringifyVerificationRelationships(doc.GetAuthenticationMethods())
	if err != nil {
		return nil, err
	}

	res.AssertionMethod, err = stringifyVerificationRelationships(doc.GetAssertionMethods())
	if err != nil {
		return nil, err
	}

	res.CapabilityInvocation, err = stringifyVerificationRelationships(doc.GetCapabilityInvocations())
	if err != nil {
		return nil, err
	}

	res.CapabilityDelegation, err = stringifyVerificationRelationships(doc.GetCapabilityDelegations())
	if err != nil {
		return nil, err
	}

	res.KeyAgreement, err = stringifyVerificationRelationships(doc.GetKeyAgreements())
	if err != nil {
		return nil, err
	}

	res.Service = servicesFromPkg(doc.GetServices())

	res.AlsoKnownAs = doc.GetAlsoKnownAs()

	return res, nil
}

func verificationMethodsFromPkg(methods did.VerificationMethods) []*VerificationMethod {
	res := make([]*VerificationMethod, len(methods))
	for i, m := range methods {
		res[i] = &VerificationMethod{
			Id:              m.ID.ID,
			Type:            string(m.Type),
			Controller:      m.Controller.ID,
			PublicKeyJwk:    make([]*KeyValuePair, 0),
			PublicKeyBase58: m.PublicKeyBase58,
		}

		for k, pubk := range m.PublicKeyJwk {
			res[i].PublicKeyJwk = append(res[i].PublicKeyJwk, &KeyValuePair{
				Key:   k,
				Value: fmt.Sprintf("%v", pubk),
			})
		}
	}
	return res
}

func servicesFromPkg(srvs did.Services) []*Service {
	res := make([]*Service, 0)
	for _, s := range srvs {
		res = append(res, &Service{
			Id:              s.ID.String(),
			Type:            s.Type,
			ServiceEndpoint: mapToKVPair(s.ServiceEndpoint),
		})
	}

	return res
}

func stringifyVerificationRelationships(auth did.VerificationRelationships) ([]string, error) {
	res := make([]string, len(auth))
	for i, a := range auth {
		ab, err := json.Marshal(a)
		if err != nil {
			return nil, err
		}
		res[i] = string(ab)
	}
	return res, nil
}

func (d *DIDDocument) DID() *did.DID {
	id, _ := did.ParseDID(d.Id)
	return id
}

func (d DIDDocument) MarshalJSON() ([]byte, error) {
	type alias DIDDocument
	tmp := alias(d)
	if data, err := json.Marshal(tmp); err != nil {
		return nil, err
	} else {
		return NormalizeDocument(data, Unplural(contextKey), Unplural(controllerKey))
	}
}

func (d *DIDDocument) UnmarshalJSON(b []byte) error {
	type Alias DIDDocument
	normalizedDoc, err := NormalizeDocument(b, pluralContext, Plural(controllerKey))
	if err != nil {
		return err
	}
	doc := Alias{}
	err = json.Unmarshal(normalizedDoc, &doc)
	if err != nil {
		return err
	}
	*d = (DIDDocument)(doc)
	return nil
}

// CopyFromBytes unmarshals a JSON document from a byte slice and copies the data into the receiver.
func (d *DIDDocument) CopyFromBytes(b []byte) error {
	var tmp DIDDocument
	err := tmp.UnmarshalJSON(b)
	if err != nil {
		return err
	}
	*d = tmp
	return nil
}

// IsController returns whether the given DID is a controller of the DID document.
func (d DIDDocument) IsController(controller string) bool {
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
func (d *DIDDocument) AddAlias(alias string) {
	if d.AlsoKnownAs == nil {
		d.AlsoKnownAs = make([]string, 0)
	}
	d.AlsoKnownAs = append(d.AlsoKnownAs, alias)
}

// NormalizeDocument accepts a JSON document and applies (in order) the given normalizers to it.
func NormalizeDocument(document []byte, normalizers ...Normalizer) ([]byte, error) {
	tmp := make(map[string]interface{}, 0)
	if err := json.Unmarshal(document, &tmp); err != nil {
		return nil, err
	}
	for _, normalizer := range normalizers {
		normalizer(tmp)
	}
	return json.Marshal(tmp)
}

type Normalizer func(map[string]interface{})

// KeyAlias returns a Normalizer that converts an aliased key to its original form. E.g. when working with
// LinkedData in JSON form, `@context` is an alias for `context`. This Normalizer would convert the `@context` key
// to `context`.
func KeyAlias(alias string, aliasFor string) Normalizer {
	return func(m map[string]interface{}) {
		for k, v := range m {
			if k == alias {
				m[aliasFor] = v
				delete(m, k)
			}
		}
	}
}

// Plural returns a Normalizer that converts a singular values (string/numeric/bool/object) to an array.
// This makes unmarshalling DID Documents or Verifiable Credentials easier, since those formats allow certain properties
// to be either a singular value or an array of values.
//
// Example input: 												{"message": "Hello, World"}
// Example output (if 'message' is supplied in 'pluralKeys'): 	{"message": ["Hello, World"]}
//
// This function does not support nested keys.
func Plural(key string) Normalizer {
	return func(m map[string]interface{}) {
		if _, isSlice := m[key].([]interface{}); m[key] != nil && !isSlice {
			m[key] = []interface{}{m[key]}
		}
	}
}

// Unplural returns a Normalizer that converts arrays with a single value into a singular value. It is the opposite
// of the Plural normalizer.
func Unplural(key string) Normalizer {
	return func(m map[string]interface{}) {
		if arr, _ := m[key].([]interface{}); len(arr) == 1 {
			m[key] = arr[0]
		}
	}
}

// PluralValueOrMap returns a Normalizer that behaves like Plural but leaves maps as simply a map. In other words,
// it only turns singular values into an array, except maps.
func PluralValueOrMap(key string) Normalizer {
	return func(m map[string]interface{}) {
		value := m[key]
		if value == nil {
			return
		} else if _, isMap := value.(map[string]interface{}); isMap {
			return
		} else if _, isSlice := value.([]interface{}); !isSlice {
			m[key] = []interface{}{m[key]}
		}
	}
}

func mapToKVPair(m map[string]string) []*KeyValuePair {
	res := make([]*KeyValuePair, 0)
	for k, v := range m {
		res = append(res, &KeyValuePair{
			Key:   k,
			Value: v,
		})
	}

	return res
}
