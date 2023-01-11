// Utility functions for DID Service - https://w3c.github.io/did-core/#services
// I.e. Service Endpoints for IPFS Cluster
package types

import (
	"encoding/json"
	fmt "fmt"

	"github.com/sonr-hq/sonr/x/identity/types/internal/marshal"
)

func NewIPNSService(id string, endpoint string) *Service {
	return &Service{
		ID:              id,
		Type:            ServiceType_ServiceType_ENCRYPTED_DATA_VAULT,
		ServiceEndpoint: endpoint,
	}
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

// ResolveEndpointURL finds the endpoint with the given type and unmarshalls it as single URL.
// It returns the endpoint ID and URL, or an error if anything went wrong;
// - holder document can't be resolved,
// - service with given type doesn't exist,
// - multiple services match,
// - serviceEndpoint isn't a string.
func (d *DidDocument) ResolveEndpointURL(serviceType string) (endpointID string, endpointURL string, err error) {
	var services []*Service
	for _, service := range d.Service.Data {
		if service.Type.FormatString() == serviceType {
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
