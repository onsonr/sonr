package resolver

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/sonrhq/core/pkg/crypto"
)

type KeyShareParseResult struct {
	CoinType     crypto.CoinType
	AccountAddress  string
	AccountName  string
	KeyShareName string
}

// Parse a DID into a WebauthnCredential struct
func ParseCredentialDID(did string) (*crypto.WebauthnCredential, error) {
	didParts := strings.SplitN(did, ":", 3)
	if len(didParts) != 3 {
		return nil, fmt.Errorf("invalid did: %s", did)
	}
	if didParts[0] != "did" {
		return nil, fmt.Errorf("invalid did: %s", did)
	}
	if didParts[1] != "key" {
		return nil, fmt.Errorf("invalid did: %s", did)
	}
	pubKey, err := base64.RawURLEncoding.DecodeString(didParts[2])
	if err != nil {
		return nil, err
	}

	idParts := strings.Split(did, "#")
	if len(idParts) != 2 {
		return nil, fmt.Errorf("invalid did: %s", did)
	}

	idBytes, err := base64.RawURLEncoding.DecodeString(idParts[1])
	if err != nil {
		return nil, err
	}

	return &crypto.WebauthnCredential{
		PublicKey: pubKey,
		Id:        idBytes,
	}, nil
}

// ParseKeyShareDID parses a keyshare DID into its components.
func ParseKeyShareDID(name string) (*KeyShareParseResult, error) {
	// Parse the DID
	parts := strings.Split(name, ":")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid keyshare DID: %s", name)
	}

	// Parse the coin type
	ct := crypto.CoinTypeFromDidMethod(parts[1])

	// Split the account address and keyshare name
	parts = strings.Split(parts[2], "#ks-")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid keyshare DID: %s", name)
	}

	// Parse the account address
	accountAddress := parts[0]

	// Parse the keyshare name
	parts = strings.Split(parts[1], "-")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid keyshare DID: %s", name)
	}


	// Parse the keyshare name
	accountName := parts[0]
	keyShareName := parts[1]



	return &KeyShareParseResult{
		CoinType:     ct,
		AccountAddress:  accountAddress,
		AccountName:  accountName,
		KeyShareName: keyShareName,
	}, nil
}

// ParseAccountDID parses an account DID into its components.
func ParseAccountDID(name string) (*KeyShareParseResult, error) {
	// Parse the DID
	parts := strings.Split(name, ":")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid account DID: %s", name)
	}

	// Parse the coin type
	ct := crypto.CoinTypeFromDidMethod(parts[1])

	// Parse the account address
	accountAddress := parts[2]

	return &KeyShareParseResult{
		CoinType:     ct,
		AccountAddress:  accountAddress,
	}, nil
}
