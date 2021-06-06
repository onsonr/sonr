package topic

import (
	"context"
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	net "github.com/sonr-io/core/internal/host"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

const LOCAL_RPC_SERVICE = "LocalService"
const LOCAL_METHOD_EXCHANGE = "ExchangeWith"
const LOCAL_METHOD_FLAT = "FlatWith"
const LOCAL_METHOD_INVITE = "InviteWith"

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
	call  ClientHandler
	lobby *md.Lobby
	user  *md.User

	respCh chan *md.AuthReply
	invite *md.AuthInvite
}

// ^ Create New Contained Topic Manager ^ //
func NewLocal(ctx context.Context, h *net.HostNode, u *md.User, name string, th ClientHandler) (*TopicManager, *md.SonrError) {
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
		lobby:        md.NewLocalLobby(u),
		lobbyType:    md.Lobby_LOCAL,
		localEvents:  make(chan *md.LocalEvent, K_MAX_MESSAGES),
		subscription: sub,
		topic:        topic,
	}

	// Start Exchange Server
	localServer := rpc.NewServer(h.Host, LOCAL_SERVICE_PID)
	psv := LocalService{
		lobby:  mgr.lobby,
		user:   u,
		call:   th,
		respCh: make(chan *md.AuthReply, K_MAX_MESSAGES),
	}

	// Register Service
	err := localServer.RegisterName(LOCAL_RPC_SERVICE, &psv)
	if err != nil {
		return nil, md.NewError(err, md.ErrorMessage_TOPIC_RPC)
	}

	// Set Service
	mgr.service = &psv
	go mgr.handleTopicEvents()
	go mgr.handleTopicMessages()
	go mgr.processTopicMessages()
	return mgr, nil
}

// ^ Send Updated Lobby ^
func (tm *TopicManager) RefreshLobby() {
	tm.handler.OnRefresh(tm.lobby)
}

// ^ SendLocal message to specific peer in topic ^
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

// ^ Flat: Handles User sent AuthInvite Response on FlatMode ^
func (tm *TopicManager) Flat(id peer.ID, inv *md.AuthInvite) error {
	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(inv)
	if err != nil {
		return err
	}

	// Initialize Data
	rpcClient := rpc.NewClient(tm.host.Host, LOCAL_SERVICE_PID)
	var reply LocalServiceResponse
	var args LocalServiceArgs
	args.Invite = msgBytes

	// Call to Peer
	err = rpcClient.Call(id, LOCAL_RPC_SERVICE, LOCAL_METHOD_FLAT, args, &reply)
	if err != nil {
		return err
	}

	tm.handler.OnReply(id, reply.InvReply)
	return nil
}

// ^ Calls Invite on Remote Peer for Flat Mode and makes Direct Response ^ //
func (ts *LocalService) FlatWith(ctx context.Context, args LocalServiceArgs, reply *LocalServiceResponse) error {
	// Received Message
	receivedMessage := md.AuthInvite{}
	err := proto.Unmarshal(args.Invite, &receivedMessage)
	if err != nil {
		return err
	}

	// Set Current Message and send Callback
	ts.invite = &receivedMessage
	ts.call.OnInvite(args.Invite)

	// Sign Contact Reply
	resp := ts.user.SignFlatReply(receivedMessage.GetFrom())

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(resp)
	if err != nil {
		return err
	}

	reply.InvReply = msgBytes
	return nil
}

// ^ Starts Exchange on Local Peer Join ^ //
func (tm *TopicManager) Exchange(id peer.ID, peerBuf []byte) error {
	// Initialize RPC
	exchClient := rpc.NewClient(tm.host.Host, LOCAL_SERVICE_PID)
	var reply LocalServiceResponse
	var args LocalServiceArgs

	// Set Args
	args.Peer = peerBuf

	// Call to Peer
	err := exchClient.Call(id, LOCAL_RPC_SERVICE, LOCAL_METHOD_EXCHANGE, args, &reply)
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
	tm.lobby.Add(remotePeer)
	tm.RefreshLobby()
	return nil
}

// ^ Calls Exchange on Local Lobby Peer ^ //
func (ts *LocalService) ExchangeWith(ctx context.Context, args LocalServiceArgs, reply *LocalServiceResponse) error {
	// Peer Data
	remotePeer := &md.Peer{}
	err := proto.Unmarshal(args.Peer, remotePeer)
	if err != nil {
		return err
	}

	// Update Peers with Lobby
	ts.lobby.Add(remotePeer)
	ts.call.OnRefresh(ts.lobby)

	// Set Message data and call done
	buf, err := ts.user.Peer.Buffer()
	if err != nil {
		return err
	}
	reply.Peer = buf
	return nil
}

// ^ Invite: Handles User sent AuthInvite Response ^
func (tm *TopicManager) Invite(id peer.ID, inv *md.AuthInvite) error {
	// Initialize Data
	rpcClient := rpc.NewClient(tm.host.Host, LOCAL_SERVICE_PID)
	var reply LocalServiceResponse
	var args LocalServiceArgs

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(inv)
	if err != nil {
		return err
	}

	args.Invite = msgBytes

	// Call to Peer
	done := make(chan *rpc.Call, 1)
	err = rpcClient.Go(id, LOCAL_RPC_SERVICE, LOCAL_METHOD_INVITE, args, &reply, done)

	// Await Response
	call := <-done
	if call.Error != nil {
		log.Println("Error Occurred: ", err)
		return err
	}
	log.Println("--- Received Reply ---")
	tm.handler.OnReply(id, reply.InvReply)
	return nil
}

// ^ Calls Invite on Local Lobby Peer ^ //
func (ts *LocalService) InviteWith(ctx context.Context, args LocalServiceArgs, reply *LocalServiceResponse) error {
	// Received Message
	receivedMessage := md.AuthInvite{}
	err := proto.Unmarshal(args.Invite, &receivedMessage)
	if err != nil {
		return err
	}

	// Set Current Message and send Callback
	ts.invite = &receivedMessage
	ts.call.OnInvite(args.Invite)
	log.Println("--- Received Invite ---")

	// Hold Select for Invite Type
	select {
	// Received Auth Channel Message
	case m := <-ts.respCh:
		log.Println("--- Sending Reply ---")
		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(m)
		if err != nil {
			log.Println("Error Occurred: ", err)
			return err
		}

		// Set Message data and call done
		reply.InvReply = msgBytes
		ctx.Done()
		return nil
	}
}

// ^ RespondToInvite to an Invitation ^ //
func (n *TopicManager) RespondToInvite(rep *md.AuthReply) {
	log.Println("--- Received Response for Invite ---")
	// Send to Channel
	n.service.respCh <- rep

	// Prepare Transfer
	if rep.Decision {
		log.Println("--- Responding for Accept ---")
		n.handler.OnResponded(n.service.invite)
	} else {
		log.Println("--- Responding for Decline ---")
	}

}
