package mpc

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/types/bech32"
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

func computeSonrAddr(pk []byte) (string, error) {
	sonrAddr, err := bech32.ConvertAndEncode("idx", pk)
	if err != nil {
		return "", err
	}
	return sonrAddr, nil
}

func NewValKeyshare(msg *protocol.Message) (*ValKeyshare, error) {
	vks := new(ValKeyshare)
	encoded, err := protocol.EncodeMessage(msg)
	if err != nil {
		return nil, err
	}
	valShare, err := dklsv1.DecodeAliceDkgResult(msg)
	if err != nil {
		return nil, err
	}

	vks.BaseKeyshare = initFromAlice(valShare, msg)
	vks.encoded = encoded
	return vks, nil
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

// PublicKey returns the uncompressed public key (65 bytes)
func (v *ValKeyshare) PublicKey() []byte {
	return v.UncompressedPubKey
}

// CompressedPublicKey returns the compressed public key (33 bytes)
func (v *ValKeyshare) CompressedPublicKey() []byte {
	return v.CompressedPubKey
}

type UserKeyshare struct {
	BaseKeyshare
	encoded string
}

func NewUserKeyshare(msg *protocol.Message) (*UserKeyshare, error) {
	uks := new(UserKeyshare)
	encoded, err := protocol.EncodeMessage(msg)
	if err != nil {
		return nil, err
	}
	out, err := dklsv1.DecodeBobDkgResult(msg)
	if err != nil {
		return nil, err
	}

	uks.BaseKeyshare = initFromBob(out, msg)
	uks.encoded = encoded
	return uks, nil
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

// PublicKey returns the uncompressed public key (65 bytes)
func (u *UserKeyshare) PublicKey() []byte {
	return u.UncompressedPubKey
}

// CompressedPublicKey returns the compressed public key (33 bytes)
func (u *UserKeyshare) CompressedPublicKey() []byte {
	return u.CompressedPubKey
}

func encodeMessage(m *protocol.Message) (string, error) {
	return protocol.EncodeMessage(m)
}

func decodeMessage(s string) (*protocol.Message, error) {
	return protocol.DecodeMessage(s)
}
