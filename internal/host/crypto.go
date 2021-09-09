package host

import (
	"context"
	"log"
	"time"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/tools/logger"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

// AuthenticateMessage Authenticates incoming p2p message
func (n *SHost) AuthenticateMessage(message proto.Message, data *common.Metadata) bool {
	// store a temp ref to signature and remove it from message data
	// sign is a string to allow easy reset to zero-value (empty string)
	sign := data.Signature
	data.Signature = nil

	// marshall data without the signature to protobufs3 binary format
	bin, err := proto.Marshal(message)
	if err != nil {
		log.Println(err, "failed to marshal pb message")
		return false
	}

	// restore sig in message data (for possible future use)
	data.Signature = sign

	// restore peer id binary format from base58 encoded node id data
	peerId, err := peer.Decode(data.NodeId)
	if err != nil {
		log.Println(err, "Failed to decode node id from base58")
		return false
	}

	// verify the data was authored by the signing peer identified by the public key
	// and signature included in the message
	return n.VerifyData(bin, []byte(sign), peerId, data.PublicKey)
}

// SignProtoMessage signs an outgoing p2p message payload
func (n *SHost) SignProtoMessage(message proto.Message) ([]byte, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	return n.SignData(data)
}

// sign binary data using the local node's private key
func (n *SHost) SignData(data []byte) ([]byte, error) {
	res, err := n.privKey.Sign(data)
	return res, err
}

// VerifyData verifies incoming p2p message data integrity
func (n *SHost) VerifyData(data []byte, signature []byte, peerId peer.ID, pubKeyData []byte) bool {
	key, err := crypto.UnmarshalPublicKey(pubKeyData)
	if err != nil {
		log.Println(err, "Failed to extract key from message key data")
		return false
	}

	// extract node id from the provided public key
	idFromKey, err := peer.IDFromPublicKey(key)

	if err != nil {
		log.Println(err, "Failed to extract peer id from public key")
		return false
	}

	// verify that message author node id matches the provided node public key
	if idFromKey != peerId {
		log.Println(err, "Node id and provided public key mismatch")
		return false
	}

	res, err := key.Verify(data, signature)
	if err != nil {
		log.Println(err, "Error authenticating data")
		return false
	}

	return res
}

// NewMetadata generates message data shared between all node's p2p protocols
func (n *SHost) NewMetadata() *common.Metadata {
	nodePubKey, err := n.privKey.GetPublic().Raw()
	if err != nil {
		log.Println(err, "Failed to extract public key")
		return nil
	}

	return &common.Metadata{
		Timestamp: time.Now().Unix(),
		PublicKey: nodePubKey,
		NodeId:    peer.Encode(n.id),
	}
}

// SendProtoMessage writes a protobuf go data object to a network stream
func (n *SHost) SendProtoMessage(id peer.ID, p protocol.ID, data proto.Message) error {
	s, err := n.NewStream(context.Background(), id, p)
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
