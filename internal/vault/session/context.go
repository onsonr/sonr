package session

import (
	"regexp"
	"strings"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"

	"github.com/onsonr/sonr/pkg/common"
)

const kWebAuthnTimeout = 6000

// TODO: Returns fixed chain ID for testing.
func GetChainID(c echo.Context) string {
	return "sonr-testnet-1"
}

// SetVaultAddress sets the address of the vault
func SetVaultAddress(c echo.Context, address string) error {
	return common.WriteCookie(c, common.SonrAddress, address)
}

// SetVaultAuthorization sets the UCAN CID of the vault
func SetVaultAuthorization(c echo.Context, ucanCID string) error {
	common.HeaderWrite(c, common.Authorization, formatAuth(ucanCID))
	return nil
}

// ╭───────────────────────────────────────────────────────────╮
// │                       Initialization                      │
// ╰───────────────────────────────────────────────────────────╯

func loadOrGenKsuid(c echo.Context) error {
	var (
		sessionID string
		err       error
	)

	// Setup genKsuid function
	genKsuid := func() string {
		return ksuid.New().String()
	}

	// Attempt to read the session ID from the "session" cookie
	if ok := common.CookieExists(c, common.SessionID); !ok {
		sessionID = genKsuid()
	} else {
		sessionID, err = common.ReadCookie(c, common.SessionID)
		if err != nil {
			sessionID = genKsuid()
		}
	}
	common.WriteCookie(c, common.SessionID, sessionID)
	return nil
}

// ╭───────────────────────────────────────────────────────────╮
// │                       Extraction                          │
// ╰───────────────────────────────────────────────────────────╯

func extractPeerInfo(c echo.Context) (string, string) {
	var chal protocol.URLEncodedBase64
	id, _ := common.ReadCookie(c, common.SessionID)
	chalRaw, _ := common.ReadCookieBytes(c, common.SessionChallenge)
	chal.UnmarshalJSON(chalRaw)

	return id, common.Base64Encode(chal)
}

func extractBrowserInfo(c echo.Context) (string, string) {
	secCHUA := common.HeaderRead(c, common.UserAgent)

	// If common.is empty, return empty BrowserInfo
	if secCHUA == "" {
		return "N/A", "-1"
	}

	// Split the common.into individual browser entries
	var (
		name string
		ver  string
	)
	entries := strings.Split(strings.TrimSpace(secCHUA), ",")
	for _, entry := range entries {
		// Remove leading/trailing spaces and quotes
		entry = strings.TrimSpace(entry)

		// Use regex to extract the browser name and version
		re := regexp.MustCompile(`"([^"]+)";v="([^"]+)"`)
		matches := re.FindStringSubmatch(entry)

		if len(matches) == 3 {
			browserName := matches[1]
			version := matches[2]

			// Skip "Not A;Brand"
			if !validBrowser(browserName) {
				continue
			}

			// Store the first valid browser info as fallback
			name = browserName
			ver = version
		}
	}
	return name, ver
}

func validBrowser(name string) bool {
	return name != common.BrowserNameUnknown.String() && name != common.BrowserNameChromium.String()
}

// ╭───────────────────────────────────────────────────────────╮
// │                        Authentication                     │
// ╰───────────────────────────────────────────────────────────╯

func buildUserEntity(userID string) protocol.UserEntity {
	return protocol.UserEntity{
		ID: userID,
	}
}

// returns the base options for registering a new user without challenge or user entity.
func baseRegisterOptions() *protocol.PublicKeyCredentialCreationOptions {
	return &protocol.PublicKeyCredentialCreationOptions{
		Timeout:     kWebAuthnTimeout,
		Attestation: protocol.PreferDirectAttestation,
		AuthenticatorSelection: protocol.AuthenticatorSelection{
			AuthenticatorAttachment: "platform",
			ResidentKey:             protocol.ResidentKeyRequirementPreferred,
			UserVerification:        "preferred",
		},
		Parameters: []protocol.CredentialParameter{
			{
				Type:      "public-key",
				Algorithm: webauthncose.AlgES256,
			},
			{
				Type:      "public-key",
				Algorithm: webauthncose.AlgES256K,
			},
			{
				Type:      "public-key",
				Algorithm: webauthncose.AlgEdDSA,
			},
		},
	}
}

func formatAuth(ucanCID string) string {
	return "Bearer " + ucanCID
}
