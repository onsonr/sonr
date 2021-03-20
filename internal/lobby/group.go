package lobby

import (
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/pkg/models"
	net "github.com/sonr-io/core/pkg/net"
)

type Group struct {
	Name         string
	Size         int
	Members      map[string]*md.Peer
	Admin        md.Peer
	subscription *pubsub.Subscription
	topic        *pubsub.Topic
	data         *md.Group
}

// ^ Creates Group for
func (lob *Lobby) NewGroup(group string) (*Group, error) {
	// Join the local pubsub Topic
	topic, err := lob.pubSub.Join(lob.router.Topic(net.SetIDForGroup(group)))
	if err != nil {
		return nil, err
	}

	// Subscribe to local Topic
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	// Return Group
	return &Group{
		Name:         group,
		topic:        topic,
		subscription: sub,
		data: &md.Group{
			Name:    group,
			Size:    1,
			Members: make(map[string]*md.Peer),
			Admin:   lob.selfPeer,
		},
	}, nil
}
