package mpc

import (
	"errors"
	"sync"

	"github.com/sonr-hq/sonr/internal/node"
	"github.com/sonr-hq/sonr/pkg/wallet"
	"github.com/sonr-io/multi-party-sig/pkg/ecdsa"
	"github.com/sonr-io/multi-party-sig/pkg/math/curve"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/multi-party-sig/pkg/pool"
	"github.com/sonr-io/multi-party-sig/pkg/protocol"
	"github.com/sonr-io/multi-party-sig/protocols/cmp"
)

func CmpKeygen(id party.ID, ids party.IDSlice, topicHandler node.TopicHandler, threshold int, wg *sync.WaitGroup, pl *pool.Pool) (wallet.WalletShare, error) {
	defer wg.Done()
	h, err := protocol.NewMultiHandler(cmp.Keygen(curve.Secp256k1{}, id, ids, threshold, pl), []byte(topicHandler.Name()))
	if err != nil {
		return nil, err
	}

	handlerLoopTopic(id, h, topicHandler)
	r, err := h.Result()
	if err != nil {
		return nil, err
	}
	conf := r.(*cmp.Config)
	//topic := fmt.Sprintf("/sonr/v0.2.0/mpc/sign/%s-%s", w.Config.ID, searchFirstNotId(w.Config.PartyIDs(), w.Config.ID))
	return &mpcConfigWalletImpl{conf}, nil
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
	handlerLoopTopic(c.ID, h, topicHandler)

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
	handlerLoopTopic(c.ID, h, topicHandler)
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
	handlerLoopTopic(c.ID, h, topicHandler)

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
