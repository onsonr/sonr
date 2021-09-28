package exchange

import (
	"context"
	"fmt"
	"time"

	psr "github.com/libp2p/go-libp2p-pubsub-router"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

// ExchangeProtocol handles Global Sonr Exchange Protocol
type ExchangeProtocol struct {
	*psr.PubsubValueStore
	ctx     context.Context
	host    *host.SNRHost  // host
	emitter *state.Emitter // Handle to signal when done
}

// NewProtocol creates new ExchangeProtocol
func NewProtocol(ctx context.Context, host *host.SNRHost, em *state.Emitter) (*ExchangeProtocol, error) {
	// Create PubSub Value Store
	r, err := psr.NewPubsubValueStore(ctx, host.Host, host.Pubsub(), ExchangeValidator{}, psr.WithRebroadcastInterval(5*time.Second))
	if err != nil {
		return nil, err
	}

	// Create Exchange Protocol
	exchProtocol := &ExchangeProtocol{
		ctx:              ctx,
		host:             host,
		emitter:          em,
		PubsubValueStore: r,
	}
	return exchProtocol, nil
}

// FindPeerId method returns PeerID by SName
func (p *ExchangeProtocol) Query(q *QueryRequest) (*common.PeerInfo, error) {
	query, val, err := q.QueryValue()
	if err != nil {
		logger.Error("Failed to Query Value", zap.Error(err))
		return nil, err
	}

	// Find peer from sName in the store
	buf, err := p.PubsubValueStore.GetValue(p.ctx, query)
	if err != nil {
		msg := fmt.Sprintf("Failed to GET peer (%s) from store, with Query Value: %s", val, query)
		logger.Error(msg, zap.Error(err))
		return nil, errors.Wrap(err, msg)
	}

	// Unmarshal Peer from buffer
	peerData := &common.Peer{}
	err = proto.Unmarshal(buf, peerData)
	if err != nil {
		return nil, err
	}

	// Get PeerID from Peer
	info, err := peerData.Info()
	if err != nil {
		msg := fmt.Sprintf("Failed to get PeerInfo from Peer: %s", val)
		logger.Error(msg, zap.Error(err))
		return nil, errors.Wrap(err, msg)
	}
	return info, nil
}

// Update method updates peer instance in the store
func (p *ExchangeProtocol) Update(peer *common.Peer) error {
	// Marshal Peer
	info, err := peer.Info()
	if err != nil {
		logger.Error("Failed to get PeerInfo from Peer", zap.Error(err))
		return err
	}

	// Marshal Peer
	buf, err := proto.Marshal(peer)
	if err != nil {
		logger.Error("Failed to Marshal Peer", zap.Error(err))
		return err
	}

	// Add Peer to SName Store
	err = p.PubsubValueStore.PutValue(p.ctx, info.StoreEntryKey, buf)
	if err != nil {
		msg := fmt.Sprintf("Failed to Add Peer Object to SName store: %s", peer.GetSName())
		logger.Error(msg, zap.Error(err))
		return errors.Wrap(err, msg)
	}
	return nil
}
