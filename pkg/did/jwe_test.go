package did

import (
	"reflect"
	"testing"

	"github.com/sonr-io/sonr/pkg/did/ssi"
)

func TestDocument_CreateJWS(t *testing.T) {
	type fields struct {
		Context              []ssi.URI
		ID                   DID
		Controller           []DID
		VerificationMethod   VerificationMethods
		Authentication       VerificationRelationships
		AssertionMethod      VerificationRelationships
		KeyAgreement         VerificationRelationships
		CapabilityInvocation VerificationRelationships
		CapabilityDelegation VerificationRelationships
		Service              []Service
		AlsoKnownAs          []string
	}
	type args struct {
		buf []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Document{
				Context:              tt.fields.Context,
				ID:                   tt.fields.ID,
				Controller:           tt.fields.Controller,
				VerificationMethod:   tt.fields.VerificationMethod,
				Authentication:       tt.fields.Authentication,
				AssertionMethod:      tt.fields.AssertionMethod,
				KeyAgreement:         tt.fields.KeyAgreement,
				CapabilityInvocation: tt.fields.CapabilityInvocation,
				CapabilityDelegation: tt.fields.CapabilityDelegation,
				Service:              tt.fields.Service,
				AlsoKnownAs:          tt.fields.AlsoKnownAs,
			}
			got, err := d.CreateJWS(tt.args.buf)
			if (err != nil) != tt.wantErr {
				t.Errorf("Document.CreateJWS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Document.CreateJWS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDocument_EncryptJWE(t *testing.T) {
	type fields struct {
		Context              []ssi.URI
		ID                   DID
		Controller           []DID
		VerificationMethod   VerificationMethods
		Authentication       VerificationRelationships
		AssertionMethod      VerificationRelationships
		KeyAgreement         VerificationRelationships
		CapabilityInvocation VerificationRelationships
		CapabilityDelegation VerificationRelationships
		Service              []Service
		AlsoKnownAs          []string
	}
	type args struct {
		buf []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Document{
				Context:              tt.fields.Context,
				ID:                   tt.fields.ID,
				Controller:           tt.fields.Controller,
				VerificationMethod:   tt.fields.VerificationMethod,
				Authentication:       tt.fields.Authentication,
				AssertionMethod:      tt.fields.AssertionMethod,
				KeyAgreement:         tt.fields.KeyAgreement,
				CapabilityInvocation: tt.fields.CapabilityInvocation,
				CapabilityDelegation: tt.fields.CapabilityDelegation,
				Service:              tt.fields.Service,
				AlsoKnownAs:          tt.fields.AlsoKnownAs,
			}
			got, err := d.EncryptJWE(tt.args.buf)
			if (err != nil) != tt.wantErr {
				t.Errorf("Document.EncryptJWE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Document.EncryptJWE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDocument_DecryptJWE(t *testing.T) {
	type fields struct {
		Context              []ssi.URI
		ID                   DID
		Controller           []DID
		VerificationMethod   VerificationMethods
		Authentication       VerificationRelationships
		AssertionMethod      VerificationRelationships
		KeyAgreement         VerificationRelationships
		CapabilityInvocation VerificationRelationships
		CapabilityDelegation VerificationRelationships
		Service              []Service
		AlsoKnownAs          []string
	}
	type args struct {
		serial string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Document{
				Context:              tt.fields.Context,
				ID:                   tt.fields.ID,
				Controller:           tt.fields.Controller,
				VerificationMethod:   tt.fields.VerificationMethod,
				Authentication:       tt.fields.Authentication,
				AssertionMethod:      tt.fields.AssertionMethod,
				KeyAgreement:         tt.fields.KeyAgreement,
				CapabilityInvocation: tt.fields.CapabilityInvocation,
				CapabilityDelegation: tt.fields.CapabilityDelegation,
				Service:              tt.fields.Service,
				AlsoKnownAs:          tt.fields.AlsoKnownAs,
			}
			got, err := d.DecryptJWE(tt.args.serial)
			if (err != nil) != tt.wantErr {
				t.Errorf("Document.DecryptJWE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Document.DecryptJWE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDocument_VerifyJWS(t *testing.T) {
	type fields struct {
		Context              []ssi.URI
		ID                   DID
		Controller           []DID
		VerificationMethod   VerificationMethods
		Authentication       VerificationRelationships
		AssertionMethod      VerificationRelationships
		KeyAgreement         VerificationRelationships
		CapabilityInvocation VerificationRelationships
		CapabilityDelegation VerificationRelationships
		Service              []Service
		AlsoKnownAs          []string
	}
	type args struct {
		serial string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Document{
				Context:              tt.fields.Context,
				ID:                   tt.fields.ID,
				Controller:           tt.fields.Controller,
				VerificationMethod:   tt.fields.VerificationMethod,
				Authentication:       tt.fields.Authentication,
				AssertionMethod:      tt.fields.AssertionMethod,
				KeyAgreement:         tt.fields.KeyAgreement,
				CapabilityInvocation: tt.fields.CapabilityInvocation,
				CapabilityDelegation: tt.fields.CapabilityDelegation,
				Service:              tt.fields.Service,
				AlsoKnownAs:          tt.fields.AlsoKnownAs,
			}
			got, err := d.VerifyJWS(tt.args.serial)
			if (err != nil) != tt.wantErr {
				t.Errorf("Document.VerifyJWS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Document.VerifyJWS() = %v, want %v", got, tt.want)
			}
		})
	}
}
