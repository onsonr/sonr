package mpc

import (
	"crypto/ecdsa"
	"encoding/json"

	"github.com/onsonr/sonr/pkg/crypto/core/curves"
	"github.com/onsonr/sonr/pkg/crypto/core/protocol"
	"github.com/onsonr/sonr/pkg/crypto/tecdsa/dklsv1"
	"github.com/onsonr/sonr/pkg/crypto/tecdsa/dklsv1/dkg"
	"golang.org/x/crypto/sha3"
)

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

type Share interface {
	Equals(o Share) bool
	GetPayloads() map[string][]byte
	GetMetadata() map[string]string
	GetPublicKey() []byte
	GetProtocol() string
	GetRole() int32
	GetVersion() uint32
	ECDSAPublicKey() (*ecdsa.PublicKey, error)
	ExtractMessage() *protocol.Message
	RefreshFunc() (RefreshFunc, error)
	SignFunc(msg []byte) (SignFunc, error)
	Marshal() (string, error)
	Unmarshal(data string) error
}

func NewKeyshareArray(val Message, user Message) ([]Share, error) {
	valShare, err := dklsv1.DecodeAliceDkgResult(val)
	if err != nil {
		return nil, err
	}
	userShare, err := dklsv1.DecodeBobDkgResult(user)
	if err != nil {
		return nil, err
	}
	return []Share{NewValKeyshare(valShare, val), NewUserKeyshare(userShare, user)}, nil
}

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

func NewValKeyshare(out *dkg.AliceOutput, msg Message) ValKeyshare {
	return ValKeyshare{
		Message:   msg,
		Role:      1,
		PublicKey: out.PublicKey.ToAffineUncompressed(),
	}
}

func (v ValKeyshare) GetPayloads() map[string][]byte {
	return v.Message.Payloads
}

func (v ValKeyshare) GetMetadata() map[string]string {
	return v.Message.Metadata
}

func (v ValKeyshare) GetPublicKey() []byte {
	return v.PublicKey
}

func (v ValKeyshare) GetProtocol() string {
	return v.Message.Protocol
}

func (v ValKeyshare) GetRole() int32 {
	return int32(v.Role)
}

func (v ValKeyshare) GetVersion() uint32 {
	return uint32(v.Message.Version)
}

func (k ValKeyshare) ECDSAPublicKey() (*ecdsa.PublicKey, error) {
	return ComputeEcdsaPublicKey(k.PublicKey)
}

func (k ValKeyshare) ExtractMessage() *protocol.Message {
	return &protocol.Message{
		Payloads: k.GetPayloads(),
		Metadata: k.GetMetadata(),
		Protocol: k.GetProtocol(),
		Version:  uint(k.GetVersion()),
	}
}

func (k ValKeyshare) RefreshFunc() (RefreshFunc, error) {
	curve := curves.K256()
	return dklsv1.NewAliceRefresh(curve, k.ExtractMessage(), protocol.Version1)
}

func (k ValKeyshare) SignFunc(msg []byte) (SignFunc, error) {
	curve := curves.K256()
	return dklsv1.NewAliceSign(curve, sha3.New256(), msg, k.ExtractMessage(), protocol.Version1)
}

func (v ValKeyshare) Equals(o Share) bool {
	return v.GetProtocol() == o.GetProtocol() &&
		v.GetVersion() == o.GetVersion() &&
		v.GetRole() == o.GetRole()
}

func (v ValKeyshare) Marshal() (string, error) {
	jsonBytes, err := json.Marshal(v)
	return string(jsonBytes), err
}

func (v ValKeyshare) Unmarshal(data string) error {
	return json.Unmarshal([]byte(data), &v)
}

type UserKeyshare struct {
	Message   Message `json:"message"` // BobOutput
	Role      int     `json:"role"`    // 2 for user, 1 for validator
	PublicKey []byte  `json:"public-key"`
}

func NewUserKeyshare(out *dkg.BobOutput, msg Message) UserKeyshare {
	return UserKeyshare{
		Message:   msg,
		Role:      2,
		PublicKey: out.PublicKey.ToAffineUncompressed(),
	}
}

func (u UserKeyshare) GetPayloads() map[string][]byte {
	return u.Message.Payloads
}

func (u UserKeyshare) GetMetadata() map[string]string {
	return u.Message.Metadata
}

func (u UserKeyshare) GetPublicKey() []byte {
	return u.PublicKey
}

func (u UserKeyshare) GetProtocol() string {
	return u.Message.Protocol
}

func (u UserKeyshare) GetRole() int32 {
	return int32(u.Role)
}

func (u UserKeyshare) GetVersion() uint32 {
	return uint32(u.Message.Version)
}

func (k UserKeyshare) ECDSAPublicKey() (*ecdsa.PublicKey, error) {
	return ComputeEcdsaPublicKey(k.PublicKey)
}

func (u UserKeyshare) Equals(o Share) bool {
	return u.GetProtocol() == o.GetProtocol() &&
		u.GetVersion() == o.GetVersion() &&
		u.GetRole() == o.GetRole()
}

func (k UserKeyshare) ExtractMessage() *protocol.Message {
	return &protocol.Message{
		Payloads: k.GetPayloads(),
		Metadata: k.GetMetadata(),
		Protocol: k.GetProtocol(),
		Version:  uint(k.GetVersion()),
	}
}

func (k UserKeyshare) RefreshFunc() (RefreshFunc, error) {
	curve := curves.K256()
	return dklsv1.NewBobRefresh(curve, k.ExtractMessage(), protocol.Version1)
}

func (k UserKeyshare) SignFunc(msg []byte) (SignFunc, error) {
	curve := curves.K256()
	return dklsv1.NewBobSign(curve, sha3.New256(), msg, k.ExtractMessage(), protocol.Version1)
}

func (u UserKeyshare) Marshal() (string, error) {
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func (u UserKeyshare) Unmarshal(data string) error {
	return json.Unmarshal([]byte(data), &u)
}

// GetRawPublicKey is the public key for the keyshare
func GetRawPublicKey(ks Share) ([]byte, error) {
	role := Role(ks.GetRole())
	if role.IsUser() {
		bobOut, err := dklsv1.DecodeBobDkgResult(ks.ExtractMessage())
		if err != nil {
			return nil, err
		}
		return bobOut.PublicKey.ToAffineUncompressed(), nil
	} else if role.IsValidator() {
		aliceOut, err := dklsv1.DecodeAliceDkgResult(ks.ExtractMessage())
		if err != nil {
			return nil, err
		}
		return aliceOut.PublicKey.ToAffineUncompressed(), nil
	}
	return nil, ErrInvalidKeyshareRole
}
