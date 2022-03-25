package wallet

import (
	"time"

	"github.com/google/uuid"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
)

// SignedMetadata is a struct to be used for signing metadata.
type SignedMetadata struct {
	Timestamp int64
	PublicKey []byte
	NodeId    string
}

// SignedUUID is a struct to be converted into a UUID.
type SignedUUID struct {
	Timestamp int64
	Signature []byte
	Value     string
}

// CreateUUID makes a new UUID value signed by the local node's private key
func CreateUUID() (*SignedUUID, error) {
	// generate new UUID
	id := uuid.New().String()

	// sign UUID using local node's private key
	sig, err := Sign([]byte(id))
	if err != nil {
		logger.Errorf("%s - Failed to sign UUID", err)
		return nil, err
	}

	// Return UUID with signature
	return &SignedUUID{
		Value:     id,
		Signature: sig,
		Timestamp: time.Now().Unix(),
	}, nil
}

// CreateMetadata makes message data shared between all node's p2p protocols
func CreateMetadata(peerID peer.ID) (*SignedMetadata, error) {
	// Get local node's public key
	pubKey, err := DevicePubKey()
	if err != nil {
		logger.Errorf("%s - Failed to get local host's public key", err)
		return nil, err
	}

	// Marshal Public key into public key data
	nodePubKey, err := crypto.MarshalPublicKey(pubKey)
	if err != nil {
		logger.Errorf("%s - Failed to Extract Public Key", err)
		return nil, err
	}

	// Generate new Metadata
	return &SignedMetadata{
		Timestamp: time.Now().Unix(),
		PublicKey: nodePubKey,
		NodeId:    peer.Encode(peerID),
	}, nil
}
