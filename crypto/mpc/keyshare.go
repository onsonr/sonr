package mpc

import (
	"crypto/ecdsa"

	"github.com/onsonr/sonr/crypto/core/protocol"
)

// Keyshare represents the common interface for both validator and user keyshares
type Keyshare interface {
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
}

// BaseKeyshare contains common fields and methods for both validator and user keyshares
type BaseKeyshare struct {
	Message            *protocol.Message `json:"message"`
	Role               int               `json:"role"`
	UncompressedPubKey []byte            `json:"public_key"`
	CompressedPubKey   []byte            `json:"compressed_public_key"`
}

func (b *BaseKeyshare) GetPayloads() map[string][]byte {
	return b.Message.Payloads
}

func (b *BaseKeyshare) GetMetadata() map[string]string {
	return b.Message.Metadata
}

func (b *BaseKeyshare) GetPublicKey() []byte {
	return b.UncompressedPubKey
}

func (b *BaseKeyshare) GetProtocol() string {
	return b.Message.Protocol
}

func (b *BaseKeyshare) GetRole() int32 {
	return int32(b.Role)
}

func (b *BaseKeyshare) GetVersion() uint32 {
	return uint32(b.Message.Version)
}

func (b *BaseKeyshare) ECDSAPublicKey() (*ecdsa.PublicKey, error) {
	return ComputeEcdsaPublicKey(b.UncompressedPubKey)
}

func (b *BaseKeyshare) ExtractMessage() *protocol.Message {
	return &protocol.Message{
		Payloads: b.GetPayloads(),
		Metadata: b.GetMetadata(),
		Protocol: b.GetProtocol(),
		Version:  uint(b.GetVersion()),
	}
}
