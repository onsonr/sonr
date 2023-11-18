package crypto

import (
	"crypto/x509"
	"encoding/base64"
	"strings"
)

// Base64UrlToBytes converts a base64url string to bytes.
func Base64UrlToBytes(base64Url string) ([]byte, error) {
	base64String := strings.ReplaceAll(strings.ReplaceAll(base64Url, "-", "+"), "_", "/")
	missingPadding := len(base64String) % 4
	if missingPadding > 0 {
		base64String += strings.Repeat("=", 4-missingPadding)
	}
	return base64.StdEncoding.DecodeString(base64String)
}

// ParseCredentialPublicKey parses a public key from a base64url string.
func ParseCredentialPublicKey(pubStr string) (interface{}, error) {
	derEncodedPublicKey, err := Base64UrlToBytes(pubStr)
	if err != nil {
		return nil, err
	}

	publicKey, err := x509.ParsePKIXPublicKey(derEncodedPublicKey)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}
