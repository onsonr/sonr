package mpc

import (
	"errors"

	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/sonr-hq/sonr/internal/node"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"

	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

type MpcProtocol struct {
	selfNode *node.Node
	sessions map[string]*Session
}

// Initialize configures the Node for the MPC Protocol
func Initialize(n *node.Node) (*MpcProtocol, error) {
	protocol := &MpcProtocol{
		selfNode: n,
		sessions: make(map[string]*Session, 0),
	}

	return protocol, nil
}

// GenerateWallet generates a new wallet
func (m *MpcProtocol) JoinCMPKeygen(peers ...peer.ID) (*cmp.Config, error) {
	if len(peers) == 0 {
		return nil, errors.New("no peers provided")
	}
	// Setup Run
	pl := pool.NewPool(0)
	defer pl.TearDown()
	ids := party.IDSlice{
		party.ID(m.selfNode.ID()),
	}

	// Add All Peers
	for _, peer := range peers {
		ids = append(ids, party.ID(peer))
	}

	// Build Session
	s, err := NewSession(m.selfNode, party.ID(m.selfNode.ID()), ids)
	if err != nil {
		return nil, err
	}

	// Run Protocol
	r, err := s.RunProtocol(cmp.Keygen(curve.Secp256k1{}, party.ID(m.selfNode.ID()), ids, 1, pl), kCmpKeygenRequest, nil)
	if err != nil {
		return nil, err
	}
	conf := r.(*cmp.Config)
	return conf, nil
}
