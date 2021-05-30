package bind

import (
	"context"
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"

	sc "github.com/sonr-io/core/pkg/client"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

type LinkerNode struct {
	// Properties
	call Callback
	ctx  context.Context

	// Client
	client *sc.Client
	device *md.Device
}

// @ Create New Mobile Node
func NewLinker(reqBytes []byte, call Callback) *LinkerNode {
	// Initialize Sentry
	sentry.Init(sentry.ClientOptions{
		Dsn: "https://cbf88b01a5a5468fa77101f7dfc54f20@o549479.ingest.sentry.io/5672329",
	})

	// Unmarshal Request
	req := &md.LinkRequest{}
	err := proto.Unmarshal(reqBytes, req)
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "Unmarshalling Connection Request"))
		return nil
	}
	// Create Mobile Node
	mn := &LinkerNode{
		call:   call,
		ctx:    context.Background(),
		device: req.GetDevice(),
	}

	// Create Client
	c, serr := sc.NewLinkClient(mn.ctx, req)
	if serr != nil {
		log.Println(err)
		return nil
	}

	// Set client
	mn.client = c
	return mn
}
