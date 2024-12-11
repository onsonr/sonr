package register

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/rand"
)

// Define the credential structure matching our frontend data
type Credential struct {
	Handle                  string            `json:"handle"`
	ID                      string            `json:"id"`
	RawID                   string            `json:"rawId"`
	Type                    string            `json:"type"`
	AuthenticatorAttachment string            `json:"authenticatorAttachment"`
	Transports              string            `json:"transports"`
	ClientExtensionResults  map[string]string `json:"clientExtensionResults"`
	Response                struct {
		AttestationObject string `json:"attestationObject"`
		ClientDataJSON    string `json:"clientDataJSON"`
	} `json:"response"`
}

type CreateProfileData struct {
	TurnstileSiteKey string
	FirstNumber      int
	LastNumber       int
}

type RegisterPasskeyData struct {
	Address       string
	Handle        string
	Name          string
	Challenge     string
	CreationBlock string
}

func (d CreateProfileData) IsHumanLabel() string {
	return fmt.Sprintf("What is %d + %d?", d.FirstNumber, d.LastNumber)
}

func extractCredentialDescriptor(jsonString string) (*Credential, error) {
	cred := &Credential{}
	// Unmarshal the credential JSON
	if err := json.Unmarshal([]byte(jsonString), cred); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid credential format: %v", err))
	}

	// Validate required fields
	if cred.ID == "" || cred.RawID == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "missing credential ID")
	}
	if cred.Type != "public-key" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid credential type")
	}
	if cred.Response.AttestationObject == "" || cred.Response.ClientDataJSON == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "missing attestation data")
	}

	// Log detailed credential information
	fmt.Printf("Credential Details:\n"+
		"ID: %s\n"+
		"Raw ID: %s\n"+
		"Type: %s\n"+
		"Authenticator Attachment: %s\n"+
		"Transports: %v\n"+
		"Attestation Object Size: %d bytes\n"+
		"Client Data Size: %d bytes\n",
		cred.ID,
		cred.RawID,
		cred.Type,
		cred.AuthenticatorAttachment,
		cred.Transports,
	)
	return cred, nil
}

func randomCreateProfileData() CreateProfileData {
	return CreateProfileData{
		FirstNumber: rand.Intn(5) + 1,
		LastNumber:  rand.Intn(4) + 1,
	}
}
