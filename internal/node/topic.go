package node

import (
	"context"

	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/kubo/core"
)

type TopicHandler interface {
	Name() string
	ListPeers() ([]string, error)
	Publish(data []byte) error
	Messages() <-chan []byte
	Close()
}

type handler struct {
	sub      icore.PubSubSubscription
	name     string
	ctx      context.Context
	messages chan []byte
	p2p      *core.IpfsNode
}

func startHandler(ctx context.Context, p2p *core.IpfsNode, sub icore.PubSubSubscription, name string) TopicHandler {
	h := &handler{
		sub:      sub,
		name:     name,
		messages: make(chan []byte),
		ctx:      ctx,
		p2p:      p2p,
	}
	go h.handle()
	return h
}

func (h *handler) handle() {
	for {
		msg, err := h.sub.Next(h.ctx)
		if err != nil {
			return
		}
		h.messages <- msg.Data()
	}
}

func (h *handler) Messages() <-chan []byte {
	return h.messages
}

func (h *handler) Close() {
	h.sub.Close()
}

func (h *handler) Publish(data []byte) error {
	return h.p2p.PubSub.Publish(h.name, data)
}

func (h *handler) ListPeers() ([]string, error) {
	peers := h.p2p.PubSub.ListPeers(h.name)
	peerIdstrs := make([]string, 0)
	for _, peer := range peers {
		peerIdstrs = append(peerIdstrs, peer.String())
	}
	return peerIdstrs, nil
}

func (h *handler) Name() string {
	return h.name
}
