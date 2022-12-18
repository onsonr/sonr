package mpc

import (
	"errors"
	"fmt"

	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/sonr-hq/sonr/pkg/node"
	"github.com/sonr-hq/sonr/pkg/wallet"
	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"

	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

type MpcProtocol struct {
	selfId   party.ID
	selfNode *node.Node
	sessions map[string]*Session
}

// Initialize configures the Node for the MPC Protocol
func Initialize(n *node.Node) (*MpcProtocol, error) {
	protocol := &MpcProtocol{
		selfId:   party.ID(n.ID()),
		selfNode: n,
		sessions: make(map[string]*Session, 0),
	}

	return protocol, nil
}

// JoinCMPKeygen generates a new wallet
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

	// Build Session
	s, err := NewSession(m.selfNode, party.ID(m.selfNode.ID()), convertToPartyIDs(peers), kCmpKeygenFeed)
	if err != nil {
		fmt.Println("Error creating handler")
		return nil, err
	}

	// Run Protocol
	r, err := s.RunProtocol(cmp.Keygen(curve.Secp256k1{}, party.ID(m.selfNode.ID()), ids, 1, pl), nil)
	if err != nil {
		fmt.Println("Error creating handler")
		return nil, err
	}
	conf := r.(*cmp.Config)
	return conf, nil
}

// JoinCMPSign signs a new signature
func (m *MpcProtocol) JoinCMPSign(ws wallet.WalletShare, msg []byte, peers ...peer.ID) (*ecdsa.Signature, error) {
	if len(peers) == 0 {
		return nil, errors.New("no peers provided")
	}
	ids := convertToPartyIDs(peers)
	ids = append(ids, m.selfId)

	// Setup Run
	pl := pool.NewPool(0)
	defer pl.TearDown()

	// Build Session
	s, err := NewSession(m.selfNode, party.ID(m.selfNode.ID()), convertToPartyIDs(peers), kCmpKeygenFeed)
	if err != nil {
		fmt.Println("Error creating handler")
		return nil, err
	}

	// Run Protocol
	r, err := s.RunProtocol(cmp.Sign(ws.MPCConfig(), ids, msg, pl), nil)
	if err != nil {
		fmt.Println("Error creating handler")
		return nil, err
	}
	sig := r.(*ecdsa.Signature)
	return sig, nil
}

// JoinCMPPreSign presigns a new signature
func (m *MpcProtocol) JoinCMPPreSign(ws wallet.WalletShare, peers ...peer.ID) (*ecdsa.PreSignature, error) {
	if len(peers) == 0 {
		return nil, errors.New("no peers provided")
	}
	ids := convertToPartyIDs(peers)
	ids = append(ids, m.selfId)

	// Setup Run
	pl := pool.NewPool(0)
	defer pl.TearDown()

	// Build Session
	s, err := NewSession(m.selfNode, party.ID(m.selfNode.ID()), convertToPartyIDs(peers), kCmpKeygenFeed)
	if err != nil {
		fmt.Println("Error creating handler")
		return nil, err
	}

	// Run Protocol
	r, err := s.RunProtocol(cmp.Presign(ws.MPCConfig(), ids, pl), nil)
	if err != nil {
		fmt.Println("Error creating handler")
		return nil, err
	}
	preSig := r.(*ecdsa.PreSignature)
	return preSig, nil
}

// JoinCMPPreSign presigns a new signature
func (m *MpcProtocol) JoinCMPPreSignOnline(ws wallet.WalletShare, peers ...peer.ID) (*ecdsa.Signature, error) {
	if len(peers) == 0 {
		return nil, errors.New("no peers provided")
	}
	ids := convertToPartyIDs(peers)
	ids = append(ids, m.selfId)

	// Setup Run
	pl := pool.NewPool(0)
	defer pl.TearDown()

	// Build Session
	s, err := NewSession(m.selfNode, party.ID(m.selfNode.ID()), convertToPartyIDs(peers), kCmpKeygenFeed)
	if err != nil {
		fmt.Println("Error creating handler")
		return nil, err
	}

	// Run Protocol
	r, err := s.RunProtocol(cmp.Presign(ws.MPCConfig(), ids, pl), nil)
	if err != nil {
		fmt.Println("Error creating handler")
		return nil, err
	}

	sig := r.(*ecdsa.Signature)
	return sig, nil
}
