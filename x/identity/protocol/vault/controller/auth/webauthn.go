package session

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/go-webauthn/webauthn/protocol"
)

// // BeginRegistration starts the registration process for the underlying Webauthn instance
// func (s *Session) GetChallengeResponse() (*v1.ChallengeResponse, error) {
// 	// Fetch Session Data
// 	opts, sessionData, err := s.webauthn.BeginRegistration(s.didDoc, webauthn.WithAuthenticatorSelection(defaultAuthSelect))
// 	if err != nil {
// 		return nil, err
// 	}
// 	s.data = sessionData
// 	bz, err := json.Marshal(opts)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &v1.ChallengeResponse{
// 		CreationOptions: string(bz),
// 		RpName:          s.AKA,
// 		SessionId:       s.AKA,
// 	}, nil
// }

// // RegisterCredential creates a credential which can be stored to use with User Authentication
// func (s *Session) RegisterCredential(credentialCreationData string) (*v1.RegisterResponse, error) {
// 	// Parse Client Credential Data
// 	pcc, err := parseCreationData(credentialCreationData)
// 	if err != nil {
// 		return nil, errors.New(fmt.Sprintf("Failed to get parsed creation data: %s", err))
// 	}
// 	cred, err := webauthn.MakeNewCredential(pcc)
// 	if err != nil {
// 		return nil, fmt.Errorf("Failed to make new credential: %s", err)
// 	}
// 	label := fmt.Sprintf("webauthn-%v", s.didDoc.AuthenticationCount()+1)
// 	pub := types.NewPubKey(cred.PublicKey, types.KeyType_KeyType_WEB_AUTHN_AUTHENTICATION_2018)
// 	vm, err := pub.VerificationMethod(types.WithIDFragmentSuffix(label))
// 	if err != nil {
// 		return nil, fmt.Errorf("Failed to add webauthn credential: %s", err)
// 	}
// 	s.didDoc.AddAuthentication(vm)
// 	return &v1.RegisterResponse{
// 		Success:     true,
// 		DidDocument: s.didDoc,
// 		Username:    s.AKA,
// 	}, nil
// }

// // GetAssertionOptions creates a new AssertionChallenge for client to verify
// func (s *Session) GetAssertionOptions() (*v1.AssertResponse, error) {
// 	opts, session, err := s.webauthn.BeginLogin(s.didDoc, webauthn.WithAllowedCredentials(s.didDoc.AllowedWebauthnCredentials()))
// 	if err != nil {
// 		return nil, err
// 	}
// 	s.data = session
// 	bz, err := json.Marshal(opts)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &v1.AssertResponse{
// 		RequestOptions: string(bz),
// 		SessionId:      s.AKA,
// 		RpName:         s.AKA,
// 	}, nil
// }

// // AuthorizeCredential authenticates from the signature provided to the client
// func (s *Session) AuthorizeCredential(credentialRequestData string) (*v1.LoginResponse, error) {
// 	pca, err := parseAssertionData(credentialRequestData)
// 	if err != nil {
// 		return nil, errors.New(fmt.Sprintf("Failed to get parsed creation data: %s", err))
// 	}
// 	_, err = s.webauthn.ValidateLogin(s.didDoc, *s.data, pca)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &v1.LoginResponse{
// 		Success:     true,
// 		DidDocument: s.didDoc,
// 		Username:    s.AKA,
// 	}, nil
// }

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
