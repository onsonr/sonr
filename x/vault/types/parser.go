package types

import (
	"fmt"
	"strings"

	"github.com/sonrhq/core/internal/crypto"
)

type KeyShareParseResult struct {
	CoinType       crypto.CoinType
	AccountAddress string
	KeyShareName   string
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
	keyShareName := parts[1]

	return &KeyShareParseResult{
		CoinType:       ct,
		AccountAddress: accountAddress,
		KeyShareName:   keyShareName,
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
		CoinType:       ct,
		AccountAddress: accountAddress,
	}, nil
}

// FormatBlockchainAddressAsDID formats a blockchain address as a DID.
func FormatBlockchainAddressAsDID(addr string) string {
	ct := crypto.CoinTypeFromAddrPrefix(addr)
	return fmt.Sprintf("did:%s:%s", ct.DidMethod(), addr)
}
