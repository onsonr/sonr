package host

import (
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/keychain"
	"github.com/sonr-io/core/tools/logger"
	"google.golang.org/protobuf/proto"
)

// AuthenticateId verifies UUID value and signature
func (h *SNRHost) AuthenticateId(id *common.UUID) (bool, error) {
	// Get local node's public key
	pubKey, err := device.KeyChain.GetPubKey(keychain.Account)
	if err != nil {
		return false, logger.Error("Failed to get local host's public key", err)
	}

	// verify UUID value
	result, err := pubKey.Verify([]byte(id.GetValue()), []byte(id.GetSignature()))
	if err != nil {
		return false, logger.Error("Failed to verify signature of UUID", err)
	}
	return result, nil
}

// AuthenticateMessage Authenticates incoming p2p message
func (n *SNRHost) AuthenticateMessage(message proto.Message, data *common.Metadata) bool {
	// store a temp ref to signature and remove it from message data
	// sign is a string to allow easy reset to zero-value (empty string)
	sign := data.Signature
	data.Signature = nil

	// marshall data without the signature to protobufs3 binary format
	bin, err := proto.Marshal(message)
	if err != nil {
		logger.Error("Failed to marshal Protobuf Message.", err)
		return false
	}

	// restore sig in message data (for possible future use)
	data.Signature = sign

	// restore peer id binary format from base58 encoded node id data
	peerId, err := peer.Decode(data.NodeId)
	if err != nil {
		logger.Error("Failed to decode node id from base58.", err)
		return false
	}

	// verify the data was authored by the signing peer identified by the public key
	// and signature included in the message
	return n.VerifyData(bin, []byte(sign), peerId, data.PublicKey)
}

// SignMessage signs an outgoing p2p message payload
func (n *SNRHost) SignMessage(message proto.Message) ([]byte, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, logger.Error("Failed to Sign Message", err)
	}
	return n.SignData(data)
}

// SignData signs an outgoing p2p message payload
func (n *SNRHost) SignData(data []byte) ([]byte, error) {
	// Get local node's private key
	res, err := device.KeyChain.SignWith(keychain.Account, data)
	if err != nil {
		return nil, logger.Error("Failed to get local host's private key", err)
	}
	return res, nil
}

// VerifyData verifies incoming p2p message data integrity
func (n *SNRHost) VerifyData(data []byte, signature []byte, peerId peer.ID, pubKeyData []byte) bool {
	key, err := crypto.UnmarshalPublicKey(pubKeyData)
	if err != nil {
		logger.Error("Failed to extract key from message key data", err)
		return false
	}

	// extract node id from the provided public key
	idFromKey, err := peer.IDFromPublicKey(key)
	if err != nil {
		logger.Error("Failed to extract peer id from public key", err)
		return false
	}

	// verify that message author node id matches the provided node public key
	if idFromKey != peerId {
		logger.Error("Node id and provided public key mismatch", err)
		return false
	}

	res, err := key.Verify(data, signature)
	if err != nil {
		logger.Error("Error authenticating data", err)
		return false
	}
	return res
}
