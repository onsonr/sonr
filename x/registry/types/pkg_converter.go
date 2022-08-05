package types

import (
	"encoding/json"
	"fmt"

	"github.com/sonr-io/sonr/pkg/crypto/did"
	"github.com/sonr-io/sonr/pkg/crypto/did/ssi"
)

func (d *DIDDocument) ToPkgDoc() (did.Document, error) {

	ctx, err := convertContext(d.Context)
	if err != nil {
		return nil, err
	}

	vm, err := convertVerificationMethods(d.VerificationMethod)
	if err != nil {
		return nil, err
	}

	auth, err := convertVerificationRelationships(d.Authentication)
	if err != nil {
		return nil, err
	}

	asrt, err := convertVerificationRelationships(d.AssertionMethod)
	if err != nil {
		return nil, err
	}

	keyAgreement, err := convertVerificationRelationships(d.KeyAgreement)
	if err != nil {
		return nil, err
	}

	capInv, err := convertVerificationRelationships(d.CapabilityInvocation)
	if err != nil {
		return nil, err
	}

	capDel, err := convertVerificationRelationships(d.CapabilityDelegation)
	if err != nil {
		return nil, err
	}

	services, err := convertServices(d.Service)
	if err != nil {
		return nil, err
	}

	return &did.DocumentImpl{
		ID:                   did.MustParseDID(d.Id),
		Context:              ctx,
		VerificationMethod:   vm,
		Authentication:       auth,
		AssertionMethod:      asrt,
		KeyAgreement:         keyAgreement,
		CapabilityInvocation: capInv,
		CapabilityDelegation: capDel,
		Service:              services,
		AlsoKnownAs:          d.AlsoKnownAs,
	}, nil
}

func convertContext(c []string) ([]ssi.URI, error) {
	res := make([]ssi.URI, len(c))
	for i, ctx := range c {
		uri, err := ssi.ParseURI(ctx)
		if err != nil {
			return nil, err
		}
		res[i] = *uri
	}
	return res, nil
}

func convertVerificationMethods(methods []*VerificationMethod) (did.VerificationMethods, error) {
	res := make(did.VerificationMethods, len(methods))
	for i, m := range methods {
		var cred did.Credential
		err := json.Unmarshal(m.CredentialJson, &cred)
		if err != nil {
			return nil, err
		}
		res[i] = &did.VerificationMethod{
			ID:              did.MustParseDID(m.Id),
			Type:            ssi.KeyType(m.Type),
			Controller:      did.MustParseDID(m.Controller),
			PublicKeyBase58: m.PublicKeyBase58,
			PublicKeyJwk:    convertKeyValuePair(m.PublicKeyJwk),
			Credential:      &cred,
		}
	}

	return res, nil
}

func convertVerificationRelationships(relationships []string) (did.VerificationRelationships, error) {
	res := make(did.VerificationRelationships, len(relationships))
	for i, r := range relationships {
		var v did.VerificationRelationship
		err := json.Unmarshal([]byte(r), &v)
		if err != nil {
			return nil, err
		}

		res[i] = v
	}

	return res, nil
}

func convertServices(srvs []*Service) (did.Services, error) {
	res := make(did.Services, len(srvs))
	for i, s := range srvs {
		endpoints := make(map[string]string)
		for k, v := range convertKeyValuePair(s.ServiceEndpoint) {
			str, ok := v.(string)
			if !ok {
				return nil, fmt.Errorf("could not convert %v to string", v)
			}
			endpoints[k] = str
		}
		res[i] = did.Service{
			ID:              ssi.MustParseURI(s.Id),
			Type:            s.Type,
			ServiceEndpoint: endpoints,
		}
	}

	return res, nil
}

func convertKeyValuePair(kvp []*KeyValuePair) map[string]interface{} {
	m := make(map[string]interface{})
	for _, v := range kvp {
		m[v.Key] = v.Value
	}
	return m
}
