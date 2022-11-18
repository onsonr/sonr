package vc

import (
	"encoding/json"

	"github.com/sonr-io/sonr/x/identity/types/internal/marshal"
	ssi "github.com/sonr-io/sonr/x/identity/types/ssi"
)

// VerifiablePresentationType is the default credential type required for every credential
const VerifiablePresentationType = "VerifiablePresentation"

// VerifiablePresentationTypeV1URI returns VerifiablePresentation as URI
func VerifiablePresentationTypeV1URI() ssi.URI {
	return ssi.MustParseURI(VerifiablePresentationType)
}

// VerifiablePresentation represents a presentation as defined by the Verifiable Credentials Data Model 1.0 specification (https://www.w3.org/TR/vc-data-model/).
type VerifiablePresentation struct {
	// Context defines the json-ld context to dereference the URIs
	Context []ssi.URI `json:"@context"`
	// ID is an unique identifier for the presentation. It is optional
	ID *ssi.URI `json:"id,omitempty"`
	// Type holds multiple types for a presentation. A presentation must always have the 'VerifiablePresentation' type.
	Type []ssi.URI `json:"type"`
	// Holder refers to the party that generated the presentation. It is optional
	Holder *ssi.URI `json:"holder,omitempty"`
	// VerifiableCredential may hold credentials that are proven with this presentation.
	VerifiableCredential []VerifiableCredential `json:"verifiableCredential,omitempty"`
	// Proof contains the cryptographic proof(s). It must be extracted using the Proofs method or UnmarshalProofValue method for non-generic proof fields.
	Proof []interface{} `json:"proof,omitempty"`
}

// Proofs returns the basic proofs for this presentation. For specific proof contents, UnmarshalProofValue must be used.
func (vp VerifiablePresentation) Proofs() ([]Proof, error) {
	var (
		target []Proof
		err    error
		asJSON []byte
	)
	asJSON, err = json.Marshal(vp.Proof)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(asJSON, &target)
	return target, err
}

func (vp VerifiablePresentation) MarshalJSON() ([]byte, error) {
	type alias VerifiablePresentation
	tmp := alias(vp)
	if data, err := json.Marshal(tmp); err != nil {
		return nil, err
	} else {
		return marshal.NormalizeDocument(data, pluralContext, marshal.Unplural(typeKey), marshal.Unplural(verifiableCredentialKey), marshal.Unplural(proofKey))
	}
}

func (vp *VerifiablePresentation) UnmarshalJSON(b []byte) error {
	type Alias VerifiablePresentation
	normalizedVC, err := marshal.NormalizeDocument(b, pluralContext, marshal.Plural(typeKey), marshal.Plural(verifiableCredentialKey), marshal.Plural(proofKey))
	if err != nil {
		return err
	}
	tmp := Alias{}
	err = json.Unmarshal(normalizedVC, &tmp)
	if err != nil {
		return err
	}
	*vp = (VerifiablePresentation)(tmp)
	return nil
}

// UnmarshalProofValue unmarshalls the proof to the given proof type. Always pass a slice as target since there could be multiple proofs.
// Each proof will result in a value, where null values may exist when the proof doesn't have the json member.
func (vp VerifiablePresentation) UnmarshalProofValue(target interface{}) error {
	if asJSON, err := json.Marshal(vp.Proof); err != nil {
		return err
	} else {
		return json.Unmarshal(asJSON, target)
	}
}

// IsType returns true when a presentation contains the requested type
func (vp VerifiablePresentation) IsType(vcType ssi.URI) bool {
	for _, t := range vp.Type {
		if t.String() == vcType.String() {
			return true
		}
	}

	return false
}

// ContainsContext returns true when a credential contains the requested context
func (vp VerifiablePresentation) ContainsContext(context ssi.URI) bool {
	for _, c := range vp.Context {
		if c.String() == context.String() {
			return true
		}
	}

	return false
}
