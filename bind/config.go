package bind

import (
	"context"
	"log"

	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// * Struct: Reference for Binded Proxy Node * //
type mobileConfig struct {
	CtxFS   context.Context
	CtxNode context.Context
	CtxUser context.Context

	HasStarted      bool
	HasBootstrapped bool
	HasJoinedLocal  bool

	Status md.Status
}

func newMobileConfig() mobileConfig {
	return mobileConfig{
		CtxFS:   context.Background(),
		CtxNode: context.Background(),
		CtxUser: context.Background(),

		HasStarted:      false,
		HasBootstrapped: false,
		HasJoinedLocal:  false,

		Status: md.Status_NONE,
	}
}

func (mn *Node) contextFS() context.Context {
	return mn.config.CtxFS
}

func (mn *Node) contextNode() context.Context {
	return mn.config.CtxNode
}

func (mn *Node) contextUser() context.Context {
	return mn.config.CtxUser
}

func (mn *Node) isReady() bool {
	return mn.config.HasBootstrapped && mn.config.HasStarted
}

func (mn *Node) setConnected(val bool) {
	// Update Status
	mn.config.HasStarted = val
	mn.config.Status = md.Status_CONNECTED

	// Build Update
	m := &md.StatusUpdate{
		Value: mn.config.Status,
	}

	// Callback Status
	data, err := proto.Marshal(m)
	if err != nil {
		log.Println(err)
		return
	}
	mn.call.OnStatus(data)
}

func (mn *Node) setBootstrapped(val bool) {
	// Update Status
	mn.config.HasBootstrapped = val
	mn.config.Status = md.Status_BOOTSTRAPPED

	// Build Update
	m := &md.StatusUpdate{
		Value: mn.config.Status,
	}

	// Callback Status
	data, err := proto.Marshal(m)
	if err != nil {
		log.Println(err)
		return
	}
	mn.call.OnStatus(data)
}

func (mn *Node) setJoinedLocal(val bool) {
	// Update Status
	mn.config.HasJoinedLocal = val
	mn.config.Status = md.Status_AVAILABLE

	// Build Update
	m := &md.StatusUpdate{
		Value: mn.config.Status,
	}

	// Callback Status
	data, err := proto.Marshal(m)
	if err != nil {
		log.Println(err)
		return
	}
	mn.call.OnStatus(data)
}

func (mn *Node) setStatus(newStatus md.Status) {
	// Set Status
	mn.config.Status = newStatus

	// Build Update
	m := &md.StatusUpdate{
		Value: mn.config.Status,
	}

	// Callback Status
	data, err := proto.Marshal(m)
	if err != nil {
		log.Println(err)
		return
	}
	mn.call.OnStatus(data)
}
