package types

import (
	"errors"

	"github.com/sonrhq/core/internal/crypto"
	algo "github.com/sonrhq/core/internal/mpc/protocol/cmp"
)

var (
	// Default Account Name is the default name of the account.
	kDefaultAccName = "primary"

	// Default Threshold is the default number of required signatures to authorize a transaction.
	kDefaultThreshold = 1

	// Default Group is the default list of all parties that can sign transactions.
	kDefaultGroup = []crypto.PartyID{crypto.PartyID("vault")}
)

// KeygenOption is a function that configures an account.
type KeygenOption func(*KeygenOpts)

// WithThreshold sets the number of required signatures to authorize a transaction.
func WithThreshold(threshold int) KeygenOption {
	return func(c *KeygenOpts) {
		c.Threshold = threshold
	}
}

// WithPartyIDs sets the list of all parties that can sign transactions.
func WithPartyIDs(ids ...crypto.PartyID) KeygenOption {
	return func(c *KeygenOpts) {
		c.Peers = ids
	}
}

// KeygenOpts is the configuration of an account.
type KeygenOpts struct {
	// Network is the network that is used to communicate with other parties.
	Network algo.Network

	// Threshold is the number of required signatures to authorize a transaction.
	Threshold int

	// Group is the list of all parties that can sign transactions.
	Peers []crypto.PartyID

	current crypto.PartyID
}

func DefaultKeygenOpts(current crypto.PartyID) *KeygenOpts {
	return &KeygenOpts{
		Threshold: kDefaultThreshold,
		Peers:     kDefaultGroup,
		current:   current,
	}
}

func (o *KeygenOpts) Apply(opts ...KeygenOption) {
	for _, opt := range opts {
		opt(o)
	}
	o.Peers = algo.EnsureSelfIDInGroup(o.current, o.Peers)
}

func (o *KeygenOpts) GetNetwork() algo.Network {
	return algo.NewOfflineNetwork(o.Peers...)
}

// - The R and S values must be in the valid range for secp256k1 scalars:
//   - Negative values are rejected
//   - Zero is rejected
//   - Values greater than or equal to the secp256k1 group order are rejected
func DeserializeECDSASecp256k1Signature(sigStr []byte) (*crypto.MPCECDSASignature, error) {
	rBytes := sigStr[:33]
	sBytes := sigStr[33:65]

	sig := crypto.NewEmptyECDSASecp256k1Signature()
	if err := sig.R.UnmarshalBinary(rBytes); err != nil {
		return nil, errors.New("malformed signature: R is not in the range [1, N-1]")
	}

	// S must be in the range [1, N-1].  Notice the check for the maximum number
	// of bytes is required because SetByteSlice truncates as noted in its
	// comment so it could otherwise fail to detect the overflow.
	if err := sig.S.UnmarshalBinary(sBytes); err != nil {
		return nil, errors.New("malformed signature: S is not in the range [1, N-1]")
	}

	// Create and return the signature.
	return &sig, nil
}

// SerializeECDSASecp256k1Signature marshals an ECDSA signature to DER format for use with the CMP protocol
func SerializeECDSASecp256k1Signature(sig *crypto.MPCECDSASignature) ([]byte, error) {
	rBytes, err := sig.R.MarshalBinary()
	if err != nil {
		return nil, err
	}
	sBytes, err := sig.S.MarshalBinary()
	if err != nil {
		return nil, err
	}

	sigBytes := make([]byte, 65)
	// 0 pad the byte arrays from the left if they aren't big enough.
	copy(sigBytes[33-len(rBytes):33], rBytes)
	copy(sigBytes[65-len(sBytes):65], sBytes)
	return sigBytes, nil
}
