package context

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
	"github.com/medama-io/go-useragent"
	hwayorm "github.com/onsonr/sonr/internal/database/hwayorm"
)

// Define the credential structure matching our frontend data
type CredentialDescriptor struct {
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

func (c *CredentialDescriptor) ToModel(handle, origin string) *hwayorm.Credential {
	return &hwayorm.Credential{
		Handle:                  handle,
		Origin:                  origin,
		CredentialID:            c.ID,
		Type:                    c.Type,
		Transports:              c.Transports,
		AuthenticatorAttachment: c.AuthenticatorAttachment,
	}
}

func CredentialArrayToDescriptors(credentials []hwayorm.Credential) []*CredentialDescriptor {
	var descriptors []*CredentialDescriptor
	for _, cred := range credentials {
		cd := &CredentialDescriptor{
			ID:                      cred.CredentialID,
			RawID:                   cred.CredentialID,
			Type:                    cred.Type,
			AuthenticatorAttachment: cred.AuthenticatorAttachment,
			Transports:              cred.Transports,
		}
		descriptors = append(descriptors, cd)
	}
	return descriptors
}

func BaseSessionCreateParams(e echo.Context) hwayorm.CreateSessionParams {
	// f := rand.Intn(5) + 1
	// l := rand.Intn(4) + 1
	challenge, _ := protocol.CreateChallenge()
	id := getOrCreateSessionID(e)
	ua := useragent.NewParser()
	s := ua.Parse(e.Request().UserAgent())

	return hwayorm.CreateSessionParams{
		ID:             id,
		BrowserName:    s.GetBrowser(),
		BrowserVersion: s.GetMajorVersion(),
		ClientIpaddr:   e.RealIP(),
		Platform:       s.GetOS(),
		IsMobile:       s.IsMobile(),
		IsTablet:       s.IsTablet(),
		IsDesktop:      s.IsDesktop(),
		IsBot:          s.IsBot(),
		IsTv:           s.IsTV(),
		// IsHumanFirst:   int64(f),
		// IsHumanLast:    int64(l),
		Challenge: challenge.String(),
	}
}

// ╭───────────────────────────────────────────────────────────╮
// │            Create Passkey (/register/passkey)             │
// ╰───────────────────────────────────────────────────────────╯

// CreatePasskeyParams represents the parameters for creating a passkey
type CreatePasskeyParams struct {
	Address       string
	Handle        string
	Name          string
	Challenge     string
	CreationBlock string
}

// ╭───────────────────────────────────────────────────────────╮
// │            Create Profile (/register/profile)             │
// ╰───────────────────────────────────────────────────────────╯

// CreateProfileParams represents the parameters for creating a profile
type CreateProfileParams struct {
	TurnstileSiteKey string
	FirstNumber      int
	LastNumber       int
}

// Sum returns the sum of the first and last number
func (d CreateProfileParams) Sum() int {
	return d.FirstNumber + d.LastNumber
}
