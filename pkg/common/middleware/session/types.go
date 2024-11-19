package session

import (
	"github.com/go-webauthn/webauthn/protocol"

	"github.com/onsonr/sonr/pkg/motr/config"
)

type (
	LoginOptions    = protocol.PublicKeyCredentialRequestOptions
	RegisterOptions = protocol.PublicKeyCredentialCreationOptions
	VaultSchema     = config.Schema
)

type BrowserName string

const (
	BrowserNameUnknown  BrowserName = " Not A;Brand"
	BrowserNameChromium BrowserName = "Chromium"
)

func (n BrowserName) String() string {
	return string(n)
}

type PeerRole string

const (
	RoleUnknown PeerRole = "none"
	RoleHway    PeerRole = "hway"
	RoleMotr    PeerRole = "motr"
)

func (r PeerRole) Is(role PeerRole) bool {
	return r == role
}

func (r PeerRole) String() string {
	return string(r)
}

type ClientConfig struct {
	ChainID    string `json:"chainID"`
	IPFSHost   string `json:"ipfsHost"`
	SonrAPIURL string `json:"sonrAPIURL"`
	SonrRPCURL string `json:"sonrRPCURL"`
	SonrWSURL  string `json:"sonrWSURL"`
}

type PeerSession struct {
	ID        string                    `json:"id"`
	Challenge protocol.URLEncodedBase64 `json:"challenge"`
}

type PlatformInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type BrowserInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type DeviceInfo struct {
	Architecture string        `json:"architecture"`
	Bitness      string        `json:"bitness"`
	Model        string        `json:"model"`
	Platform     *PlatformInfo `json:"platform"`
}

type UserAgent struct {
	Browser  *BrowserInfo `json:"browser"`
	Device   *DeviceInfo  `json:"device"`
	IsMobile bool         `json:"isMobile"`
}

type VaultConfig struct {
	Schema  *VaultSchema `json:"schema"`
	Address string       `json:"address"`
}
