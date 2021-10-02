package common

import "github.com/sonr-io/core/internal/keychain"

// SignedMetadataToProto converts a SignedMetadata to a protobuf.
func SignedMetadataToProto(m *keychain.SignedMetadata) *Metadata {
	return &Metadata{
		Timestamp: m.Timestamp,
		NodeId:    m.NodeId,
		PublicKey: m.PublicKey,
	}
}

// SignedUUIDToProto converts a SignedUUID to a protobuf.
func SignedUUIDToProto(m *keychain.SignedUUID) *UUID {
	return &UUID{
		Timestamp: m.Timestamp,
		Signature: m.Signature,
		Value:     m.Value,
	}
}
