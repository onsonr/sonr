package common

import (
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/mr-tron/base58"
)

// Parse parses a DID string into a DID struct
func ParseDid(s string) (*Did, error) {
	var did Did

	if !IsValidDid(Method, s) {
		return nil, ErrParseInvalid
	}
	// methodDid, methodFragment := SplitDidUrlIntoDidAndFragment(methodId)
	// if len(methodDid) == 0 {
	// 	result = did + "#" + methodFragment
	// }
	// Parse Items from string into DID struct
	did.Method = Method

	return &did, nil
}

// CreateBaseDID creates a base DID with a given users libp2p public key
func CreateBaseDID(pubKey crypto.PubKey) (*Did, string, error) {
	// Marshal the public key into bytes
	pubBuf, err := crypto.MarshalPublicKey(pubKey)
	if err != nil {
		return nil, "", err
	}

	pubStr := base58.Encode(pubBuf)

	// Encode the public key into base58 and return the result
	did, err := NewDID(pubStr)
	if err != nil {
		return nil, "", err
	}
	return did, pubStr, nil
}

// CreateServiceDID creates a service DID with a developers libp2p public key
func CreateServiceDID(pubKey crypto.PubKey, name string) (*Did, string, error) {
	// Marshal the public key into bytes
	pubBuf, err := crypto.MarshalPublicKey(pubKey)
	if err != nil {
		return nil, "", err
	}

	pubStr := base58.Encode(pubBuf)

	// Encode the public key into base58 and return the result
	did, err := NewDID(pubStr, WithPath(name))
	if err != nil {
		return nil, "", err
	}
	return did, pubStr, nil
}

// NewDID creates a new DID
func NewDID(id string, options ...Option) (*Did, error) {
	var did Did
	did.Id = id
	did.Method = Method

	// Apply options
	for _, option := range options {
		option(&did)
	}

	// Check if the DID is valid
	if !did.IsValid() {
		return nil, ErrFragmentAndQuery
	}
	return &did, nil
}

// GetBase returns the base DID string: Method + Network
func (d *Did) GetBase() string {
	if d.HasNetwork() {
		return d.Method + ":" + d.Network + ":"
	}
	return d.Method + ":"
}

// HasNetwork returns true if the DID has a network
func (d *Did) HasNetwork() bool {
	return len(d.Network) > 0
}

// HasPath returns true if the DID has a path
func (d *Did) HasPath() bool {
	return len(d.Paths) > 0
}

// HasQuery returns true if the DID has a query
func (d *Did) HasQuery() bool {
	return len(d.Query) > 0
}

// HasFragment returns true if the DID has a fragment
func (d *Did) HasFragment() bool {
	return len(d.Fragment) > 0
}

// IsValid checks if a DID is valid and does not contain both a Fragment and a Query
func (d *Did) IsValid() bool {
	hq := d.HasQuery()
	hf := d.HasFragment()

	if hq && hf {
		return false
	}
	return true
}

// String combines all DID parts into a string
func (d *Did) ToString() string {
	return d.GetBase() + d.Id + d.GetQuery() + d.GetFragment() + d.GetFragment()
}
