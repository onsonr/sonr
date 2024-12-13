package mpc

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/onsonr/sonr/crypto/core/curves"
	"github.com/onsonr/sonr/crypto/core/protocol"
	"github.com/onsonr/sonr/crypto/keys"
)

type Role int

const (
	RoleUnknown Role = iota
	RoleUser
	RoleValidator
)

func (r Role) IsUser() bool      { return r == RoleUser }
func (r Role) IsValidator() bool { return r == RoleValidator }

// Message is the protocol.Message that is used for MPC
type (
	Message   *protocol.Message
	Signature *curves.EcdsaSignature
)

// RefreshFunc is the type for the refresh function
type RefreshFunc interface{ protocol.Iterator }

// SignFunc is the type for the sign function
type SignFunc interface{ protocol.Iterator }

func GetIssuerDID(pk keys.PubKey) (string, string, error) {
	addr, err := bech32.ConvertAndEncode("idx", pk.Bytes())
	if err != nil {
		return "", "", err
	}
	return fmt.Sprintf("did:sonr:%s", addr), addr, nil
}
