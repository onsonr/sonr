package bind

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	ath "github.com/sonr-io/core/internal/auth"
	tpc "github.com/sonr-io/core/internal/topic"
	sc "github.com/sonr-io/core/pkg/client"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// * Struct: Reference for Binded Proxy Node * //
type Node struct {
	md.NodeCallback

	// Properties
	call    Callback
	ctx     context.Context
	request *md.ConnectionRequest

	// Client
	auth   ath.AuthService
	client *sc.Client
	user   *md.User

	// Groups
	local  *tpc.TopicManager
	topics map[string]*tpc.TopicManager

	// Miscellaneous
	initialized    bool
	storageEnabled bool
	store          md.Store
}

// @ Initializes New Node
func NewNode(reqBytes []byte, call Callback) *Node {
	// Initialize Sentry
	sentry.Init(sentry.ClientOptions{
		Dsn: "https://cbf88b01a5a5468fa77101f7dfc54f20@o549479.ingest.sentry.io/5672329",
	})

	// Unmarshal Request
	req := &md.ConnectionRequest{}
	err := proto.Unmarshal(reqBytes, req)
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "Unmarshalling Connection Request"))
		return nil
	}
	// Initialize Node
	mn := &Node{
		call:        call,
		ctx:         context.Background(),
		initialized: false,
		topics:      make(map[string]*tpc.TopicManager, 10),
	}
	mn.initialize(req)
	return mn
}

// **-------------------** //
// ** LifeCycle Actions ** //
// **-------------------** //
// @ Close Ends All Network Communication
func (mn *Node) Pause() {
	md.GetState().Pause()
}

// @ Close Ends All Network Communication
func (mn *Node) Resume() {
	md.GetState().Resume()
}

// @ Close Ends All Network Communication
func (mn *Node) Stop() {
	mn.client.Close()
	mn.ctx.Done()
}
