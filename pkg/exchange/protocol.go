package exchange

import (
	"context"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/beam"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/common"
	"google.golang.org/protobuf/proto"
)

var (
	logger             = golog.Child("protocols/exchange")
	ErrParameters      = errors.New("Failed to create new ExchangeProtocol, invalid parameters")
	ErrInvalidPeer     = errors.New("Peer object provided to ExchangeProtocol is Nil")
	ErrTopicNotCreated = errors.New("Lobby Topic has not been Created")
)

// ExchangeProtocol handles Global and Local Sonr Peer Exchange Protocol
type ExchangeProtocol struct {
	node       api.NodeImpl
	ctx        context.Context
	beamStore  beam.Beam
	host       *host.SNRHost // host
	authRecord api.Record
	nameRecord api.Record
	lobby      *Lobby
}

// NewProtocol creates new ExchangeProtocol
func NewProtocol(ctx context.Context, host *host.SNRHost, node api.NodeImpl, options ...Option) (*ExchangeProtocol, error) {
	// Set options
	opts := defaultOptions()
	for _, opt := range options {
		opt(opts)
	}

	// Create BeamStore
	b, err := beam.New(ctx, host, beam.ID("exchange"))
	if err != nil {
		return nil, err
	}

	// Create Exchange Protocol
	exchProtocol := &ExchangeProtocol{
		ctx:       ctx,
		beamStore: b,
		host:      host,
		node:      node,
	}
	logger.Debug("✅  ExchangeProtocol is Activated \n")

	// Set Peer in Exchange
	peer, err := node.Peer()
	if err != nil {
		logger.Errorf("%s - Failed to get Profile", err)
		return nil, err
	}
	exchProtocol.Put(peer)

	// Join Topic
	topic, err := host.Join(createOlc(opts.location))
	if err != nil {
		logger.Errorf("%s - Failed to create Lobby Topic", err)
		return nil, err
	}

	// Create Lobby
	if err := exchProtocol.initLobby(topic, opts); err != nil {
		logger.Errorf("%s - Failed to initialize Lobby", err)
		return nil, err
	}
	return exchProtocol, nil
}

// FindPeerId method returns PeerID by SName
func (p *ExchangeProtocol) Get(sname string) (*common.Peer, error) {
	peer := &common.Peer{}
	// Get Peer from KadDHT store
	if buf, err := p.beamStore.Get(sname); err == nil {
		// Unmarshal Peer
		err := proto.Unmarshal(buf, peer)
		if err != nil {
			logger.Errorf("%s - Failed to unmarshal Peer", err)
			return nil, err
		}
		return peer, nil
	} else {
		logger.Warn("Failed to get Peer from BeamStore: %s", err)
	}
	return p.Resolve(sname)
}

// Put method updates peer instance in the store
func (p *ExchangeProtocol) Put(peer *common.Peer) error {
	logger.Debug("Updating Peer in BeamStore")
	// Marshal Peer
	buf, err := peer.Buffer()
	if err != nil {
		logger.Errorf("Failed to Marshal Peer: %s", err)
		return err
	}

	// Add Peer to KadDHT store
	err = p.beamStore.Put(peer.GetSName(), buf)
	if err != nil {
		logger.Errorf("Failed to put item in BeamStore: %s", err)
		return err
	}
	return nil
}

// Resolve method resolves SName from DNS Table
func (p *ExchangeProtocol) Resolve(sname string) (*common.Peer, error) {
	logger.Debug("Attempting to resolve from DNS Table")
	// Get Peer from DNS Resolver
	recs, err := api.LookupTXT(p.ctx, sname)
	if err != nil {
		logger.Errorf("Failed to resolve DNS record for SName: %s", err)
		return nil, err
	}

	// Get Name Record
	rec, err := recs.GetNameRecord()
	if err != nil {
		logger.Errorf("Failed to get Name Record: %s", err)
		return nil, err
	}
	return rec.Peer()
}

// Verify method uses resolver to check if Peer is registered,
// returns true if Peer is registered
func (p *ExchangeProtocol) Verify(sname string) (bool, api.Record, error) {
	// Check if NamebaseClient is Nil
	empty := api.Record{}
	// Verify Peer is registered
	recs, err := api.LookupTXT(p.ctx, sname)
	if err != nil {
		logger.Errorf("Failed to resolve DNS record for SName: %s", err)
		return false, empty, err
	}

	// Get Name Record
	rec, err := recs.GetNameRecord()
	if err != nil {
		logger.Errorf("Failed to get Name Record: %s", err)
		return false, empty, err
	}

	// Check peer record
	pubKey, err := rec.PubKey()
	if err != nil {
		logger.Errorf("Failed to get public key from record: %s", err)
		return false, rec, err
	}

	compId, err := peer.IDFromPublicKey(pubKey)
	if err != nil {
		logger.Errorf("Failed to extract PeerID from PublicKey: %s", err)
		return false, rec, err
	}
	return rec.ComparePeerID(compId), rec, nil
}

// RegisterDomain registers a domain with Namebase.
func (p *ExchangeProtocol) Register(sName string, records ...api.Record) (api.DomainMap, error) {
	// Put records into Namebase

	ok, err := api.PutRecords(p.ctx, records...)
	if err != nil {
		logger.Error("Failed to Register SName", err)
		return nil, err
	}

	// API Call was Unsuccessful
	if !ok {
		return nil, err
	}

	// Get records from Namebase
	recs, err := api.FindRecords(p.ctx, sName)
	if err != nil {
		logger.Error("Failed to Find SName after Registering", err)
		return nil, err
	}

	// Map records to DomainMap
	m := make(api.DomainMap)
	for _, r := range recs {
		m[r.Host] = r.Value
	}
	return m, nil
}

// Update method publishes peer data to the topic
func (p *ExchangeProtocol) Update() error {
	// Verify Peer is not nil
	peer, err := p.node.Peer()
	if err != nil {
		return err
	}

	// Publish Event
	p.lobby.Publish(peer)
	return nil
}

// Close method closes the ExchangeProtocol
func (p *ExchangeProtocol) Close() error {
	p.lobby.eventHandler.Cancel()
	p.lobby.subscription.Cancel()
	err := p.lobby.topic.Close()
	if err != nil {
		logger.Errorf("%s - Failed to Close Local Lobby Topic for Exchange", err)
	}
	return nil
}
