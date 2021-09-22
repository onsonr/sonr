package exchange

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	psr "github.com/libp2p/go-libp2p-pubsub-router"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/emitter"
	"github.com/sonr-io/core/tools/logger"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

// TransferProtocol type
type ExchangeProtocol struct {
	*psr.PubsubValueStore
	ctx     context.Context
	host    *host.SNRHost    // host
	emitter *emitter.Emitter // Handle to signal when done
}

// NewProtocol creates new ExchangeProtocol
func NewProtocol(ctx context.Context, host *host.SNRHost, em *emitter.Emitter) (*ExchangeProtocol, error) {
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

	// Handle Background Processes, return Protocol
	// go exchProtocol.HandleEvents()
	// go exchProtocol.HandleMessages()
	return exchProtocol, nil
}

// Find method returns PeerID by SName
func (p *ExchangeProtocol) Find(sName string) (peer.ID, error) {
	// Set Lowercase Name
	sName = strings.ToLower(sName)

	// Find peer from sName in the store
	buf, err := p.PubsubValueStore.GetValue(p.ctx, fmt.Sprintf("store/%s", sName))
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to GET peer (%s) from store", sName), zap.Error(err))
		return "", err
	}

	// Unmarshal Peer from buffer
	profile := &common.Peer{}
	err = proto.Unmarshal(buf, profile)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to Unmarshal Peer (%s)", sName), zap.Error(err))
		return "", err
	}

	// Fetch public key from peer data
	pubKey, err := crypto.UnmarshalPublicKey(profile.PublicKey)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to Unmarshal Public Key for (%s)", sName), zap.Error(err))
		return "", err
	}

	// Get peer ID from public key
	id, err := peer.IDFromPublicKey(pubKey)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to get peer ID from Public Key for (%s)", sName), zap.Error(err))
		return "", err
	}
	return id, nil
}

// Ping method finds peer Profile by name
func (p *ExchangeProtocol) Ping(sName string) (*common.Peer, error) {
	// Set Lowercase Name
	sName = strings.ToLower(sName)

	// Find peer from sName in the store
	buf, err := p.PubsubValueStore.GetValue(p.ctx, fmt.Sprintf("store/%s", sName))
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to GET peer (%s) from store", sName), zap.Error(err))
		return nil, err
	}

	// Unmarshal Peer from buffer
	profile := &common.Peer{}
	err = proto.Unmarshal(buf, profile)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to Unmarshal Peer (%s)", sName), zap.Error(err))
		return nil, err
	}
	return profile, nil
}

// Update method updates peer instance in the store
func (p *ExchangeProtocol) Update(peer *common.Peer) error {
	// Marshal Peer
	peerBuf, err := proto.Marshal(peer)
	if err != nil {
		logger.Error("Failed to Marshal Peer", zap.Error(err))
		return err
	}

	// Determine Key and Add Value to Store
	err = p.PubsubValueStore.PutValue(p.ctx, fmt.Sprintf("store/%s", strings.ToLower(peer.GetSName())), peerBuf)
	if err != nil {
		logger.Error("Failed to PUT peer from store", zap.Error(err))
		return err
	}
	return nil
}
