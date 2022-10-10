package common

import (
	"encoding/base64"
	"errors"
	"strings"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/multiformats/go-multiaddr"
)

type AccountId string

func AccountIdFromString(id string) (AccountId, error) {
	return AccountId(id), nil
}

// Checks if the AccountId is a AccountAddress on the Cosmos Blockchain
func (id AccountId) IsAddress() bool {
	return strings.HasPrefix(id.String(), "snr") && len(id.String()) == 44
}

func (id AccountId) IsAlias() bool {
	if strings.Contains(id.String(), ".snr") {
		spts := strings.Split(id.String(), ".")
		if len(spts) == 2 && spts[1] == "snr" {
			return true
		}
	}
	return false
}

func (id AccountId) IsDid() bool {
	if strings.Contains(string(id), "did:snr:") {
		spts := strings.Split(string(id), ".")
		if len(spts) == 2 && spts[1] == "snr" {
			return true
		}
	}
	return false
}

func (id AccountId) IsPeerId() bool {
	if strings.Contains(string(id), "Qm") {
		return true
	}
	return false
}

func (id AccountId) String() string {
	return string(id)
}

func AddrInfoFromBase64(aistr string) (*AddrInfo, error) {
	bz, err := base64.StdEncoding.DecodeString(aistr)
	if err != nil {
		return nil, err
	}
	var ai AddrInfo
	err = ai.Unmarshal(bz)
	if err != nil {
		return nil, err
	}
	return &ai, nil
}

// Base64 returns the base64 encoded string of the Marshaled AddrInfo
func (ai *AddrInfo) Base64() (string, error) {
	// Write AddrInfo to bytes
	bz, err := ai.Marshal()
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bz), nil
}

func (ai *AddrInfo) GetLinkProtocolId() (protocol.ID, error) {
	for _, proto := range ai.Protocols {
		if strings.Contains(proto, "/sonr/link/") {
			return protocol.ID(proto), nil
		}
	}
	return "", errors.New("no link protocol found")
}

// ListProtocols returns a list of all the protocols that are supported by the Host as a Libp2p Protocol ID
func (ai *AddrInfo) ListProtocols() []protocol.ID {
	protos := make([]protocol.ID, len(ai.Protocols))
	for i, proto := range ai.Protocols {
		protos[i] = protocol.ID(proto)
	}
	return protos
}

// ToLibp2pAddrInfo converts the Sonr common AddrInfo to a libp2p peer.AddrInfo
func (ai *AddrInfo) ToLibp2pAddrInfo() (peer.AddrInfo, error) {
	var err error
	maddrs := make([]multiaddr.Multiaddr, len(ai.Addrs))
	for i, addr := range ai.Addrs {
		maddrs[i], err = multiaddr.NewMultiaddr(addr)
	}
	if err != nil {
		return peer.AddrInfo{}, err
	}
	return peer.AddrInfo{
		ID:    peer.ID(ai.Id),
		Addrs: maddrs,
	}, nil
}
