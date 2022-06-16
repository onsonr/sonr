package crypto

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/golang/protobuf/proto"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp/config"
)

type PublicKey struct {
	proto.Message
	Key *config.Public
}

func NewPublicKey(key *config.Public) cryptotypes.PubKey {
	return &PublicKey{
		Key: key,
	}
}

func (pk *PublicKey) Address() cryptotypes.Address {
	return nil
}

func (pk *PublicKey) Bytes() []byte {
	return nil
}
func (pk *PublicKey) VerifySignature(msg []byte, sig []byte) bool {
	return true
}
func (pk *PublicKey) Equals(cryptotypes.PubKey) bool {
	return false
}

func (pk *PublicKey) Type() string {
	return "public"
}
