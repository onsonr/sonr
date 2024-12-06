package common

import (
	"encoding/base64"
)

type Payment struct {
	IsPayment bool `json:"isPayment"`
}

type LargeBlob struct {
	Support string `json:"support"`
	Write   string `json:"write"`
}

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

func Base64Encode(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

func Base64Decode(data string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(data)
}
