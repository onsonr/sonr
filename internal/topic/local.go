package topic

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	net "github.com/sonr-io/core/internal/host"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

// ExchangeArgs is Peer protobuf
type LocalServiceArgs struct {
	Peer   []byte
	Invite []byte
}

// ExchangeResponse is also Peer protobuf
type LocalServiceResponse struct {
	InvReply []byte
	Peer     []byte
}

// Service Struct
type LocalService struct {
	// Current Data
	call ClientHandler
	user *md.User

	respCh chan *md.InviteResponse
	invite *md.InviteRequest
}

// ^ Create New Contained Topic Manager ^ //
func NewLocal(ctx context.Context, h net.HostNode, u *md.User, name string, th ClientHandler) (*TopicManager, *md.SonrError) {
	// Join Topic
	topic, sub, handler, serr := h.Join(name)
	if serr != nil {
		return nil, serr
	}

	// Create Lobby Manager
	mgr := &TopicManager{
		handler:      th,
		user:         u,
		ctx:          ctx,
		host:         h,
		eventHandler: handler,
		lobbyType:    md.Lobby_LOCAL,
		localEvents:  make(chan *md.LocalEvent, util.TOPIC_MAX_MESSAGES),
		subscription: sub,
		topic:        topic,
	}

	// Start Exchange Server
	localServer := rpc.NewServer(h.Host(), util.LOCAL_PROTOCOL)
	psv := LocalService{
		user:   u,
		call:   th,
		respCh: make(chan *md.InviteResponse, util.TOPIC_MAX_MESSAGES),
	}

	// Register Service
	err := localServer.RegisterName(util.LOCAL_RPC_SERVICE, &psv)
	if err != nil {
		return nil, md.NewError(err, md.ErrorMessage_TOPIC_RPC)
	}

	// Set Service
	mgr.service = &psv
	go mgr.handleTopicEvents(context.Background())
	go mgr.handleTopicMessages(context.Background())
	go mgr.processTopicMessages(context.Background())
	return mgr, nil
}

// @ Send Updated Lobby
func (tm *TopicManager) RefreshLobby(event *md.LocalEvent) {
	tm.handler.OnEvent(event)
}

// @ SendLocal message to specific peer in topic
func (tm *TopicManager) SendLocal(msg *md.LocalEvent) error {
	// Convert Event to Proto Binary
	bytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	// Publish to Topic
	err = tm.topic.Publish(tm.ctx, bytes)
	if err != nil {
		return err
	}
	return nil
}

// @ Starts Exchange on Local Peer Join
func (tm *TopicManager) Exchange(id peer.ID, peerBuf []byte) error {
	// Initialize RPC
	exchClient := rpc.NewClient(tm.host.Host(), util.LOCAL_PROTOCOL)
	var reply LocalServiceResponse
	var args LocalServiceArgs

	// Set Args
	args.Peer = peerBuf

	// Call to Peer
	err := exchClient.Call(id, util.LOCAL_RPC_SERVICE, util.LOCAL_METHOD_EXCHANGE, args, &reply)
	if err != nil {
		return err
	}

	// Received Message
	remotePeer := &md.Peer{}
	err = proto.Unmarshal(reply.Peer, remotePeer)

	// Send Error
	if err != nil {
		return err
	}

	// Update Peer with new data
	tm.RefreshLobby(md.NewJoinLocalEvent(remotePeer))
	return nil
}

// # Calls Exchange on Local Lobby Peer
func (ts *LocalService) ExchangeWith(ctx context.Context, args LocalServiceArgs, reply *LocalServiceResponse) error {
	// Peer Data
	remotePeer := &md.Peer{}
	err := proto.Unmarshal(args.Peer, remotePeer)
	if err != nil {
		return err
	}

	// Update Peers with Lobby
	ts.call.OnEvent(md.NewJoinLocalEvent(remotePeer))

	// Set Message data and call done
	buf, err := ts.user.Peer.Buffer()
	if err != nil {
		return err
	}
	reply.Peer = buf
	return nil
}

// @ Invite: Handles User sent InviteRequest Response
func (tm *TopicManager) Invite(id peer.ID, inv *md.InviteRequest) error {
	// Initialize Data
	isFlat := inv.IsFlatInvite()
	rpcClient := rpc.NewClient(tm.host.Host(), util.LOCAL_PROTOCOL)
	var reply LocalServiceResponse
	var args LocalServiceArgs

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(inv)
	if err != nil {
		return err
	}

	// Set Args
	args.Invite = msgBytes

	// Check Invite for Flat/Default
	if isFlat {
		// Call to Peer
		err = rpcClient.Call(id, util.LOCAL_RPC_SERVICE, util.LOCAL_METHOD_INVITE, args, &reply)
		if err != nil {
			return err
		}

		tm.handler.OnReply(id, reply.InvReply)
		return nil
	} else {
		// Call to Peer
		done := make(chan *rpc.Call, 1)
		err = rpcClient.Go(id, util.LOCAL_RPC_SERVICE, util.LOCAL_METHOD_INVITE, args, &reply, done)

		// Await Response
		call := <-done
		if call.Error != nil {
			return err
		}
		tm.handler.OnReply(id, reply.InvReply)
		return nil
	}
}

// # Calls Invite on Local Lobby Peer
func (ts *LocalService) InviteWith(ctx context.Context, args LocalServiceArgs, reply *LocalServiceResponse) error {
	// Received Message
	inv := md.InviteRequest{}
	err := proto.Unmarshal(args.Invite, &inv)
	if err != nil {
		return err
	}

	// Set Current Message and send Callback
	isFlat := inv.IsFlatInvite()
	ts.invite = &inv
	ts.call.OnInvite(args.Invite)

	// Check Invite for Flat/Default
	if isFlat {
		// Sign Contact Reply
		resp := ts.user.SignFlatReply(inv.GetFrom())

		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(resp)
		if err != nil {
			return err
		}

		reply.InvReply = msgBytes
		return nil
	} else {
		// Hold Select for Invite Type
		select {
		// Received Auth Channel Message
		case m := <-ts.respCh:
			// Convert Protobuf to bytes
			msgBytes, err := proto.Marshal(m)
			if err != nil {
				return err
			}

			// Set Message data and call done
			reply.InvReply = msgBytes
			ctx.Done()
			return nil
		}
	}
}

// @ RespondToInvite to an Invitation
func (n *TopicManager) RespondToInvite(rep *md.InviteResponse) {
	// Send to Channel
	n.service.respCh <- rep

	// Prepare Transfer
	if rep.Decision {
		n.handler.OnResponded(n.service.invite)
	}
}
