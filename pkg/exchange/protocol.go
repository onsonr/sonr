package exchange

import (
	"context"
	"errors"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/internet"
	"google.golang.org/protobuf/proto"
)

var (
	ErrParameters  = errors.New("Failed to create new ExchangeProtocol, invalid parameters")
	ErrInvalidPeer = errors.New("Peer object provided to ExchangeProtocol is Nil")
)

// ExchangeProtocol handles Global Sonr Exchange Protocol
type ExchangeProtocol struct {
	api.NodeImpl
	ctx      context.Context
	host     *host.SNRHost // host
	resolver internet.HDNSResolver
}

// NewProtocol creates new ExchangeProtocol
func NewProtocol(ctx context.Context, host *host.SNRHost, node api.NodeImpl) (*ExchangeProtocol, error) {
	// Check parameters
	if err := checkParams(host); err != nil {
		logger.Error("Failed to create ExchangeProtocol", err)
		return nil, err
	}

	// Create Exchange Protocol
	exchProtocol := &ExchangeProtocol{
		ctx:      ctx,
		host:     host,
		NodeImpl: node,
		resolver: internet.NewHDNSResolver(),
	}
	logger.Info("✅  ExchangeProtocol is Activated \n")
	return exchProtocol, nil
}

// FindPeerId method returns PeerID by SName
func (p *ExchangeProtocol) Query(sname string) (*common.PeerInfo, error) {
	// Get Peer from KadDHT store
	buf, err := p.host.GetValue(p.ctx, sname)
	if err != nil {
		logger.Error("Failed to get item from KadDHT", err)
		return nil, err
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
		logger.Error("Failed to get PeerInfo from Peer", err)
		return nil, err
	}

	// Verify Peer is registered
	ok, rec, err := p.Verify(sname)
	if err != nil {
		logger.Warn("Failed to verify Peer", err)
		return info, nil
	}

	// Update PeerInfo
	if ok {
		info.NameRecord = rec
		return info, nil
	}
	logger.Error("Peer is not registered", err)
	return info, err
}

// Update method updates peer instance in the store
func (p *ExchangeProtocol) Update(peer *common.Peer) error {
	// Create a cid manually by specifying the 'prefix' parameters
	key, err := peer.CID()
	if err != nil {
		return err
	}

	// Marshal Peer
	buf, err := peer.Buffer()
	if err != nil {
		logger.Error("Failed to Marshal Peer", err)
		return err
	}

	// Add Peer to KadDHT store
	err = p.host.PutValue(p.ctx, key, buf)
	if err != nil {
		logger.Error("Failed to put Item in KDHT", err)
		return err
	}
	return nil
}

// Verify method uses resolver to check if Peer is registered,
// returns true if Peer is registered
func (p *ExchangeProtocol) Verify(sname string) (bool, internet.Record, error) {
	// Create Context
	empty := internet.Record{}
	ctx, cancel := context.WithTimeout(p.ctx, time.Second*5)
	defer cancel()

	// Verify Peer is registered
	recs, err := p.resolver.LookupTXT(ctx, sname)
	if err != nil {
		logger.Error("Failed to resolve DNS record for SName", err)
		return false, empty, err
	}

	// Get Name Record
	rec, err := recs.GetNameRecord()
	if err != nil {
		logger.Error("Failed to get Name Record", err)
		return false, empty, err
	}

	// Check peer record
	pubKey, err := rec.PubKey()
	if err != nil {
		logger.Error("Failed to get public key from record", err)
		return false, rec, err
	}

	compId, err := peer.IDFromPublicKey(pubKey)
	if err != nil {
		logger.Error("Failed to extract PeerID from PublicKey", err)
		return false, rec, err
	}

	ok, err := compareRecordtoID(rec, compId)
	if err != nil {
		logger.Error("Failed to compare PeerID to record", err)
		return false, rec, err
	}
	return ok, rec, nil
}

func compareRecordtoID(r internet.Record, target peer.ID) (bool, error) {
	// Check peer record
	pid, err := r.PeerID()
	if err != nil {
		logger.Error("Failed to extract PeerID from PublicKey", err)
		return false, err
	}
	return pid == target, nil
}
