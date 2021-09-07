package service

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	"github.com/sonr-io/core/pkg/data"
	"github.com/sonr-io/core/pkg/util"
	"github.com/sonr-io/core/tools/emitter"
	"github.com/sonr-io/core/tools/logger"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

// AuthServiceArgs ExchangeArgs is Peer protobuf
type AuthServiceArgs struct {
	Peer   []byte
	Invite []byte
	Link   []byte
}

// AuthServiceResponse ExchangeResponse is also Peer protobuf
type AuthServiceResponse struct {
	InvReply     []byte
	Peer         []byte
	LinkResponse []byte
	LinkResult   bool
}

type AuthService struct {
	emitter         *emitter.Emitter
	device          *data.Device
	respCh          chan *data.InviteResponse
	linkCh          chan *data.LinkRequest
	invite          *data.InviteRequest
	isLinkingActive bool
}

// Starts New Auth Instance
func (sc *serviceClient) StartAuth() *data.SonrError {
	// Start Exchange Server
	localServer := rpc.NewServer(sc.host.Host(), util.AUTH_PROTOCOL)
	psv := AuthService{
		device:  sc.device,
		emitter: sc.emitter,
		respCh:  make(chan *data.InviteResponse, util.MAX_CHAN_DATA),
		linkCh:  make(chan *data.LinkRequest, util.MAX_CHAN_DATA),
	}

	// Register Service
	err := localServer.RegisterName(util.AUTH_RPC_SERVICE, &psv)
	if err != nil {
		return data.NewError(err, data.ErrorEvent_ROOM_RPC)
	}
	sc.Auth = &psv
	return nil
}

// Enable/Disable Linking from LinkRequest
func (sc *serviceClient) HandleLinking(req *data.LinkRequest) {
	// Check Link Request Type
	if req.Type == data.LinkRequest_RECEIVE {
		// Set Active
		sc.Auth.isLinkingActive = true
	} else if req.Type == data.LinkRequest_CANCEL {
		// Set Inactive
		sc.Auth.isLinkingActive = false
	}
}

// Invite @ Invite: Handles User sent InviteRequest Response
func (tm *serviceClient) Invite(id peer.ID, inv *data.InviteRequest) error {
	// Initialize Data
	rpcClient := rpc.NewClient(tm.host.Host(), util.AUTH_PROTOCOL)
	var reply AuthServiceResponse
	var args AuthServiceArgs

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(inv)
	if err != nil {
		return err
	}

	// Set Args
	args.Invite = msgBytes

	// Call to Peer
	done := make(chan *rpc.Call, 1)
	err = rpcClient.Go(id, util.AUTH_RPC_SERVICE, util.AUTH_METHOD_INVITE, args, &reply, done)

	// Await Response
	call := <-done
	if call.Error != nil {
		logger.Error("Failed to Invite Peer", zap.Error(err))
		return err
	}
	tm.emitter.Emit(emitter.EMIT_REPLY, id, reply.InvReply)
	return nil
}

// InviteWith # Calls Invite on Local Lobby Peer
func (ts *AuthService) InviteWith(ctx context.Context, args AuthServiceArgs, reply *AuthServiceResponse) error {
	// Received Message
	inv := data.InviteRequest{}
	err := proto.Unmarshal(args.Invite, &inv)
	if err != nil {
		logger.Error("Failed to Unmarshal Invite Request", zap.Error(err))
		return err
	}

	// Set Current Message and send Callback
	ts.invite = &inv
	ts.emitter.Emit(emitter.EMIT_INVITE, args.Invite)

	// Hold Select for Invite Type
	select {
	// Received Auth Channel Message
	case m := <-ts.respCh:
		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(m)
		if err != nil {
			logger.Error("Failed to Marshal InviteResponse", zap.Error(err))
			return err
		}

		// Set Message data and call done
		reply.InvReply = msgBytes
		return nil
	}
}

// Sends Link Request to Linker type Peer
func (tm *serviceClient) Link(id peer.ID, inv *data.LinkRequest) error {
	// Check Invite
	if inv.Type == data.LinkRequest_SEND {
		// Initialize Data
		rpcClient := rpc.NewClient(tm.host.Host(), util.AUTH_PROTOCOL)
		var reply AuthServiceResponse
		var args AuthServiceArgs

		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(inv)
		if err != nil {
			return err
		}

		// Set Args
		args.Link = msgBytes

		// Call to Peer
		err = rpcClient.Call(id, util.AUTH_RPC_SERVICE, util.AUTH_METHOD_LINK, args, &reply)
		if err != nil {
			tm.emitter.Emit(emitter.EMIT_LINK, false, false, id, nil)
			return err
		}
		tm.emitter.Emit(emitter.EMIT_LINK, reply.LinkResult, false, id, reply.LinkResponse)
		return nil
	}
	return errors.New("Invite is not a Link Invite")
}

// InviteWith # Calls Invite on Local Lobby Peer
func (ts *AuthService) LinkWith(ctx context.Context, args AuthServiceArgs, reply *AuthServiceResponse) error {
	if ts.isLinkingActive {
		// Received Message
		inv := data.LinkRequest{}
		err := proto.Unmarshal(args.Link, &inv)
		if err != nil {
			return err
		}

		// Handle Status
		ok, result := ts.device.VerifyLink(&inv)
		buf, err := proto.Marshal(result)
		if err != nil {
			return err
		}

		// Update Properties
		reply.LinkResponse = buf
		reply.LinkResult = ok
		ts.isLinkingActive = !reply.LinkResult

		// Return Result
		ts.emitter.Emit(emitter.EMIT_LINK, ok, true, peer.ID(inv.GetFrom().PeerID()), reply.LinkResponse)
		return nil
	} else {
		return errors.New("Linking is not Active")
	}
}

// RespondToInvite @ RespondToInvite to an Invitation
func (tm *serviceClient) Respond(rep *data.InviteResponse) {
	// Send to Channel
	tm.Auth.respCh <- rep

	// Prepare Transfer
	if rep.Decision.Accepted() {
		tm.emitter.Emit(emitter.EMIT_CONFIRMED, tm.Auth.invite)
	}
}
