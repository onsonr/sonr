package bind

import (
	"context"
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// * Struct: Reference for Binded Proxy Node * //
type nodeConfig struct {
	Ctx context.Context

	HasStarted      bool
	HasBootstrapped bool
	HasJoinedLocal  bool

	Status md.Status
}

func newNodeConfig() nodeConfig {
	return nodeConfig{
		Ctx: context.Background(),

		HasStarted:      false,
		HasBootstrapped: false,
		HasJoinedLocal:  false,

		Status: md.Status_NONE,
	}
}

func (mn *Node) contextNode() context.Context {
	return mn.config.Ctx
}

func (mn *Node) contextUser() context.Context {
	return mn.config.Ctx
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
		sentry.CaptureException(errors.Wrap(err, "Unmarshalling StatusUpdate Protobuf"))
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
		sentry.CaptureException(errors.Wrap(err, "Unmarshalling StatusUpdate Protobuf"))
		return
	}
	mn.call.OnStatus(data)
}
