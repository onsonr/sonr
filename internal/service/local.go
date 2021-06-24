package service

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	rpc "github.com/libp2p/go-libp2p-gorpc"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

// LocalServiceArgs ExchangeArgs is Peer protobuf
type LocalServiceArgs struct {
	Peer   []byte
	Invite []byte
}

// LocalServiceResponse ExchangeResponse is also Peer protobuf
type LocalServiceResponse struct {
	InvReply []byte
	Peer     []byte
}

type LocalService struct {
	ServiceClient
	handler ServiceHandler
	user    *md.User
	respCh  chan *md.InviteResponse
	invite  *md.InviteRequest
}

func (sc *serviceClient) StartLocal() *md.SonrError {
	// Start Exchange Server
	localServer := rpc.NewServer(sc.host.Host(), util.LOCAL_PROTOCOL)
	psv := LocalService{
		user:    sc.user,
		handler: sc.handler,
		respCh:  make(chan *md.InviteResponse, util.TOPIC_MAX_MESSAGES),
	}

	// Register Service
	err := localServer.RegisterName(util.LOCAL_RPC_SERVICE, &psv)
	if err != nil {
		return md.NewError(err, md.ErrorMessage_TOPIC_RPC)
	}
	sc.Local = &psv
	return nil
}

// Invite @ Invite: Handles User sent InviteRequest Response
func (tm *serviceClient) Invite(id peer.ID, inv *md.InviteRequest) error {
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

// InviteWith # Calls Invite on Local Lobby Peer
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
	ts.handler.OnInvite(args.Invite)

	// Check Invite for Flat/Default
	if isFlat {
		// Sign Contact Reply
		resp := ts.user.ReplyToFlat(inv.GetFrom())

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

// RespondToInvite @ RespondToInvite to an Invitation
func (tm *serviceClient) Respond(rep *md.InviteResponse) {
	// Send to Channel
	tm.Local.respCh <- rep

	// Prepare Transfer
	if rep.Decision {
		tm.handler.OnConfirmed(tm.Local.invite)
	}
}
