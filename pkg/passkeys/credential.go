package passkeys

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
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

func ExtractCredential(jsonString string) (*Credential, error) {
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
		cred.ID,
		cred.RawID,
		cred.Type,
		cred.AuthenticatorAttachment,
		cred.Transports,
	)
	return cred, nil
}
