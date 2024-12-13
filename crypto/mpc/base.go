package mpc

import (
	"github.com/onsonr/sonr/crypto/core/protocol"
	"github.com/onsonr/sonr/crypto/keys"
	"github.com/onsonr/sonr/crypto/tecdsa/dklsv1/dkg"
)

// BaseKeyshare contains common fields and methods for both validator and user keyshares
type BaseKeyshare struct {
	Message *protocol.Message `json:"message"`
	Role    int               `json:"role"`
	PubKey  keys.PubKey
}

func initFromAlice(aliceOut *dkg.AliceOutput, originalMsg *protocol.Message) BaseKeyshare {
	return BaseKeyshare{
		Message: originalMsg,
		Role:    1,
		PubKey:  keys.NewPubKey(aliceOut.PublicKey, keys.DIDMethodSonr),
	}
}

func initFromBob(bobOut *dkg.BobOutput, originalMsg *protocol.Message) BaseKeyshare {
	return BaseKeyshare{
		Message: originalMsg,
		Role:    2,
		PubKey:  keys.NewPubKey(bobOut.PublicKey, keys.DIDMethodSonr),
	}
}

func (b *BaseKeyshare) PublicKey() keys.PubKey {
	return b.PubKey
}

func (b *BaseKeyshare) GetPayloads() map[string][]byte {
	return b.Message.Payloads
}

func (b *BaseKeyshare) GetMetadata() map[string]string {
	return b.Message.Metadata
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

func (b *BaseKeyshare) ExtractMessage() *protocol.Message {
	return &protocol.Message{
		Payloads: b.GetPayloads(),
		Metadata: b.GetMetadata(),
		Protocol: b.GetProtocol(),
		Version:  uint(b.GetVersion()),
	}
}
