package mpc

import (
	"errors"
	"sync"

	"github.com/sonr-hq/sonr/internal/node"
	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

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
