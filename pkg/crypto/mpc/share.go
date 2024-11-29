package mpc

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"

	"github.com/onsonr/sonr/pkg/crypto/core/curves"
	"github.com/onsonr/sonr/pkg/crypto/core/protocol"
	"github.com/onsonr/sonr/pkg/crypto/tecdsa/dklsv1"
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

type PublicKey *ecdsa.PublicKey

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
	Message   Message `json:"message"`
	Role      int     `json:"role"` // 1 for validator, 2 for user
	PublicKey []byte  `json:"public-key"`
}

func NewValKeyshare(msg Message) (*ValKeyshare, error) {
	valShare, err := dklsv1.DecodeAliceDkgResult(msg)
	if err != nil {
		return nil, err
	}
	return &ValKeyshare{
		Message:   msg,
		Role:      1,
		PublicKey: valShare.PublicKey.ToAffineUncompressed(),
	}, nil
}

func (v *ValKeyshare) GetPayloads() map[string][]byte {
	return v.Message.Payloads
}

func (v *ValKeyshare) GetMetadata() map[string]string {
	return v.Message.Metadata
}

func (v *ValKeyshare) GetPublicKey() []byte {
	return v.PublicKey
}

func (v *ValKeyshare) GetProtocol() string {
	return v.Message.Protocol
}

func (v *ValKeyshare) GetRole() int32 {
	return int32(v.Role)
}

func (v *ValKeyshare) GetVersion() uint32 {
	return uint32(v.Message.Version)
}

func (v *ValKeyshare) ECDSAPublicKey() (*ecdsa.PublicKey, error) {
	return ComputeEcdsaPublicKey(v.PublicKey)
}

func (v *ValKeyshare) ExtractMessage() *protocol.Message {
	return &protocol.Message{
		Payloads: v.GetPayloads(),
		Metadata: v.GetMetadata(),
		Protocol: v.GetProtocol(),
		Version:  uint(v.GetVersion()),
	}
}

func (v *ValKeyshare) RefreshFunc() (RefreshFunc, error) {
	curve := curves.K256()
	return dklsv1.NewAliceRefresh(curve, v.ExtractMessage(), protocol.Version1)
}

func (v *ValKeyshare) SignFunc(msg []byte) (SignFunc, error) {
	curve := curves.K256()
	return dklsv1.NewAliceSign(curve, sha3.New256(), msg, v.ExtractMessage(), protocol.Version1)
}

func (v *ValKeyshare) Marshal() ([]byte, error) {
	return json.Marshal(v.Message)
}

type UserKeyshare struct {
	Message   Message `json:"message"` // BobOutput
	Role      int     `json:"role"`    // 2 for user, 1 for validator
	PublicKey []byte  `json:"public-key"`
}

func NewUserKeyshare(msg Message) (*UserKeyshare, error) {
	out, err := dklsv1.DecodeBobDkgResult(msg)
	if err != nil {
		return nil, err
	}
	return &UserKeyshare{
		Message:   msg,
		Role:      2,
		PublicKey: out.PublicKey.ToAffineUncompressed(),
	}, nil
}

func (u *UserKeyshare) GetPayloads() map[string][]byte {
	return u.Message.Payloads
}

func (u *UserKeyshare) GetMetadata() map[string]string {
	return u.Message.Metadata
}

func (u *UserKeyshare) GetPublicKey() []byte {
	return u.PublicKey
}

func (u *UserKeyshare) GetProtocol() string {
	return u.Message.Protocol
}

func (u *UserKeyshare) GetRole() int32 {
	return int32(u.Role)
}

func (u *UserKeyshare) GetVersion() uint32 {
	return uint32(u.Message.Version)
}

func (u *UserKeyshare) ECDSAPublicKey() (*ecdsa.PublicKey, error) {
	return ComputeEcdsaPublicKey(u.PublicKey)
}

func (u *UserKeyshare) ExtractMessage() *protocol.Message {
	return &protocol.Message{
		Payloads: u.GetPayloads(),
		Metadata: u.GetMetadata(),
		Protocol: u.GetProtocol(),
		Version:  uint(u.GetVersion()),
	}
}

func (u *UserKeyshare) RefreshFunc() (RefreshFunc, error) {
	curve := curves.K256()
	return dklsv1.NewBobRefresh(curve, u.ExtractMessage(), protocol.Version1)
}

func (u *UserKeyshare) SignFunc(msg []byte) (SignFunc, error) {
	curve := curves.K256()
	return dklsv1.NewBobSign(curve, sha3.New256(), msg, u.ExtractMessage(), protocol.Version1)
}

func (u *UserKeyshare) Marshal() ([]byte, error) {
	jsonBytes, err := json.Marshal(u.Message)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}
