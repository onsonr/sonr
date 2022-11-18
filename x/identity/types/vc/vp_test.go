package vc

import (
	"encoding/json"
	"testing"

	ssi "github.com/sonr-io/sonr/x/identity/types/ssi"

	"github.com/stretchr/testify/assert"
)

func TestVerifiablePresentation_MarshalJSON(t *testing.T) {
	t.Run("ok - single credential and proof", func(t *testing.T) {
		input := VerifiablePresentation{
			VerifiableCredential: []VerifiableCredential{
				{
					Type: []ssi.URI{VerifiableCredentialTypeV1URI()},
				},
			},
			Proof: []interface{}{
				JSONWebSignature2020Proof{
					Jws: "",
				},
			},
		}

		bytes, err := json.Marshal(input)

		if !assert.NoError(t, err) {
			return
		}
		assert.Contains(t, string(bytes), "\"proof\":{")
		assert.Contains(t, string(bytes), "\"verifiableCredential\":{")
	})

	t.Run("ok - multiple credential and proof", func(t *testing.T) {
		input := VerifiablePresentation{
			VerifiableCredential: []VerifiableCredential{
				{
					Type: []ssi.URI{VerifiableCredentialTypeV1URI()},
				},
				{
					Type: []ssi.URI{VerifiableCredentialTypeV1URI()},
				},
			},
			Proof: []interface{}{
				JSONWebSignature2020Proof{
					Jws: "",
				},
				JSONWebSignature2020Proof{
					Jws: "",
				},
			},
		}

		bytes, err := json.Marshal(input)

		if !assert.NoError(t, err) {
			return
		}
		assert.Contains(t, string(bytes), "\"proof\":[")
		assert.Contains(t, string(bytes), "\"verifiableCredential\":[")
	})
}

func TestVerifiablePresentation_UnmarshalProof(t *testing.T) {
	type jsonWebSignature struct {
		Jws string
	}
	t.Run("ok - single proof", func(t *testing.T) {
		input := VerifiablePresentation{}
		json.Unmarshal([]byte(`{
		  "id":"did:example:123#vc-1",
		  "type":["VerifiablePresentation", "custom"],
		  "proof": {"jws": "test"}
		}`), &input)
		var target []jsonWebSignature

		err := input.UnmarshalProofValue(&target)

		assert.NoError(t, err)
		assert.Equal(t, "test", target[0].Jws)
	})

	t.Run("ok - multiple proof", func(t *testing.T) {
		input := VerifiablePresentation{}
		json.Unmarshal([]byte(`{
		  "id":"did:example:123#vc-1",
		  "type":["VerifiablePresentation", "custom"],
		  "proof": [{"jws": "test"}, {"not-jws": "test"}]
		}`), &input)
		var target []jsonWebSignature

		err := input.UnmarshalProofValue(&target)

		assert.NoError(t, err)
		assert.Len(t, target, 2)
		assert.Equal(t, "test", target[0].Jws)
		assert.Equal(t, "", target[1].Jws)
	})
}

func TestVerifiablePresentation_Proofs(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		input := VerifiablePresentation{}
		json.Unmarshal([]byte(`{
		  "id":"did:example:123#vc-1",
		  "type":["VerifiablePresentation", "custom"],
		  "proof": [{"type": "JsonWebSignature2020"}, {"type": "other"}]
		}`), &input)

		proofs, err := input.Proofs()

		assert.NoError(t, err)
		assert.Len(t, proofs, 2)
		assert.Equal(t, ssi.JsonWebSignature2020, proofs[0].Type)
	})
}

func TestVerifiablePresentation_IsType(t *testing.T) {
	input := VerifiablePresentation{}
	json.Unmarshal([]byte(`{
		  "id":"did:example:123#vp-1",
		  "type":"VerifiablePresentation"
		}`), &input)

	t.Run("true", func(t *testing.T) {
		assert.True(t, input.IsType(VerifiablePresentationTypeV1URI()))
	})

	t.Run("false", func(t *testing.T) {
		u, _ := ssi.ParseURI("type")
		assert.False(t, input.IsType(*u))
	})
}

func TestVerifiablePresentation_ContainsContext(t *testing.T) {
	input := VerifiablePresentation{}
	json.Unmarshal([]byte(`{
		  "id":"did:example:123#vp-1",
		  "@context":["https://www.w3.org/2018/credentials/v1"]
		}`), &input)

	t.Run("true", func(t *testing.T) {
		assert.True(t, input.ContainsContext(VCContextV1URI()))
	})

	t.Run("false", func(t *testing.T) {
		u, _ := ssi.ParseURI("context")
		assert.False(t, input.ContainsContext(*u))
	})
}
