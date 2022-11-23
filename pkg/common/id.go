package common

import (
	"errors"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/sonr-io/multi-party-sig/pkg/math/curve"
)

// ID is a string type used for common cross-operations between the Libp2p Protocol, the MPC Protocol and the Blockchain
type ID string

// NewIDFromAddress creates a new ID from a blockchain address or a proper did address. All Addresses are normalized to bech32 snr addresses
func NewIDFromAddress(str string) (ID, error) {
	if strings.Contains(str, "did:snr") {
		ptrs := strings.Split(str, ":")
		if len(ptrs) != 3 {
			return "", errors.New("invalid address")
		}
		return ID(strings.Join(ptrs[1:], "")), nil
	}
	if strings.Contains(str, "snr") {
		return ID(str), nil
	}
	return "", errors.New("invalid address")
}

// NewIDFromPublicPoint creates a new ID from a public point that is the result of a MPC protocol
func NewIDFromPublicPoint(point curve.Point) (ID, error) {
	pubBz, err := point.MarshalBinary()
	if err != nil {
		return "", err
	}
	bech32Pub, err := bech32.ConvertAndEncode("snr", pubBz)
	if err != nil {
		return "", err
	}
	return ID(bech32Pub), nil
}

// NewIDFromPeerID creates a new ID from a libp2p peer ID
func NewIDFromPeerID(id peer.ID) (ID, error) {
	pbBz, err := id.ExtractPublicKey()
	if err != nil {
		return "", err
	}
	pubkeyBytes, err := pbBz.Raw()
	if err != nil {
		return "", err
	}
	bech32Pub, err := bech32.ConvertAndEncode("snr", pubkeyBytes)
	if err != nil {
		return "", err
	}
	return ID(bech32Pub), nil
}

// GetPublicBytes returns the public bytes of the ID
func (i ID) GetPublicBytes() ([]byte, error) {
	preFix, bz, err := bech32.DecodeAndConvert(i.String())
	if err != nil {
		return nil, err
	}
	if preFix != "snr" {
		return nil, errors.New("invalid prefix")
	}
	return bz, nil
}

// GetMPCPublicPoint returns the MPC public point of the ID
func (i ID) GetMPCPublicPoint() (curve.Point, error) {
	bz, err := i.GetPublicBytes()
	p := curve.Secp256k1Point{}
	err = p.UnmarshalBinary(bz)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPeerID returns the libp2p peer ID of the ID
func (i ID) GetPeerID() (peer.ID, error) {
	bz, err := i.GetPublicBytes()
	if err != nil {
		return "", err
	}
	pubkey, err := crypto.UnmarshalPublicKey(bz)
	if err != nil {
		return "", err
	}
	return peer.IDFromPublicKey(pubkey)
}

// GetP2PPubKey returns the libp2p public key of the ID
func (i ID) GetP2PPubKey() (crypto.PubKey, error) {
	pbBz, err := i.GetPublicBytes()
	if err != nil {
		return nil, err
	}
	return crypto.UnmarshalSecp256k1PublicKey(pbBz)
}

// String returns the string representation of the Blockchain address
func (i ID) String() string {
	return string(i)
}
