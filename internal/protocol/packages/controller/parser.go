package controller

import (
	"fmt"
	"strings"

	"github.com/sonrhq/core/pkg/crypto"
)

type KeyShareParseResult struct {
	DID 		  string
	CoinType     crypto.CoinType
	AccountAddress  string
	AccountName  string
	KeyShareName string
}

// ParseKeyShareDid parses a keyshare DID into its components. The DID format is:
// did:{coin_type}:{account_address}#ks-{account_name}-{keyshare_name}
func ParseKeyShareDid(name string) (*KeyShareParseResult, error) {
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
		DID: name,
		AccountAddress:  accountAddress,
		AccountName:  accountName,
		KeyShareName: keyShareName,
	}, nil
}
