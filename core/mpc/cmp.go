package mpc

import (
	"errors"
	"sync"

	"github.com/sonr-hq/sonr/internal/node"
	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"

	p2p_protocol "github.com/libp2p/go-libp2p/core/protocol"
)

const (
	// MPC_KEYGEN_PROTOCOL is the protocol ID for the MPC keygen protocol that is attached to the node.
	kCmpKeygenRequest = p2p_protocol.ID("/mpc-cmp/keygen-request/0.1.0")

	// MPC_KEYGEN_PROTOCOL is the protocol ID for the MPC keygen protocol that is attached to the node.
	kCmpKeygenResponse = p2p_protocol.ID("/mpc-cmp/keygen-response/0.1.0")

	// MPC_KEYGEN_FEED_PROTOCOL
	kCmpKeygenFeed = p2p_protocol.ID("/mpc-cmp/keygen-feed/0.1.0")

	// MPC_SIGN_PROTOCOL is the protocol ID for the MPC sign protocol that is attached to the node.
	kCmpSign = p2p_protocol.ID("/mpc-cmp/sign/0.1.0")

	// MPC_REFRESH_PROTOCOL is the protocol ID for the MPC refresh protocol that is attached to the node.
	kCmpRefresh = p2p_protocol.ID("/mpc-cmp/refresh/0.1.0")

	// MPC_PRE_SIGN_PROTOCOL is the protocol ID for the MPC pre-sign protocol that is attached to the node.
	kCmpPreSign = p2p_protocol.ID("/mpc-cmp/pre-sign/0.1.0")

	// MPC_PRE_SIGN_ONLINE_PROTOCOL is the protocol ID for the MPC pre-sign online protocol that is attached to the node.
	kCmpPreSignOnline = p2p_protocol.ID("/mpc-cmp/pre-sign-online/0.1.0")
)

func CmpKeygen(id party.ID, ids party.IDSlice, topicHandler node.TopicHandler, threshold int, wg *sync.WaitGroup, pl *pool.Pool) (*cmp.Config, error) {
	tph, err := protocol.NewMultiHandler(cmp.Keygen(curve.Secp256k1{}, id, ids, 1, pl), []byte(topicHandler.Name()))
	if err != nil {
		return nil, err
	}
	// handlerLoopTopic(id, tph, topicHandler)
	r, err := tph.Result()
	if err != nil {
		return nil, err
	}
	return r.(*cmp.Config), nil
}

// func cmpRefresh(c *cmp.Config, topicHandler node.TopicHandler, wg *sync.WaitGroup, pl *pool.Pool) (*cmp.Config, error) {

// 	handlerLoopChannel(c.ID, h, n)
// 	r, err := h.Result()
// 	if err != nil {
// 		return nil, err
// 	}
// 	conf := r.(*cmp.Config)
// 	return conf, nil
// }

func CmpSign(c *cmp.Config, m []byte, signers party.IDSlice, topicHandler node.TopicHandler, wg *sync.WaitGroup, pl *pool.Pool) (*ecdsa.Signature, error) {
	defer wg.Done()
	h, err := protocol.NewMultiHandler(cmp.Sign(c, signers, m, pl), []byte(topicHandler.Name()))
	if err != nil {
		return nil, err
	}
	// handlerLoopTopic(c.ID, h, topicHandler)

	signResult, err := h.Result()
	if err != nil {
		return nil, err
	}
	signature := signResult.(*ecdsa.Signature)
	if !signature.Verify(c.PublicPoint(), m) {
		return nil, errors.New("failed to verify cmp signature")
	}
	return signature, nil
}

func CmpPreSign(c *cmp.Config, signers party.IDSlice, topicHandler node.TopicHandler, wg *sync.WaitGroup, pl *pool.Pool) (*ecdsa.PreSignature, error) {
	defer wg.Done()
	h, err := protocol.NewMultiHandler(cmp.Presign(c, signers, pl), []byte(topicHandler.Name()))
	if err != nil {
		return nil, err
	}
	// handlerLoopTopic(c.ID, h, topicHandler)
	signResult, err := h.Result()
	if err != nil {
		return nil, err
	}
	preSignature := signResult.(*ecdsa.PreSignature)
	if err = preSignature.Validate(); err != nil {
		return nil, errors.New("failed to verify cmp presignature")
	}
	return preSignature, nil
}

func CmpPreSignOnline(c *cmp.Config, preSignature *ecdsa.PreSignature, m []byte, topicHandler node.TopicHandler, wg *sync.WaitGroup, pl *pool.Pool) (*ecdsa.Signature, error) {
	defer wg.Done()
	h, err := protocol.NewMultiHandler(cmp.PresignOnline(c, preSignature, m, pl), []byte(topicHandler.Name()))
	if err != nil {
		return nil, err
	}
	// handlerLoopTopic(c.ID, h, topicHandler)

	signResult, err := h.Result()
	if err != nil {
		return nil, err
	}
	signature := signResult.(*ecdsa.Signature)
	if !signature.Verify(c.PublicPoint(), m) {
		return nil, errors.New("failed to verify cmp signature")
	}
	return signature, nil
}
