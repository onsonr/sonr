package remote

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/libp2p/go-libp2p-core/peer"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ****************** //
// ** GRPC Service ** //
// ****************** //
// DirectArgs is Peer protobuf
type DirectArgs struct {
	Data []byte
}

// DirectResponse is also Peer protobuf
type DirectResponse struct {
	Data []byte
}

// Service Struct
type DirectService struct {
	getUser    md.ReturnPeer
	updatePeer md.UpdatePeer
}

// ^ Calls Invite on Remote Peer ^ //
func (ps *DirectService) DirectWith(ctx context.Context, args DirectArgs, reply *DirectResponse) error {
	// Peer Data
	remotePeer := &md.Peer{}
	err := proto.Unmarshal(args.Data, remotePeer)
	if err != nil {
		return err
	}

	// Update Peers
	ps.updatePeer(remotePeer)
	userPeer := ps.getUser()

	// Convert Protobuf to bytes
	replyData, err := proto.Marshal(userPeer)
	if err != nil {
		return err
	}

	// Set Message data and call done
	reply.Data = replyData
	return nil
}

// ^ Calls Invite on Remote Peer ^ //
func (rp *RemotePoint) Direct(id peer.ID) {
	// Get Peer Data
	bytes, err := proto.Marshal(rp.invite)
	if err != nil {
		rp.call.Error(err, "Direct")
	}
	var reply DirectResponse
	var args DirectArgs
	args.Data = bytes

	// Initialize RPC
	rpcClient := gorpc.NewClient(rp.host, rp.router.Direct(rp.point))

	// Call to Peer
	done := make(chan *gorpc.Call, 1)
	err = rpcClient.Go(id, "DirectService", "DirectWith", args, &reply, done)

	// Await Response
	call := <-done
	if call.Error != nil {
		sentry.CaptureException(err)
		rp.call.Error(err, "Direct")
	}

	// Send Callback and Reset
	rp.call.Responded(reply.Data)
	transDecs, from := rp.handleReply(reply.Data)

	// Check Response for Accept
	if transDecs {
		rp.transfer.StartOutgoing(rp.host, id, from)
	}
}

// @ Helper Method to Handle Reply
func (rp *RemotePoint) handleReply(data []byte) (bool, *md.Peer) {
	// Received Message
	resp := md.AuthReply{}
	err := proto.Unmarshal(data, &resp)
	if err != nil {
		rp.call.Error(err, "handleReply")
		sentry.CaptureException(err)
		return false, nil
	}
	return resp.Decision && resp.Type == md.AuthReply_Transfer, resp.From
}
