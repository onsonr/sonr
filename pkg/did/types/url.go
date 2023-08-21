package types

import (
	"fmt"
	"strings"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	servicetypes "github.com/sonr-io/sonr/x/service/types"
)

// Credential is a webauthn credential
type Credential = servicetypes.WebauthnCredential

// BroadcastTxResponse is a tx response
type BroadcastTxResponse = txtypes.BroadcastTxResponse

// ParseDID parses a DID string
func ParseDID(did string) (DIDMethod, DIDIdentifier, error) {
	parts := strings.Split(did, ":")
	if len(parts) != 3 {
		return "", "", fmt.Errorf("invalid did: %s", did)
	}
	return DIDMethod(parts[1]), DIDIdentifier(parts[2]), nil
}

// FormatDID formats a DID string
func FormatDID(method DIDMethod, identifier DIDIdentifier) string {
	return fmt.Sprintf("%s:%s:%s", "did", method.String(), identifier.String())
}

// DIDUrl is a DID URL
type DIDUrl string

// EmptyDIDURL is an empty DID URL
var EmptyDIDURL DIDUrl = ""

// NewDIDUrl creates a new DID URL
func NewDIDUrl(method DIDMethod, id DIDIdentifier) DIDUrl {
	return DIDUrl(fmt.Sprintf("did:%s:%s", method, id))
}

// ParseDIDUrl parses a DID URL
func ParseDIDUrl(did string) (DIDUrl, error) {
	m, id, err := ParseDID(did)
	if err != nil {
		return "", err
	}
	return NewDIDUrl(m, id), nil
}

// Method returns the method of the DID URL
func (d DIDUrl) Method() (DIDMethod, error) {
	err := d.Valid()
	if err != nil {
		return "", err
	}
	ptrs := strings.Split(string(d), ":")
	return DIDMethod(ptrs[1]), nil
}

// Identifier returns the identifier of the DID URL
func (d DIDUrl) Identifier() (DIDIdentifier, error) {
	err := d.Valid()
	if err != nil {
		return "", err
	}
	ptrs := strings.Split(string(d), ":")
	return DIDIdentifier(ptrs[2]), nil
}

// String returns the string representation of the DID URL
func (d DIDUrl) String() string {
	return string(d)
}

// Valid returns whether the DID URL is valid
func (d DIDUrl) Valid() error {
	ptrs := strings.Split(string(d), ":")
	if len(ptrs) != 3 {
		return fmt.Errorf("invalid did url: %s. Needs minimum 3 parts", d)
	}
	if ptrs[0] != "did" {
		return fmt.Errorf("invalid did url: %s. First part must be 'did'", d)
	}
	return nil
}
