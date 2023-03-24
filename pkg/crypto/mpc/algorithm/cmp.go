package algorithm

import (
	"errors"
	"sync"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// It creates a new handler for the keygen protocol, runs the handler loop, and returns the result
func CmpKeygen(id party.ID, ids party.IDSlice, n crypto.Network, threshold int, wg *sync.WaitGroup, pl *pool.Pool) (*cmp.Config, error) {
	defer wg.Done()
	h, err := protocol.NewMultiHandler(cmp.Keygen(curve.Secp256k1{}, id, ids, threshold, pl), nil)
	if err != nil {
		return nil, err
	}

	HandleNetworkProtocol(id, h, n)
	r, err := h.Result()
	if err != nil {
		return nil, err
	}
	conf := r.(*cmp.Config)
	return conf, nil
}

// It creates a new handler for the refresh protocol, runs the handler loop, and returns the result
func CmpRefresh(c *cmp.Config, n crypto.Network, wg *sync.WaitGroup, pl *pool.Pool) (*cmp.Config, error) {
	defer wg.Done()
	h, err := protocol.NewMultiHandler(cmp.Refresh(c, pl), nil)
	if err != nil {
		return nil, err
	}

	HandleNetworkProtocol(c.ID, h, n)
	r, err := h.Result()
	if err != nil {
		return nil, err
	}
	conf := r.(*cmp.Config)
	return conf, nil
}

// It creates a new `protocol.MultiHandler` for the `cmp.Sign` protocol, and then runs the handler loop
func CmpSign(c *cmp.Config, m []byte, signers party.IDSlice, n crypto.Network, wg *sync.WaitGroup, pl *pool.Pool) (*crypto.MPCECDSASignature, error) {
	defer wg.Done()
	h, err := protocol.NewMultiHandler(cmp.Sign(c, signers, m, pl), nil)
	if err != nil {
		return nil, err
	}
	HandleNetworkProtocol(c.ID, h, n)

	r, err := h.Result()
	if err != nil {
		return nil, err
	}
	sig, ok := r.(*ecdsa.Signature)
	if !ok {
		return nil, errors.New("failed to cast result to ecdsa.Signature")
	}
	if !sig.Verify(c.PublicPoint(), m) {
		return nil, errors.New("failed to verify cmp signature")
	}
	return sig, nil
}

// CmpVerify verifies a signature using the public key of the signer
func CmpVerify(c *cmp.Config, m []byte, sig *crypto.MPCECDSASignature) bool {
	return sig.Verify(c.PublicPoint(), m)
}
