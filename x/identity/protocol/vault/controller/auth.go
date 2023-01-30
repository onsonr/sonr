package controller

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/ucan-wg/go-ucan"
)

// This method is used to get the challenge response from the DID controller.
func (d *DIDControllerImpl) BeginRegistration(aka string) ([]byte, error) {
	if d.authentication != nil {
		return nil, errors.New("already registered")
	}
	audId := d.primaryAccount.DID() + "#" + aka
	d.authentication = &types.VerificationMethod{
		Id:         d.primaryAccount.DID() + "#" + aka,
		Controller: d.primaryAccount.DID(),
		Type:       "EcdsaSecp256k1VerificationKey2019",
	}
	caps := ucan.NewNestedCapabilities("DELEGATOR", "AUTHENTICATOR", "CREATE", "READ", "UPDATE")
	att := ucan.Attenuations{
		{Cap: caps.Cap("AUTHENTICATOR"), Rsc: ucan.NewStringLengthResource("mpc/acc", "*")},
		{Cap: caps.Cap("SUPER_USER"), Rsc: ucan.NewStringLengthResource("mpc/acc", "b5:world_bank_population:*")},
	}
	zero := time.Time{}
	origin, err := d.primaryAccount.NewOriginToken(audId, att, nil, zero, zero)
	if err != nil {
		return nil, err
	}
	challenge, err := d.authentication.IssueChallenge(origin)
	if err != nil {
		return nil, err
	}
	webAuthnUser := protocol.UserEntity{
		ID:          d.didDocument.WebAuthnID(),
		DisplayName: d.didDocument.WebAuthnDisplayName(),
		CredentialEntity: protocol.CredentialEntity{
			Name: d.didDocument.WebAuthnName(),
			Icon: d.didDocument.WebAuthnIcon(),
		},
	}

	relyingParty := protocol.RelyingPartyEntity{
		ID: "localhost",
		CredentialEntity: protocol.CredentialEntity{
			Name: d.didDocument.WebAuthnDisplayName(),
			Icon: defaultRpIcon,
		},
	}

	credentialParams := defaultRegistrationCredentialParameters()
	creationOptions := protocol.PublicKeyCredentialCreationOptions{
		Challenge:              challenge,
		RelyingParty:           relyingParty,
		User:                   webAuthnUser,
		Parameters:             credentialParams,
		AuthenticatorSelection: defaultAuthSelect,
		Timeout:                defaultTimeout,
		Attestation:            defaultAttestationPreference,
	}

	// Marshal the response into JSON.
	response := protocol.CredentialCreation{Response: creationOptions}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return jsonResponse, nil
}

// This is the method that will be called when the user clicks on the "Register" button.
func (d *DIDControllerImpl) FinishRegistration(aka string, challengResp string) ([]byte, error) {
	if d.authentication == nil {
		return nil, errors.New("no authentication method")
	}

	// Unmarshal the challenge response into a CredentialCreationResponse.
	pccd, err := parseCreationData(challengResp)
	if err != nil {
		return nil, err
	}
	cred := crypto.NewWebAuthnCredential(pccd)
	pk := crypto.NewWebAuthnPubKey(cred.PublicKey)
	vm, err := types.NewVMFromPubKey(pk, types.WithController(d.primaryAccount.DID()), types.WithIDFragmentSuffix(aka))
	if err != nil {
		return nil, err
	}
	d.didDocument.AddAuthentication(vm)
	didBz, err := d.didDocument.Marshal()
	if err != nil {
		return nil, err
	}
	didB64 := base64.StdEncoding.EncodeToString(didBz)
	return []byte(didB64), nil
}

// This method is used to get the options for the assertion.
func (d *DIDControllerImpl) BeginLogin(aka string) ([]byte, error) {
	vm := d.didDocument.GetVerificationMethodByFragment(aka)
	if vm == nil {
		return nil, errors.New("no authentication method")
	}
	d.authentication = vm
	return nil, nil
}

// This is the method that will be called when the user clicks the "Login" button on the login page.
func (d *DIDControllerImpl) FinishLogin(aka string, challengResp string) ([]byte, error) {
	return nil, nil
}

func defaultRegistrationCredentialParameters() []protocol.CredentialParameter {
	return []protocol.CredentialParameter{
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgES256,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgES384,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgES512,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgRS256,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgRS384,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgRS512,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgPS256,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgPS384,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgPS512,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgEdDSA,
		},
	}
}

// It takes a JSON string, converts it to a struct, and then converts that struct to a different struct
func parseCreationData(bz string) (*protocol.ParsedCredentialCreationData, error) {
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

// parseAssertionData takes a JSON string, converts it to a struct, and then converts that struct to a different struct
func parseAssertionData(bz string) (*protocol.ParsedCredentialAssertionData, error) {
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
