package handler

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/sonrhq/core/pkg/client/chain"
	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/x/identity/types"
	v1 "github.com/sonrhq/core/x/identity/types/vault/v1"
)

type ServiceHandler interface {
	// This method is used to get the challenge response from the DID controller.
	BeginRegistration(req *v1.RegisterStartRequest) ([]byte, error)

	// This is the method that will be called when the user clicks on the "Register" button.
	FinishRegistration(req *v1.RegisterFinishRequest) (bool, error)

	// This method is used to get the options for the assertion.
	BeginLogin(req *v1.LoginStartRequest) ([]byte, error)

	// This is the method that will be called when the user clicks the "Login" button on the login page.
	FinishLogin(req *v1.LoginFinishRequest) (bool, error)
}

type serviceHandlerImpl struct {
	// Service is the service that will be used to handle the requests.
	service *types.Service

	// Client is the client that will be used to interact with the chain.
	sonrQueryClient *chain.SonrQueryClient
}

func NewServiceHandler(origin string, apiEndpoint chain.APIEndpoint) (ServiceHandler, error) {
	// Get the service from the chain.
	sonrQueryClient := chain.NewClient(apiEndpoint)
	service, err := sonrQueryClient.GetService(context.Background(), origin)
	if err != nil {
		return nil, fmt.Errorf("failed to get service from chain: %w", err)
	}

	return &serviceHandlerImpl{
		service:         service,
		sonrQueryClient: sonrQueryClient,
	}, nil
}

// BeginRegistration is the method that will be called when the user clicks on the "Register" button.
func (s *serviceHandlerImpl) BeginRegistration(req *v1.RegisterStartRequest) ([]byte, error) {
	// Get the parameters from the chain.
	params := types.NewParams()

	// Issue the challenge.
	chal, err := s.service.IssueChallenge()
	if err != nil {
		return nil, fmt.Errorf("failed to issue challenge: %w", err)
	}

	// Build the credential creation options.
	creationOptions := protocol.PublicKeyCredentialCreationOptions{
		// Generated Challenge.
		Challenge: chal,

		// Service resulting properties.
		RelyingParty: s.service.RelyingPartyEntity(),
		User:         s.service.GetUserEntity(req.Uuid, req.DeviceLabel),

		// Preconfigured parameters.
		Parameters:             params.WebauthnRegistrationCredentialParameters(),
		Timeout:                params.WebauthnTimeoutInteger(),
		AuthenticatorSelection: params.WebauthnAuthenticatorSelection(),
		Attestation:            params.WebauthnConveyancePreference(),
	}

	// Marshal the response into JSON.
	response := protocol.CredentialCreation{Response: creationOptions}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}
	return jsonResponse, nil
}

// FinishRegistration is the method that will be called when the user clicks on the "Register" button.
func (s *serviceHandlerImpl) FinishRegistration(req *v1.RegisterFinishRequest) (bool, error) {
	// Get the parameters from the chain.
	pccd, err := parseCreationData(req.CredentialResponse)
	if err != nil {
		return false, fmt.Errorf("failed to parse credential response: %w", err)
	}

	// Verify the challenge.
	err = s.service.VerifyChallenge(pccd)
	if err != nil {
		return false, fmt.Errorf("failed to verify challenge: %w", err)
	}

	cred := crypto.NewWebAuthnCredential(pccd)
	pk := crypto.NewWebAuthnPubKey(cred.PublicKey)
	fmt.Println(pk)
	return true, nil
	// vm, err := types.NewVMFromPubKey(pk, types.WithController(d.primaryAccount.DID()), types.WithIDFragmentSuffix(aka))
	// if err != nil {
	// 	return false, err
	// }
}

// BeginLogin is the method that will be called when the user clicks on the "Login" button.
func (s *serviceHandlerImpl) BeginLogin(req *v1.LoginStartRequest) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

// FinishLogin is the method that will be called when the user clicks on the "Login" button.
func (s *serviceHandlerImpl) FinishLogin(req *v1.LoginFinishRequest) (bool, error) {
	return false, fmt.Errorf("not implemented")
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
