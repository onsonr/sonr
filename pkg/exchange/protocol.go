package exchange

import (
	"context"
	"errors"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
	"google.golang.org/protobuf/proto"
)

var (
	logger         = golog.Child("protocols/exchange")
	ErrParameters  = errors.New("Failed to create new ExchangeProtocol, invalid parameters")
	ErrInvalidPeer = errors.New("Peer object provided to ExchangeProtocol is Nil")
)

// ExchangeProtocol handles Global Sonr Exchange Protocol
type ExchangeProtocol struct {
	node           api.NodeImpl
	ctx            context.Context
	host           *host.SNRHost // host
	namebaseClient *NamebaseAPIClient
}

// NewProtocol creates new ExchangeProtocol
func NewProtocol(ctx context.Context, host *host.SNRHost, node api.NodeImpl) (*ExchangeProtocol, error) {
	key, secret, err := fetchApiKeys()
	if err != nil {
		logger.Error("Failed to fetch API Keys", err)
		return nil, err
	}

	// Create Exchange Protocol
	exchProtocol := &ExchangeProtocol{
		ctx:      ctx,
		host:     host,
		node:     node,
		namebaseClient: NewNamebaseClient(ctx, key, secret),
	}
	logger.Debug("✅  ExchangeProtocol is Activated \n")

	// Set Peer in Exchange
	peer, err := node.Peer()
	if err != nil {
		logger.Error("Failed to get Profile", err)
		return nil, err
	}
	exchProtocol.Put(peer)
	return exchProtocol, nil
}

// FindPeerId method returns PeerID by SName
func (p *ExchangeProtocol) Get(sname string) (*common.PeerInfo, error) {
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
	ok, _, err := p.Verify(sname)
	if err != nil {
		logger.Warn("Failed to verify Peer", err)
		return info, nil
	}

	// Update PeerInfo
	if ok {
		return info, nil
	}
	logger.Error("Peer is not registered", err)
	return info, err
}

// Put method updates peer instance in the store
func (p *ExchangeProtocol) Put(peer *common.Peer) error {
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
func (p *ExchangeProtocol) Verify(sname string) (bool, host.Record, error) {
	// Create Context
	empty := host.Record{}

	// Verify Peer is registered
	recs, err := p.host.LookupTXT(p.ctx, sname)
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

// RegisterDomain registers a domain with Namebase.
func (p *ExchangeProtocol) Register(sName string, records ...host.Record) (host.DomainMap, error) {
	// Put records into Namebase
	req := host.NewNBAddRequest(records...)
	ok, err := p.namebaseClient.PutRecords(req)
	if err != nil {
		logger.Error("Failed to Register SName", err)
		return nil, err
	}

	// API Call was Unsuccessful
	if !ok {
		return nil, err
	}

	// Get records from Namebase
	recs, err := p.namebaseClient.FindRecords(sName)
	if err != nil {
		logger.Error("Failed to Find SName after Registering", err)
		return nil, err
	}

	// Map records to DomainMap
	m := make(host.DomainMap)
	for _, r := range recs {
		m[r.Host] = r.Value
	}
	return m, nil
}

func compareRecordtoID(r host.Record, target peer.ID) (bool, error) {
	// Check peer record
	pid, err := r.PeerID()
	if err != nil {
		logger.Error("Failed to extract PeerID from PublicKey", err)
		return false, err
	}
	return pid == target, nil
}
