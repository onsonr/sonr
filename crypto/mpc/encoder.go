package mpc

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/onsonr/sonr/crypto/core/curves"
	"github.com/onsonr/sonr/crypto/core/protocol"
	"github.com/onsonr/sonr/crypto/keys"
)

// ╭───────────────────────────────────────────────────────────╮
// │                    Exported Generics                      │
// ╰───────────────────────────────────────────────────────────╯

type (
	Role        int                            // Role is the type for the role
	Message     *protocol.Message              // Message is the protocol.Message that is used for MPC
	Signature   *curves.EcdsaSignature         // Signature is the type for the signature
	RefreshFunc interface{ protocol.Iterator } // RefreshFunc is the type for the refresh function
	SignFunc    interface{ protocol.Iterator } // SignFunc is the type for the sign function
)

func GetIssuerDID(pk keys.PubKey) (string, string, error) {
	addr, err := bech32.ConvertAndEncode("idx", pk.Bytes())
	if err != nil {
		return "", "", err
	}
	return fmt.Sprintf("did:sonr:%s", addr), addr, nil
}

// ╭───────────────────────────────────────────────────────────╮
// │                  MPC Share Roles (Alice/Bob)              │
// ╰───────────────────────────────────────────────────────────╯

const (
	RoleUnknown Role = iota
	RoleUser
	RoleValidator
)

func (r Role) IsUser() bool      { return r == RoleUser }
func (r Role) IsValidator() bool { return r == RoleValidator }

// ╭───────────────────────────────────────────────────────────╮
// │                      Keyshare Encoding                    │
// ╰───────────────────────────────────────────────────────────╯

type KeyShare string
