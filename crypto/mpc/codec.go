package mpc

import (
	"fmt"
	"strings"

	"github.com/onsonr/sonr/crypto/core/curves"
	"github.com/onsonr/sonr/crypto/core/protocol"
	"github.com/onsonr/sonr/crypto/tecdsa/dklsv1/dkg"
)

// ╭───────────────────────────────────────────────────────────╮
// │                    Exported Generics                      │
// ╰───────────────────────────────────────────────────────────╯

type (
	AliceOut    *dkg.AliceOutput
	BobOut      *dkg.BobOutput
	Point       curves.Point
	Role        string                         // Role is the type for the role
	Message     *protocol.Message              // Message is the protocol.Message that is used for MPC
	Signature   *curves.EcdsaSignature         // Signature is the type for the signature
	RefreshFunc interface{ protocol.Iterator } // RefreshFunc is the type for the refresh function
	SignFunc    interface{ protocol.Iterator } // SignFunc is the type for the sign function
)

// ╭───────────────────────────────────────────────────────────╮
// │                      Keyshare Encoding                    │
// ╰───────────────────────────────────────────────────────────╯

type KeyShare string

func EncodeKeyshare(m Message, role Role) (KeyShare, error) {
	enc, err := protocol.EncodeMessage(m)
	if err != nil {
		return "", err
	}
	return KeyShare(fmt.Sprintf("%s.%s", role.String(), enc)), nil
}

func DecodeKeyshare(s string) (KeyShare, error) {
	parts := strings.Split(s, ".")
	role := Role(parts[0])
	if role != RoleUser && role != RoleValidator {
		return "", fmt.Errorf("invalid share role")
	}
	_, err := protocol.DecodeMessage(parts[1])
	if err != nil {
		return "", err
	}
	return KeyShare(fmt.Sprintf("%s.%s", role.String(), parts[1])), nil
}

func (k KeyShare) AliceOut() (AliceOut, error) {
	if k.Role() != RoleValidator {
		return nil, fmt.Errorf("invalid share role")
	}
	msg, err := k.Message()
	if err != nil {
		return nil, err
	}
	return getAliceOut(msg)
}

func (k KeyShare) BobOut() (BobOut, error) {
	if k.Role() != RoleUser {
		return nil, fmt.Errorf("invalid share role")
	}
	msg, err := k.Message()
	if err != nil {
		return nil, err
	}
	return getBobOut(msg)
}

func (k KeyShare) Bytes() []byte {
	return []byte(k)
}

func (k KeyShare) Message() (*protocol.Message, error) {
	parts := strings.Split(k.String(), ".")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid share format")
	}
	return protocol.DecodeMessage(parts[1])
}

func (k KeyShare) Role() Role {
	parts := strings.Split(k.String(), ".")
	return Role(parts[0])
}

func (k KeyShare) String() string {
	return string(k)
}

// ╭───────────────────────────────────────────────────────────╮
// │                  MPC Share Roles (Alice/Bob)              │
// ╰───────────────────────────────────────────────────────────╯

const (
	RoleUser      Role = "user"
	RoleValidator Role = "validator"
	RoleUnknown   Role = "nahh"
)

func (r Role) String() string {
	return string(r)
}

func (r Role) IsUser() bool      { return r == RoleUser }
func (r Role) IsValidator() bool { return r == RoleValidator }
