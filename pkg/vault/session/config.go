package session

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
	"github.com/sonr-hq/sonr/x/identity/types"
)

// Default Variables
var (
	// Default Origins
	defaultRpOrigins = []string{
		"https://auth.sonr.io",
		"https://sonr.id",
		"https://sandbox.sonr.network",
		"http://localhost:3000",
	}

	// Default Icon to display
	defaultRpIcon = "https://raw.githubusercontent.com/sonr-hq/sonr/master/docs/static/favicon.png"

	// Default name to display
	defaultRpName = "Sonr"

	// defaultAuthSelect
	defaultAuthSelect = protocol.AuthenticatorSelection{
		AuthenticatorAttachment: protocol.AuthenticatorAttachment("platform"),
		RequireResidentKey:      protocol.ResidentKeyUnrequired(),
		UserVerification:        protocol.VerificationRequired,
	}
)

func makeDefaultRandomVars() (*types.VerificationMethod, error) {
	sessionID := uuid.New().String()[:8]
	vm := &types.VerificationMethod{
		ID:   sessionID,
		Type: types.KeyType_KeyType_WEB_AUTHN_AUTHENTICATION_2018,
	}
	return vm, nil
}

// It takes a JSON string, converts it to a struct, and then converts that struct to a different struct
func getParsedCredentialCreationData(bz string) (*protocol.ParsedCredentialCreationData, error) {
	// Get Credential Creation Respons
	ccr := protocol.CredentialCreationResponse{}
	err := json.Unmarshal([]byte(bz), &ccr)
	if err != nil {
		return nil, err
	}

	// Parse the response
	var pcc protocol.ParsedCredentialCreationData
	pcc.ID, pcc.RawID, pcc.Type, pcc.ClientExtensionResults = ccr.ID, ccr.RawID, ccr.Type, ccr.ClientExtensionResults
	pcc.Raw = ccr

	// Parse the attestation object
	for _, t := range ccr.Transports {
		pcc.Transports = append(pcc.Transports, protocol.AuthenticatorTransport(t))
	}

	// Parse the attestation object
	parsedAttestationResponse, err := ccr.AttestationResponse.Parse()
	if err != nil {
		return nil, err
	}

	pcc.Response = *parsedAttestationResponse
	return &pcc, nil
}

func getParsedCredentialRequestData(bz string) (*protocol.ParsedCredentialAssertionData, error) {
	car := protocol.CredentialAssertionResponse{}
	err := json.Unmarshal([]byte(bz), &car)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, errors.New("Parse error for Assertion")
	}

	if car.ID == "" {
		return nil, errors.New("CredentialAssertionResponse with ID missing")
	}

	_, err = base64.RawURLEncoding.DecodeString(car.ID)
	if err != nil {
		return nil, errors.New("CredentialAssertionResponse with ID not base64url encoded")
	}
	if car.Type != "public-key" {
		return nil, errors.New("CredentialAssertionResponse with bad type")
	}
	var par protocol.ParsedCredentialAssertionData
	par.ID, par.RawID, par.Type, par.ClientExtensionResults = car.ID, car.RawID, car.Type, car.ClientExtensionResults
	par.Raw = car

	par.Response.Signature = car.AssertionResponse.Signature
	par.Response.UserHandle = car.AssertionResponse.UserHandle

	// Step 5. Let JSONtext be the result of running UTF-8 decode on the value of cData.
	// We don't call it cData but this is Step 5 in the spec.
	err = json.Unmarshal(car.AssertionResponse.ClientDataJSON, &par.Response.CollectedClientData)
	if err != nil {
		return nil, err
	}

	err = par.Response.AuthenticatorData.Unmarshal(car.AssertionResponse.AuthenticatorData)
	if err != nil {
		return nil, errors.New("Error unmarshalling auth data")
	}
	return &par, nil
}

func (s *SessionEntry) SetRPID(copts *protocol.CredentialCreation) *protocol.CredentialCreation {
	oldRp := copts.Response.RelyingParty
	newRp := protocol.RelyingPartyEntity{}
	newRp.ID = s.ID
	newRp.CredentialEntity = oldRp.CredentialEntity
	newCopts := copts
	newCopts.Response.RelyingParty = newRp
	return newCopts
}

func (s *SessionEntry) WebAuthn() (*webauthn.WebAuthn, error) {
	rawRpId := "localhost"
	if len(s.RPID) > 0 {
		rawRpId = s.RPID
	}
	// Create the Webauthn Instance
	return webauthn.New(&webauthn.Config{
		RPDisplayName:          defaultRpName,
		RPID:                   rawRpId,
		RPIcon:                 defaultRpIcon,
		RPOrigins:              defaultRpOrigins,
		Timeout:                60000,
		AttestationPreference:  protocol.PreferDirectAttestation,
		AuthenticatorSelection: defaultAuthSelect,
		Debug:                  true,
	})
}
