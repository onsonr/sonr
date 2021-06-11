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
	reqType md.ConnectionRequest_Type

	// Client
	auth   ath.AuthService
	client *sc.Client
	user   *md.User

	// Groups
	local  *tpc.TopicManager
	topics map[string]*tpc.TopicManager

	// Storage
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
		call:   call,
		ctx:    context.Background(),
		topics: make(map[string]*tpc.TopicManager, 10),
	}
	mn.initialize(req)
	return mn
}

// **-----------------** //
// ** Network Actions ** //
// **-----------------** //
// @ Start Host and Connect
func (mn *Node) Connect() []byte {
	// Connect Host
	err := mn.client.Connect(mn.user.KeyPrivate())
	if err != nil {
		mn.handleError(err)
		mn.setConnected(false)
		return nil
	} else {
		// Update Status
		mn.setConnected(true)
	}

	// Bootstrap Node
	mn.local, err = mn.client.Bootstrap()
	if err != nil {
		mn.handleError(err)
		mn.setAvailable(false)
		return nil
	} else {
		mn.setAvailable(true)
	}

	// Create ConnectResponse
	bytes, rerr := proto.Marshal(&md.ConnectionResponse{
		User: mn.user,
		Id:   mn.user.ID(),
	})

	// Handle Error
	if rerr != nil {
		mn.handleError(md.NewMarshalError(rerr))
		return nil
	}
	return bytes
}

// @ Returns Node Location Protobuf as Bytes
func (mn *Node) Location() []byte {
	bytes, err := proto.Marshal(mn.user.Location)
	if err != nil {
		return nil
	}
	return bytes
}

// @ Returns Node User Protobuf as Bytes
func (mn *Node) User() []byte {
	bytes, err := proto.Marshal(mn.user)
	if err != nil {
		return nil
	}
	return bytes
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
