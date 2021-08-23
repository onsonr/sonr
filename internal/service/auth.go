package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/libp2p/go-libp2p-core/peer"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	"github.com/sonr-io/core/pkg/data"
	"github.com/sonr-io/core/pkg/util"
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
	handler         ServiceHandler
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
		handler: sc.handler,
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
	isDirect := inv.IsDirectInvite()
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

	// Check Invite for Direct transfer
	if isDirect {
		// Call to Peer
		err = rpcClient.Call(id, util.AUTH_RPC_SERVICE, util.AUTH_METHOD_INVITE, args, &reply)
		if err != nil {
			data.LogError(err)
			return err
		}

		tm.handler.OnReply(id, reply.InvReply)
		return nil
	} else {
		// Call to Peer
		done := make(chan *rpc.Call, 1)
		err = rpcClient.Go(id, util.AUTH_RPC_SERVICE, util.AUTH_METHOD_INVITE, args, &reply, done)

		// Await Response
		call := <-done
		if call.Error != nil {
			data.LogError(err)
			return err
		}
		tm.handler.OnReply(id, reply.InvReply)
		return nil
	}
}

// InviteWith # Calls Invite on Local Lobby Peer
func (ts *AuthService) InviteWith(ctx context.Context, args AuthServiceArgs, reply *AuthServiceResponse) error {
	// Received Message
	inv := data.InviteRequest{}
	err := proto.Unmarshal(args.Invite, &inv)
	if err != nil {
		data.LogError(err)
		return err
	}

	// Set Current Message and send Callback
	isFlat := inv.IsDirectInvite()
	ts.invite = &inv
	ts.handler.OnInvite(args.Invite)

	// Check Invite for Flat/Default
	if isFlat {
		// Sign Contact Reply
		resp := ts.device.ReplyToFlat(inv.GetFrom())

		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(resp)
		if err != nil {
			data.LogError(err)
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
				data.LogError(err)
				return err
			}

			// Set Message data and call done
			reply.InvReply = msgBytes
			ctx.Done()
			return nil
		}
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
			tm.handler.OnLink(false, false, id, nil)
			return err
		}
		tm.handler.OnLink(reply.LinkResult, false, id, reply.LinkResponse)
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
		res, err := proto.Marshal(result)
		if err != nil {
			return err
		}

		// Update Properties
		reply.LinkResponse = buf
		reply.LinkResult = ok
		ts.isLinkingActive = !reply.LinkResult

		// Return Result
		data.LogInfo(fmt.Sprintf("Link Result: %v", result))
		ts.handler.OnLink(ok, true, peer.ID(inv.GetFrom().PeerID()), reply.LinkResponse)
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
		tm.handler.OnConfirmed(tm.Auth.invite)
	}
}
