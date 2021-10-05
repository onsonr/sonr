package exchange

import (
	"context"
	"errors"
	"fmt"

	"time"

	psr "github.com/libp2p/go-libp2p-pubsub-router"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/net"
	"github.com/sonr-io/core/tools/state"
	"google.golang.org/protobuf/proto"
)

var (
	ErrParameters  = errors.New("Failed to create new ExchangeProtocol, invalid parameters")
	ErrInvalidPeer = errors.New("Peer object provided to ExchangeProtocol is Nil")
)

// ExchangeProtocol handles Global Sonr Exchange Protocol
type ExchangeProtocol struct {
	*psr.PubsubValueStore
	ctx      context.Context
	host     *host.SNRHost  // host
	emitter  *state.Emitter // Handle to signal when done
	resolver net.HDNSResolver
}

// NewProtocol creates new ExchangeProtocol
func NewProtocol(ctx context.Context, host *host.SNRHost, em *state.Emitter) (*ExchangeProtocol, error) {
	// Check parameters
	if err := checkParams(host, em); err != nil {
		return nil, logger.Error("Failed to create TransferProtocol", err)
	}

	// Create PubSub Value Store
	r, err := psr.NewPubsubValueStore(ctx, host.Host, host.Pubsub(), ExchangeValidator{}, psr.WithRebroadcastInterval(5*time.Second))
	if err != nil {
		return nil, logger.Error("Failed to create Exchange PubSubValueStore", err)
	}

	// Create Exchange Protocol
	exchProtocol := &ExchangeProtocol{
		ctx:              ctx,
		host:             host,
		emitter:          em,
		resolver:         net.NewHDNSResolver(),
		PubsubValueStore: r,
	}
	return exchProtocol, nil
}

// FindPeerId method returns PeerID by SName
func (p *ExchangeProtocol) Query(q *QueryRequest) (*common.PeerInfo, error) {
	query, val, err := q.QueryValue()
	if err != nil {
		return nil, logger.Error("Failed to Query Value", err)
	}

	// Find peer from sName in the store
	buf, err := p.PubsubValueStore.GetValue(p.ctx, query)
	if err != nil {
		return nil, logger.Error(fmt.Sprintf("Failed to GET peer (%s) from store, with Query Value: %s", val, query), err)
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
		return nil, logger.Error("Failed to get PeerInfo from Peer", err)
	}
	return info, nil
}

// Update method updates peer instance in the store
func (p *ExchangeProtocol) Update(peer *common.Peer) error {
	// Verify Peer is not nil
	if peer == nil {
		return ErrInvalidPeer
	}

	// Marshal Peer
	info, err := peer.Info()
	if err != nil {
		return logger.Error("Failed to get PeerInfo from Peer", err)
	}

	// Marshal Peer
	buf, err := proto.Marshal(peer)
	if err != nil {
		return logger.Error("Failed to Marshal Peer", err)
	}

	// Add Peer to SName Store
	err = p.PubsubValueStore.PutValue(p.ctx, info.StoreEntryKey, buf)
	if err != nil {
		return logger.Error("Failed to Put Value in Exchange Store", err)
	}
	return nil
}
