package mpc

import (
	"errors"

	"github.com/onsonr/sonr/crypto/core/curves"
	"github.com/onsonr/sonr/crypto/core/protocol"
	"github.com/onsonr/sonr/crypto/tecdsa/dklsv1"
	"golang.org/x/crypto/sha3"
)

var ErrInvalidKeyshareRole = errors.New("invalid keyshare role")

type Role int

const (
	RoleUnknown Role = iota
	RoleUser
	RoleValidator
)

func (r Role) IsUser() bool {
	return r == RoleUser
}

func (r Role) IsValidator() bool {
	return r == RoleValidator
}

// Message is the protocol.Message that is used for MPC
type Message *protocol.Message

type Signature *curves.EcdsaSignature

// RefreshFunc is the type for the refresh function
type RefreshFunc interface {
	protocol.Iterator
}

// SignFunc is the type for the sign function
type SignFunc interface {
	protocol.Iterator
}

type ValKeyshare struct {
	BaseKeyshare
	encoded string
}

func NewValKeyshare(msg *protocol.Message) (*ValKeyshare, error) {
	valShare, err := dklsv1.DecodeAliceDkgResult(msg)
	if err != nil {
		return nil, err
	}
	return &ValKeyshare{
		BaseKeyshare: BaseKeyshare{
			Message:   msg,
			Role:      1,
			PublicKey: valShare.PublicKey.ToAffineUncompressed(),
		},
		encoded: "",
	}, nil
}

func (v *ValKeyshare) RefreshFunc() (RefreshFunc, error) {
	curve := curves.K256()
	return dklsv1.NewAliceRefresh(curve, v.ExtractMessage(), protocol.Version1)
}

func (v *ValKeyshare) SignFunc(msg []byte) (SignFunc, error) {
	curve := curves.K256()
	return dklsv1.NewAliceSign(curve, sha3.New256(), msg, v.ExtractMessage(), protocol.Version1)
}

func (v *ValKeyshare) String() string {
	return v.encoded
}

type UserKeyshare struct {
	BaseKeyshare
	encoded string
}

func NewUserKeyshare(msg *protocol.Message) (*UserKeyshare, error) {
	out, err := dklsv1.DecodeBobDkgResult(msg)
	if err != nil {
		return nil, err
	}
	return &UserKeyshare{
		BaseKeyshare: BaseKeyshare{
			Message:   msg,
			Role:      2,
			PublicKey: out.PublicKey.ToAffineUncompressed(),
		},
		encoded: "",
	}, nil
}

func (u *UserKeyshare) RefreshFunc() (RefreshFunc, error) {
	curve := curves.K256()
	return dklsv1.NewBobRefresh(curve, u.ExtractMessage(), protocol.Version1)
}

func (u *UserKeyshare) SignFunc(msg []byte) (SignFunc, error) {
	curve := curves.K256()
	return dklsv1.NewBobSign(curve, sha3.New256(), msg, u.ExtractMessage(), protocol.Version1)
}

func (u *UserKeyshare) String() string {
	return u.encoded
}

func encodeMessage(m *protocol.Message) (string, error) {
	return protocol.EncodeMessage(m)
}

func decodeMessage(s string) (*protocol.Message, error) {
	return protocol.DecodeMessage(s)
}
