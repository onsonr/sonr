package didcommon

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sonrhq/sonr/internal/crypto"
)

// TxResponse is a type alias for sdk.TxResponse
type TxResponse = sdk.TxResponse

// DIDResource is a byte array that is stored under a DID identifier in its internal store.
type DIDResource string

// NewResource returns a new DID resource from a DID URL, and key
func NewResource(id string, key string) DIDResource {
	return DIDResource(fmt.Sprintf("%s@%s", id, key))
}

// Key returns the key of the resource
func (d DIDResource) Key() string {
	ptrs := strings.Split(string(d), "@")
	if len(ptrs) < 2 {
		return ""
	}
	return ptrs[1]
}

// String returns the string representation of the resource
func (d DIDResource) String() string {
	return string(d)
}

func EncodeResource(data []byte) string {
	return crypto.Base64Encode(data)
}

func DecodeResource(data string) []byte {
	bz, err := crypto.Base64Decode(data)
	if err != nil {
		panic(err)
	}
	return bz
}
