package resolver

import (
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

// parseKeyShareDid parses a keyshare DID into its components.
func parseKeyShareDid(name string) (*KeyShareParseResult, error) {
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

// parseAccountDid parses an account DID into its components.
func parseAccountDid(name string) (*KeyShareParseResult, error) {
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
