package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sonrhq/core/pkg/crypto"
)

// TxResponse is a type alias for sdk.TxResponse
type TxResponse = sdk.TxResponse

// DIDResource is a byte array that is stored under a DID identifier in its internal store.
type DIDResource string

// NewResource returns a new DID resource from a DID URL, and key
func NewResource(id DIDIdentifier, key string) DIDResource {
	return DIDResource(fmt.Sprintf("%s@%s", id.String(), key))
}

// Key returns the key of the resource
func (d DIDResource) Key() string {
	ptrs := strings.Split(string(d), "@")
	if len(ptrs) < 2 {
		return ""
	}
	return ptrs[1]
}

// Data returns the data stored in the resource
func (d DIDResource) Data() ([]byte, error) {
	datStr := d.Identifier().GetKey(d.Key())
	if datStr == "" {
		return nil, fmt.Errorf("resource not found")
	}
	return decodeResource(datStr), nil
}

// Identifier returns the identifier of the resource
func (d DIDResource) Identifier() DIDIdentifier {
	ptrs := strings.Split(string(d), "@")
	if len(ptrs) < 2 {
		return ""
	}
	return DIDIdentifier(ptrs[0])
}

// String returns the string representation of the resource
func (d DIDResource) String() string {
	return string(d)
}

// Update replaces the data stored in the resource
func (d DIDResource) Update(data []byte) error {
	if !d.Identifier().HasKey(d.Key()) {
		return fmt.Errorf("resource not found")
	}
	d.Identifier().SetKey(d.Key(), encodeResource(data))
	return nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                Utility functions                               ||
// ! ||--------------------------------------------------------------------------------||

func encodeResource(data []byte) string {
	return crypto.Base64Encode(data)
}

func decodeResource(data string) []byte {
	bz, err := crypto.Base64Decode(data)
	if err != nil {
		panic(err)
	}
	return bz
}
