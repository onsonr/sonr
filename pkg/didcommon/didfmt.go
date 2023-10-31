package didcommon

import (
	"fmt"
	"strings"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"

	"github.com/sonrhq/sonr/types/webauthn"
)

// Credential is a webauthn credential
type Credential = webauthn.Credential

// BroadcastTxResponse is a tx response
type BroadcastTxResponse = txtypes.BroadcastTxResponse

// Parse parses a  string
func Parse(did string) (Method, Identifier, error) {
	parts := strings.Split(did, ":")
	if len(parts) != 3 {
		return "", "", fmt.Errorf("invalid did: %s", did)
	}
	return Method(parts[1]), Identifier(parts[2]), nil
}

// Format formats a  string
func Format(method Method, identifier Identifier) string {
	return fmt.Sprintf("%s:%s:%s", "did", method.String(), identifier.String())
}

// Url is a  URL
type Url string

// EmptyURL is an empty  URL
var EmptyURL Url = ""

// NewUrl creates a new  URL
func NewUrl(method Method, id Identifier) Url {
	return Url(fmt.Sprintf("did:%s:%s", method, id))
}

// ParseUrl parses a  URL
func ParseUrl(did string) (Url, error) {
	m, id, err := Parse(did)
	if err != nil {
		return "", err
	}
	return NewUrl(m, id), nil
}

// Method returns the method of the  URL
func (d Url) Method() (Method, error) {
	err := d.Valid()
	if err != nil {
		return "", err
	}
	ptrs := strings.Split(string(d), ":")
	return Method(ptrs[1]), nil
}

// Identifier returns the identifier of the  URL
func (d Url) Identifier() (Identifier, error) {
	err := d.Valid()
	if err != nil {
		return "", err
	}
	ptrs := strings.Split(string(d), ":")
	return Identifier(ptrs[2]), nil
}

// String returns the string representation of the  URL
func (d Url) String() string {
	return string(d)
}

// Valid returns whether the  URL is valid
func (d Url) Valid() error {
	ptrs := strings.Split(string(d), ":")
	if len(ptrs) != 3 {
		return fmt.Errorf("invalid did url: %s. Needs minimum 3 parts", d)
	}
	if ptrs[0] != "did" {
		return fmt.Errorf("invalid did url: %s. First part must be 'did'", d)
	}
	return nil
}
