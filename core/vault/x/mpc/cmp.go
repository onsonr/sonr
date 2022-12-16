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
	"github.com/sonr-io/multi-party-sig/protocols/doerner"
)

func CmpKeygen(id party.ID, ids party.IDSlice, topicHandler node.TopicHandler, threshold int, wg *sync.WaitGroup, pl *pool.Pool) (wallet.WalletShare, error) {
	tph, err := protocol.NewTwoPartyHandler(doerner.Keygen(curve.Secp256k1{}, false, id, ids[0], pl), []byte(topicHandler.Name()), true)
	if err != nil {
		return nil, err
	}
	handlerLoopTopic(id, tph, topicHandler)
	r, err := tph.Result()
	if err != nil {
		return nil, err
	}
	return &mpcConfigWalletImpl{r.(*cmp.Config)}, nil
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
