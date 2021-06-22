package host

import (
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	psub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/pkg/models"
)

// ** ─── Stream/Pubsub Methods ────────────────────────────────────────────────────────
// Join New Topic with Name
func (h *hostNode) Join(name string) (*psub.Topic, *psub.Subscription, *psub.TopicEventHandler, *md.SonrError) {
	// Join Topic
	topic, err := h.pubsub.Join(name)
	if err != nil {
		return nil, nil, nil, md.NewError(err, md.ErrorMessage_TOPIC_JOIN)
	}

	// Subscribe to Topic
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, nil, nil, md.NewError(err, md.ErrorMessage_TOPIC_SUB)
	}

	// Create Topic Handler
	handler, err := topic.EventHandler()
	if err != nil {
		return nil, nil, nil, md.NewError(err, md.ErrorMessage_TOPIC_HANDLER)
	}
	return topic, sub, handler, nil
}

// Set Stream Handler for Host
func (h *hostNode) HandleStream(pid protocol.ID, handler network.StreamHandler) {
	h.host.SetStreamHandler(pid, handler)
}

// Start Stream for Host
func (h *hostNode) StartStream(p peer.ID, pid protocol.ID) (network.Stream, error) {
	return h.host.NewStream(h.ctxHost, p, pid)
}
