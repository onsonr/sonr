package local

import (
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/pkg/models"
)

// @ Helper: Finds Peer in Lobby
func (lm *LocalManager) findPeer(q string) (*md.Peer, error) {
	// Iterate Through Peers, Return Matched Peer
	val, ok := lm.lobby.Find(q)
	if ok {
		return val, nil
	}
	return nil, errors.New("Peer Data not found in Lobby.")
}

// @ Helper: Joins IP Topic - Default
func (lm *LocalManager) joinIPTopic() (*pubsub.Subscription, *pubsub.TopicEventHandler, *md.SonrError) {
	// Join IP Topic
	ipTopic, ipSub, ipHandler, serr := lm.host.Join(lm.user.LocalIPTopic())
	if serr != nil {
		return nil, nil, serr
	}

	// Set IP Topic Return Handlers
	lm.ipTopic = ipTopic
	return ipSub, ipHandler, nil
}

// @ Helper: Joins Geo Topic - Mobile
func (lm *LocalManager) joinGeoTopic() (*pubsub.Subscription, *pubsub.TopicEventHandler, *md.SonrError) {
	// Find GeoTopic Name
	geoName, err := lm.user.LocalGeoTopic()
	if err != nil {
		return nil, nil, md.NewError(err, md.ErrorMessage_TOPIC_INVALID)
	}

	// Join Geo Topic
	geoTopic, geoSub, geoHandler, serr := lm.host.Join(geoName)
	if serr != nil {
		return nil, nil, serr
	}

	// Set GeoTopic Return Handlers
	lm.geoTopic = geoTopic
	return geoSub, geoHandler, nil
}

// @ Helper: Registers Local Service
func (lm *LocalManager) registerService() *md.SonrError {
	// Start Exchange Server
	peersvServer := rpc.NewServer(lm.host.Host, K_SERVICE_PID)
	lsv := TopicService{
		lobby:  lm.lobby,
		user:   lm.user,
		call:   lm.callback,
		respCh: make(chan *md.AuthReply, 1),
		linkCh: make(chan *md.LinkResponse, 1),
	}

	// Register Service
	err := peersvServer.RegisterName(K_RPC_SERVICE, &lsv)
	if err != nil {
		return md.NewError(err, md.ErrorMessage_TOPIC_RPC)
	}

	// Set Service
	lm.service = &lsv
	return nil
}

// @ Helper: Searches IPTopic for PeerID
func (lm *LocalManager) searchIPTopic(q string) (peer.ID, error) {
	// Iterate through Topic Peers
	for _, id := range lm.ipTopic.ListPeers() {
		// If Found Match
		if id.String() == q {
			return id, nil
		}
	}
	return "", errors.New("Peer ID was not found in topic.")
}

// @ Helper: Searches GeoTopic for PeerID
func (lm *LocalManager) searchGeoTopic(q string) (peer.ID, error) {
	if lm.geoTopic != nil {
		// Iterate through Topic Peers
		for _, id := range lm.geoTopic.ListPeers() {
			// If Found Match
			if id.String() == q {
				return id, nil
			}
		}
	}
	return "", errors.New("Peer ID was not found in topic.")
}
