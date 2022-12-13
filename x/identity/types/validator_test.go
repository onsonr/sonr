package types

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestW3CSpecValidator(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		assert.NoError(t, W3CSpecValidator{}.Validate(document()))
	})
	t.Run("base", func(t *testing.T) {

		didUrl, err := ParseDID("did:snr:123#fragment")
		if !assert.NoError(t, err) {
			return
		}

		t.Run("context is missing DIDv1", func(t *testing.T) {
			input := document()
			input.Context = []string{}
			assertIsError(t, ErrInvalidContext, W3CSpecValidator{}.Validate(input))
		})
		t.Run("invalid ID - is empty", func(t *testing.T) {
			input := document()
			input.ID = ""
			assertIsError(t, ErrInvalidID, W3CSpecValidator{}.Validate(input))
		})
		t.Run("invalid ID - is URL", func(t *testing.T) {
			input := document()
			input.ID = "https://example.com"
			assertIsError(t, ErrInvalidID, W3CSpecValidator{}.Validate(input))
		})

		t.Run("invalid controller - is empty", func(t *testing.T) {
			input := document()
			input.Controller = append(make([]string, 0), "")
			assertIsError(t, ErrInvalidController, W3CSpecValidator{}.Validate(input))
		})

		t.Run("invalid controller - is URL", func(t *testing.T) {
			input := document()

			input.Controller = append(make([]string, 1), didUrl.String())
			assertIsError(t, ErrInvalidController, W3CSpecValidator{}.Validate(input))
		})
	})
	t.Run("verificationMethod", func(t *testing.T) {
		t.Run("invalid ID", func(t *testing.T) {
			input := document()
			input.VerificationMethod.Data[0].ID = ""
			assertIsError(t, ErrInvalidVerificationMethod, W3CSpecValidator{}.Validate(input))
		})
		t.Run("invalid controller", func(t *testing.T) {
			input := document()
			input.VerificationMethod.Data[0].Controller = ""
			assertIsError(t, ErrInvalidVerificationMethod, W3CSpecValidator{}.Validate(input))
		})
		t.Run("invalid type", func(t *testing.T) {
			input := document()
			input.VerificationMethod.Data[0].Type = KeyType_KeyType_UNSPECIFIED
			assertIsError(t, ErrInvalidVerificationMethod, W3CSpecValidator{}.Validate(input))
		})
	})
	t.Run("authentication", func(t *testing.T) {
		t.Run("invalid ID", func(t *testing.T) {
			input := document()
			// Make copy first because it is a reference instead of embedded
			vm := *input.VerificationMethod.Data[0]
			input.Authentication.Data[0] = &VerificationRelationship{VerificationMethod: &vm}
			// Then alter
			input.Authentication.Data[0].VerificationMethod.ID = ""
			input.Authentication.Data[0].Reference = ""
			assertIsError(t, ErrInvalidAuthentication, W3CSpecValidator{}.Validate(input))
		})
		t.Run("invalid controller", func(t *testing.T) {
			input := document()
			// Make copy first because it is a reference instead of embedded
			vm := *input.VerificationMethod.Data[0]
			input.Authentication.Data[0] = &VerificationRelationship{VerificationMethod: &vm}
			// Then alter
			input.Authentication.Data[0].VerificationMethod.Controller = ""
			input.Authentication.Data[0].Reference = "did:snr:123#fragment"
			assertIsError(t, ErrInvalidAuthentication, W3CSpecValidator{}.Validate(input))
		})
	})
	t.Run("service", func(t *testing.T) {
		t.Run("invalid ID", func(t *testing.T) {
			input := document()
			input.Service.Data[0].ID = ""
			assertIsError(t, ErrInvalidService, W3CSpecValidator{}.Validate(input))
		})
		t.Run("invalid type", func(t *testing.T) {
			input := document()
			input.Service.Data[0].Type = ServiceType_ServiceType_UNSPECIFIED
			assertIsError(t, ErrInvalidService, W3CSpecValidator{}.Validate(input))
		})
		t.Run("endpoint is nil", func(t *testing.T) {
			input := document()
			input.Service.Data[0].ServiceEndpoint = ""
			assertIsError(t, ErrInvalidService, W3CSpecValidator{}.Validate(input))
		})
	})
}

func TestMultiValidator(t *testing.T) {
	t.Run("no validators", func(t *testing.T) {
		assert.NoError(t, MultiValidator{}.Validate(document()))
	})
	t.Run("no errors", func(t *testing.T) {
		v := W3CSpecValidator{}
		assert.NoError(t, MultiValidator{Validators: []Validator{v, v}}.Validate(document()))
	})
	t.Run("returns first", func(t *testing.T) {
		v1 := W3CSpecValidator{}
		v2 := funcValidator{fn: func(_ *DidDocument) error {
			return errors.New("failed")
		}}
		assert.Error(t, MultiValidator{Validators: []Validator{v2, v1}}.Validate(document()))
	})
	t.Run("returns second", func(t *testing.T) {
		v1 := W3CSpecValidator{}
		v2 := funcValidator{fn: func(_ *DidDocument) error {
			return errors.New("failed")
		}}
		assert.Error(t, MultiValidator{Validators: []Validator{v1, v2}}.Validate(document()))
	})
}

func assertIsError(t *testing.T, expected error, actual error) {
	if !errors.Is(actual, expected) {
		t.Errorf("\ngot error: %v\nwanted error: %v", actual, expected)
	}
}

func document() *DidDocument {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	vm, _ := NewVerificationMethod("did:test:12345", KeyType_KeyType_JSON_WEB_KEY_2020, "", privateKey.Public())
	srv := &Service{
		ID:              "did:test:12345",
		Type:            ServiceType_ServiceType_UNSPECIFIED,
		ServiceEndpoint: "tcp://awesome-service",
	}
	doc := &DidDocument{
		Context: []string{DIDContextV1URI().String()},
		ID:      "did:test:12345",
		VerificationMethod: &VerificationMethods{
			Data: []*VerificationMethod{vm},
		},
		Authentication:       &VerificationRelationships{},
		AssertionMethod:      &VerificationRelationships{},
		CapabilityInvocation: &VerificationRelationships{},
		CapabilityDelegation: &VerificationRelationships{},
		KeyAgreement:         &VerificationRelationships{},
		Service: &Services{
			Data: []*Service{srv},
		},
	}
	doc.AddAuthenticationMethod(vm)
	doc.AddAssertionMethod(vm)
	return doc
}

type funcValidator struct {
	fn func(document *DidDocument) error
}

func (f funcValidator) Validate(document *DidDocument) error {
	return f.fn(document)
}
