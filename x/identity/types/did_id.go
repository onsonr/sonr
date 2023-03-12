package types

import (
	"fmt"
	"strings"

	"github.com/sonrhq/core/pkg/crypto"
)

type DIDParseResult struct {
	AccountName string
	Address     string
	CoinType    crypto.CoinType
}

// NewSonrID creates a new DID URI for the given Sonr Account address
func NewSonrID(addr string) string {
	return fmt.Sprintf("did:sonr:%s", addr)
}

// NewWebID creates a new DID URI for the given Sonr Account address
func NewWebID(addr string) string {
	return DIDMethod_DIDMethod_WEB.Format(addr)
}

// NewKeyID creates a new DID URI for the given Sonr Account address
func NewKeyID(addr string, keyName string) string {
	return DIDMethod_DIDMethod_KEY.Format(addr, WithFragment(keyName))
}

// NewIpfsID creates a new DID URI for the given Content ID
func NewIpfsID(addr string) string {
	return DIDMethod_DIDMethod_IPFS.Format(addr)
}

// NewPeerID creates a new DID URI for the given Peer ID
func NewPeerID(addr string) string {
	return DIDMethod_DIDMethod_PEER.Format(addr)
}

// NewBlockchainID creates a new DID URI for the given blockchain account address
func NewBlockchainID(addr string, name string) string {
	return DIDMethod_DIDMethod_BLOCKCHAIN.Format(addr, WithFragment(name))
}

//
// Helper Functions
//

// findCoinTypeFromAddress returns the CoinType for the given address
func findCoinTypeFromAddress(addr string) crypto.CoinType {
	for _, ct := range crypto.AllCoinTypes() {
		if strings.Contains(addr, ct.AddrPrefix()) {
			return ct
		}
	}
	return crypto.TestCoinType
}
