package builder

import (
	"encoding/json"

	"github.com/onsonr/sonr/x/did/types"
)

type AuthenticatorResponse struct {
	// From the spec https://www.w3.org/TR/webauthn/#dom-authenticatorresponse-clientdatajson
	// This attribute contains a JSON serialization of the client data passed to the authenticator
	// by the client in its call to either create() or get().
	ClientDataJSON URLEncodedBase64 `json:"clientDataJSON"`
}

type AuthenticatorAttestationResponse struct {
	// The byte slice of clientDataJSON, which becomes CollectedClientData
	AuthenticatorResponse

	Transports []string `json:"transports,omitempty"`

	AuthenticatorData URLEncodedBase64 `json:"authenticatorData"`

	PublicKey URLEncodedBase64 `json:"publicKey"`

	PublicKeyAlgorithm int64 `json:"publicKeyAlgorithm"`

	// AttestationObject is the byte slice version of attestationObject.
	// This attribute contains an attestation object, which is opaque to, and
	// cryptographically protected against tampering by, the client. The
	// attestation object contains both authenticator data and an attestation
	// statement. The former contains the AAGUID, a unique credential ID, and
	// the credential public key. The contents of the attestation statement are
	// determined by the attestation statement format used by the authenticator.
	// It also contains any additional information that the Relying Party's server
	// requires to validate the attestation statement, as well as to decode and
	// validate the authenticator data along with the JSON-serialized client data.
	AttestationObject URLEncodedBase64 `json:"attestationObject"`
}

type PublicKeyCredentialCreationOptions struct {
	RelyingParty           RelyingPartyEntity         `json:"rp"`
	User                   UserEntity                 `json:"user"`
	Challenge              URLEncodedBase64           `json:"challenge"`
	Parameters             []CredentialParameter      `json:"pubKeyCredParams,omitempty"`
	Timeout                int                        `json:"timeout,omitempty"`
	CredentialExcludeList  []CredentialDescriptor     `json:"excludeCredentials,omitempty"`
	AuthenticatorSelection AuthenticatorSelection     `json:"authenticatorSelection,omitempty"`
	Hints                  []PublicKeyCredentialHints `json:"hints,omitempty"`
	Attestation            ConveyancePreference       `json:"attestation,omitempty"`
	AttestationFormats     []AttestationFormat        `json:"attestationFormats,omitempty"`
	Extensions             AuthenticationExtensions   `json:"extensions,omitempty"`
}

func NewRegistrationOptions(origin string, subject string, vaultCID string, params *types.Params) (*PublicKeyCredentialCreationOptions, error) {
	chal, err := CreateChallenge()
	if err != nil {
		return nil, err
	}
	return &PublicKeyCredentialCreationOptions{
		RelyingParty:           NewRelayingParty(origin, subject),
		User:                   NewUserEntity(subject, subject, vaultCID),
		Parameters:             ExtractCredentialParameters(params),
		Timeout:                20,
		CredentialExcludeList:  nil,
		Challenge:              chal,
		AuthenticatorSelection: AuthenticatorSelection{},
		Hints:                  nil,
		Attestation:            ExtractConveyancePreference(params),
		AttestationFormats:     ExtractAttestationFormats(params),
		Extensions:             nil,
	}, nil
}

func UnmarshalAuthenticatorResponse(data []byte) (*AuthenticatorResponse, error) {
	var ar AuthenticatorResponse
	err := json.Unmarshal(data, &ar)
	if err != nil {
		return nil, err
	}
	return &ar, nil
}
