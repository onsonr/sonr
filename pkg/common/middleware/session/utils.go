package session

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"

	"github.com/onsonr/sonr/pkg/common"
	"github.com/onsonr/sonr/pkg/common/middleware/cookie"
	"github.com/onsonr/sonr/pkg/common/middleware/header"
	"github.com/onsonr/sonr/pkg/common/types"
	"github.com/onsonr/sonr/pkg/core/dwn"
)

const kWebAuthnTimeout = 6000

// ╭───────────────────────────────────────────────────────────╮
// │                       Initialization                      │
// ╰───────────────────────────────────────────────────────────╯

func loadOrGenChallenge(c echo.Context) error {
	var (
		chal    protocol.URLEncodedBase64
		chalRaw []byte
		err     error
	)

	// Setup genChal function
	genChal := func() []byte {
		ch, _ := protocol.CreateChallenge()
		bz, _ := ch.MarshalJSON()
		return bz
	}

	// Check if there is a session challenge cookie
	if !cookie.Exists(c, cookie.SessionChallenge) {
		chalRaw = genChal()
		cookie.WriteBytes(c, cookie.SessionChallenge, chalRaw)
	} else {
		chalRaw, err = cookie.ReadBytes(c, cookie.SessionChallenge)
		if err != nil {
			return err
		}
	}

	// Attempt to read the session challenge from the "session" cookie
	err = chal.UnmarshalJSON(chalRaw)
	if err != nil {
		return err
	}
	return nil
}

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
	sessionID, err = cookie.Read(c, cookie.SessionID)
	if err != nil {
		sessionID = genKsuid()
		cookie.Write(c, cookie.SessionID, sessionID)
	}
	return nil
}

// ╭───────────────────────────────────────────────────────────╮
// │                       Extraction                          │
// ╰───────────────────────────────────────────────────────────╯

func extractConfigClient(c echo.Context) *types.ClientConfig {
	return &types.ClientConfig{
		ChainID:    header.Read(c, header.ChainID),
		IpfsHost:   header.Read(c, header.IPFSHost),
		SonrAPIURL: header.Read(c, header.SonrAPIURL),
		SonrRPCURL: header.Read(c, header.SonrRPCURL),
		SonrWSURL:  header.Read(c, header.SonrWSURL),
	}
}

func extractConfigVault(c echo.Context) (*types.VaultDetails, error) {
	schema := &dwn.Schema{}
	schemaBz, _ := cookie.ReadBytes(c, cookie.VaultSchema)
	err := json.Unmarshal(schemaBz, schema)
	if err != nil {
		return nil, err
	}
	addr, err := cookie.Read(c, cookie.SonrAddress)
	if err != nil {
		return nil, err
	}
	return &types.VaultDetails{
		Schema:  schema,
		Address: addr,
	}, nil
}

func extractPeerRole(c echo.Context) common.PeerRole {
	r, _ := cookie.Read(c, cookie.SessionRole)
	return common.PeerRole(r)
}

func extractPeerInfo(c echo.Context) *types.PeerInfo {
	var chal protocol.URLEncodedBase64

	id, _ := cookie.Read(c, cookie.SessionID)
	chalRaw, _ := cookie.ReadBytes(c, cookie.SessionChallenge)
	chal.UnmarshalJSON(chalRaw)

	return &types.PeerInfo{Id: id, Challenge: common.Base64Encode(chal)}
}

func extractBrowserInfo(c echo.Context) *types.BrowserInfo {
	secCHUA := header.Read(c, header.UserAgent)

	// If header is empty, return empty BrowserInfo
	if secCHUA == "" {
		return unknownBrowser()
	}

	// Split the header into individual browser entries
	var selectedBrowser *types.BrowserInfo
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
			selectedBrowser = newBrowserInfo(browserName, version)
		}
	}
	return selectedBrowser
}

func extractUserAgent(c echo.Context) *types.UserAgent {
	ua := &types.UserAgent{
		Browser:         extractBrowserInfo(c),
		Architecture:    header.Read(c, header.Architecture),
		Bitness:         header.Read(c, header.Bitness),
		Model:           header.Read(c, header.Model),
		PlatformName:    header.Read(c, header.Platform),
		PlatformVersion: header.Read(c, header.PlatformVersion),
		IsMobile:        header.Equals(c, header.Mobile, "?1"),
	}
	return ua
}

func newBrowserInfo(name string, version string) *types.BrowserInfo {
	return &types.BrowserInfo{
		Name:    name,
		Version: version,
	}
}

func unknownBrowser() *types.BrowserInfo {
	return &types.BrowserInfo{
		Name:    "Unknown",
		Version: "-1",
	}
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
func baseRegisterOptions() *common.RegisterOptions {
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
