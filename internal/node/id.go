package node

import (
	"errors"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/sonr-io/multi-party-sig/pkg/math/curve"
	"github.com/sonr-io/multi-party-sig/pkg/party"
)

type Address string

type PeerID = peer.ID

type MPCPublicPoint = curve.Point

type PartyID = party.ID

type ID interface {
	GetPublicBytes() ([]byte, error)
	GetMPCPublicPoint() (curve.Point, error)
	GetPartyID() party.ID
	GetPeerID() (peer.ID, error)
	GetP2PPubKey() (crypto.PubKey, error)
	IsEquals(interface{}) bool
	String() string
}

// IDSlice is a list of IDs
type IDSlice []ID

func (idl IDSlice) ToPartyIDSlice() party.IDSlice {
	ids := make(party.IDSlice, len(idl))
	for i, id := range idl {
		ids[i] = id.GetPartyID()
	}
	return ids
}

// ID is a string type used for common cross-operations between the Libp2p Protocol, the MPC Protocol and the Blockchain
type iD string

// ParseID parses an ID from a string, curve.Point, party.ID, peer.ID or crypto.PubKey by checking all the possible types
func ParseID(id interface{}) (ID, error) {
	switch id.(type) {
	case string:
		return newIDFromAddress(id.(string))
	case curve.Point:
		return newIDFromPublicPoint(id.(curve.Point))
	case party.ID:
		return newIDFromPartyID(id.(party.ID))
	case peer.ID:
		return newIDFromPeerID(id.(peer.ID))
	case crypto.PubKey:
		return newIDFromP2PPubKey(id.(crypto.PubKey))
	}
	return nil, errors.New("invalid ID type")
}

// NewIDFromAddress creates a new ID from a blockchain address or a proper did address. All Addresses are normalized to bech32 snr addresses
func newIDFromAddress(str string) (ID, error) {
	if strings.Contains(str, "did:snr") {
		ptrs := strings.Split(str, ":")
		if len(ptrs) != 3 {
			return iD(""), errors.New("invalid address")
		}
		return iD(strings.Join(ptrs[1:], "")), nil
	}
	if strings.Contains(str, "snr") {
		return iD(str), nil
	}
	return iD(""), errors.New("invalid address")
}

// NewIDFromPublicPoint creates a new ID from a public point that is the result of a MPC protocol
func newIDFromPublicPoint(point curve.Point) (iD, error) {
	pubBz, err := point.MarshalBinary()
	if err != nil {
		return iD(""), err
	}
	bech32Pub, err := bech32.ConvertAndEncode("snr", pubBz)
	if err != nil {
		return iD(""), err
	}
	return iD(bech32Pub), nil
}

// newIDFromPartyID creates a new ID from a party ID
func newIDFromPartyID(id party.ID) (ID, error) {
	return iD(id), nil
}

// NewIDFromPeerID creates a new ID from a libp2p peer ID
func newIDFromPeerID(id peer.ID) (ID, error) {
	pbBz, err := id.ExtractPublicKey()
	if err != nil {
		return iD(""), err
	}
	pubkeyBytes, err := pbBz.Raw()
	if err != nil {
		return iD(""), err
	}
	bech32Pub, err := bech32.ConvertAndEncode("snr", pubkeyBytes)
	if err != nil {
		return iD(""), err
	}
	return iD(bech32Pub), nil
}

// newIDFromP2PPubKey creates a new ID from a libp2p public key
func newIDFromP2PPubKey(pubkey crypto.PubKey) (ID, error) {
	pubkeyBytes, err := pubkey.Raw()
	if err != nil {
		return iD(""), err
	}
	bech32Pub, err := bech32.ConvertAndEncode("snr", pubkeyBytes)
	if err != nil {
		return iD(""), err
	}
	return iD(bech32Pub), nil
}

// GetPublicBytes returns the public bytes of the ID
func (i iD) GetPublicBytes() ([]byte, error) {
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
func (i iD) GetMPCPublicPoint() (curve.Point, error) {
	bz, err := i.GetPublicBytes()
	p := curve.Secp256k1Point{}
	err = p.UnmarshalBinary(bz)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPartyID returns the party ID of the ID
func (i iD) GetPartyID() party.ID {
	return party.ID(i.String())
}

// GetPeerID returns the libp2p peer ID of the ID
func (i iD) GetPeerID() (peer.ID, error) {
	bz, err := i.GetPublicBytes()
	if err != nil {
		return "", err
	}
	return peer.IDFromBytes(bz)
}

// GetP2PPubKey returns the libp2p public key of the ID
func (i iD) GetP2PPubKey() (crypto.PubKey, error) {
	pbBz, err := i.GetPublicBytes()
	if err != nil {
		return nil, err
	}
	return crypto.UnmarshalSecp256k1PublicKey(pbBz)
}

// IsEquals checks if the ID is equal to either another ID, MPCPublicPoint, PartyID, PeerID or P2PPubKey by checking all the possible types
func (i iD) IsEquals(id interface{}) bool {
	switch id.(type) {
	case ID:
		return i.String() == id.(ID).String()
	case curve.Point:
		return i.isEqualsMPCPublicPoint(id.(curve.Point))
	case party.ID:
		return i.isEqualsPartyID(id.(party.ID))
	case peer.ID:
		return i.isEqualsPeerID(id.(peer.ID))
	case crypto.PubKey:
		return i.isEqualsP2PPubKey(id.(crypto.PubKey))
	}
	return false
}

// String returns the string representation of the Blockchain address
func (i iD) String() string {
	return string(i)
}

// IsEqualsMPCPublicPoint checks if the ID is a MPC public point
func (i iD) isEqualsMPCPublicPoint(point curve.Point) bool {
	bz, err := i.GetPublicBytes()
	if err != nil {
		return false
	}
	p := curve.Secp256k1Point{}
	err = p.UnmarshalBinary(bz)
	if err != nil {
		return false
	}
	return p.Equal(point)
}

// IsEqualsPartyID checks if the ID is a party ID
func (i iD) isEqualsPartyID(id party.ID) bool {
	return i.String() == string(id)
}

// isEqualsPeerID checks if the ID is a libp2p peer ID
func (i iD) isEqualsPeerID(id peer.ID) bool {
	peerID, err := i.GetPeerID()
	if err != nil {
		return false
	}
	return peerID == id
}

// isEqualsP2PPubKey checks if the ID is a libp2p public key
func (i iD) isEqualsP2PPubKey(key crypto.PubKey) bool {
	p2pKey, err := i.GetP2PPubKey()
	if err != nil {
		return false
	}
	return p2pKey == key
}
