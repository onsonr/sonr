package host

import (
	"time"

	"github.com/google/uuid"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-msgio"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/tools/logger"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

// NewID generates a new UUID value signed by the local node's private key
func (n *SNRHost) NewID() (*common.UUID, error) {
	// generate new UUID
	id := uuid.New().String()

	// sign UUID using local node's private key
	sig, err := n.SignData([]byte(id))
	if err != nil {
		logger.Error("Failed to sign UUID", zap.Error(err))
		return nil, err
	}

	// Return UUID with signature
	return &common.UUID{
		Value:     id,
		Signature: sig,
		Timestamp: time.Now().Unix(),
	}, nil
}

// NewMetadata generates message data shared between all node's p2p protocols
func (h *SNRHost) NewMetadata() (*common.Metadata, error) {
	// Get local node's public key
	pubKey, err := device.KeyChain.GetPubKey(device.Account)
	if err != nil {
		logger.Error("Failed to get local host's public key", zap.Error(err))
		return nil, err
	}

	// Marshal Public key into public key data
	nodePubKey, err := crypto.MarshalPublicKey(pubKey)
	if err != nil {
		logger.Error("Failed to Extract Public Key", zap.Error(err))
		return nil, err
	}

	// Generate new Metadata
	return &common.Metadata{
		Timestamp: time.Now().Unix(),
		PublicKey: nodePubKey,
		NodeId:    peer.Encode(h.ID()),
	}, nil
}

// AuthenticateId verifies UUID value and signature
func (h *SNRHost) AuthenticateId(id *common.UUID) (bool, error) {
	// Get local node's public key
	pubKey, err := device.KeyChain.GetPubKey(device.Account)
	if err != nil {
		logger.Error("Failed to get local host's public key", zap.Error(err))
		return false, err
	}

	// verify UUID value
	result, err := pubKey.Verify([]byte(id.GetValue()), []byte(id.GetSignature()))
	if err != nil {
		logger.Error("Failed to verify signature of UUID", zap.Error(err))
		return false, err
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
		logger.Error("Failed to marshal Protobuf Message.", zap.Error(err))
		return false
	}

	// restore sig in message data (for possible future use)
	data.Signature = sign

	// restore peer id binary format from base58 encoded node id data
	peerId, err := peer.Decode(data.NodeId)
	if err != nil {
		logger.Error("Failed to decode node id from base58.", zap.Error(err))
		return false
	}

	// verify the data was authored by the signing peer identified by the public key
	// and signature included in the message
	return n.VerifyData(bin, []byte(sign), peerId, data.PublicKey)
}

// SendMessage writes a protobuf go data object to a network stream
func (n *SNRHost) SendMessage(id peer.ID, p protocol.ID, data proto.Message) error {
	s, err := n.NewStream(n.ctxHost, id, p)
	if err != nil {
		logger.Error("Failed to start stream", zap.Error(err))
		return err
	}
	defer s.Close()

	// marshall data to protobufs3 binary format
	bin, err := proto.Marshal(data)
	if err != nil {
		logger.Error("Failed to marshal pb", zap.Error(err))
		return err
	}

	// Create Writer and write data to stream
	w := msgio.NewWriter(s)
	if err := w.WriteMsg(bin); err != nil {
		logger.Error("Failed to write message to stream.", zap.Error(err))
		return err
	}
	return nil
}

// SignMessage signs an outgoing p2p message payload
func (n *SNRHost) SignMessage(message proto.Message) ([]byte, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	return n.SignData(data)
}

// SignData signs an outgoing p2p message payload
func (n *SNRHost) SignData(data []byte) ([]byte, error) {
	// Get local node's private key
	privKey, err := device.KeyChain.GetPrivKey(device.Account)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get private key")
	}

	// sign data using local node's private key
	res, err := privKey.Sign(data)
	if err != nil {
		logger.Error("Failed to sign data.", zap.Error(err))
		return nil, err
	}
	return res, nil
}

// VerifyData verifies incoming p2p message data integrity
func (n *SNRHost) VerifyData(data []byte, signature []byte, peerId peer.ID, pubKeyData []byte) bool {
	key, err := crypto.UnmarshalPublicKey(pubKeyData)
	if err != nil {
		logger.Error("Failed to extract key from message key data", zap.Error(err))
		return false
	}

	// extract node id from the provided public key
	idFromKey, err := peer.IDFromPublicKey(key)
	if err != nil {
		logger.Error("Failed to extract peer id from public key", zap.Error(err))
		return false
	}

	// verify that message author node id matches the provided node public key
	if idFromKey != peerId {
		logger.Error("Node id and provided public key mismatch", zap.Error(err))
		return false
	}

	res, err := key.Verify(data, signature)
	if err != nil {
		logger.Error("Error authenticating data", zap.Error(err))
		return false
	}

	return res
}
