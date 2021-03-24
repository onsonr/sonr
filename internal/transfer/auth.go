package transfer

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
	dt "github.com/sonr-io/core/internal/data"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ****************** //
// ** GRPC Service ** //
// ****************** //
// Argument is AuthMessage protobuf
type AuthArgs struct {
	Data []byte
}

// Reply is also AuthMessage protobuf
type AuthResponse struct {
	Data []byte
}

// Service Struct
type AuthService struct {
	// Current Data
	call   dt.TransferCallback
	respCh chan *md.AuthReply
	invite *md.AuthInvite
}

// ^ Calls Invite on Remote Peer ^ //
func (as *AuthService) Invited(ctx context.Context, args AuthArgs, reply *AuthResponse) error {
	// Received Message
	receivedMessage := md.AuthInvite{}
	err := proto.Unmarshal(args.Data, &receivedMessage)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	// Set Current Message
	as.invite = &receivedMessage

	// Send Callback
	as.call.Invited(args.Data)

	// Hold Select for Invite Type
	select {
	// Received Auth Channel Message
	case m := <-as.respCh:

		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(m)
		if err != nil {
			sentry.CaptureException(err)
			return err
		}

		// Set Message data and call done
		reply.Data = msgBytes
		ctx.Done()
		return nil
		// Context is Done
	case <-ctx.Done():
		return nil
	}
}

// ^ Send RequestInvite to a Peer ^ //
func (pc *TransferController) RequestInvite(h host.Host, id peer.ID, msgBytes []byte) {
	// Initialize Data
	rpcClient := gorpc.NewClient(h, pc.router.Auth())
	var reply AuthResponse
	var args AuthArgs
	args.Data = msgBytes

	// Call to Peer
	done := make(chan *gorpc.Call, 1)
	err := rpcClient.Go(id, "AuthService", "Invited", args, &reply, done)

	// Await Response
	call := <-done
	if call.Error != nil {
		sentry.CaptureException(err)
		pc.call.Error(err, "Request")
	}

	// Send Callback and Reset
	pc.call.Responded(reply.Data)
	transDecs, from := pc.handleReply(reply.Data)

	// Check Response for Accept
	if transDecs {
		pc.StartOutgoing(h, id, from)
	}
}

// ^ Send Authorize transfer on RPC ^ //
func (pc *TransferController) Authorize(decision bool, contact *md.Contact, peer *md.Peer) {
	// Get Offer Message
	offerMsg := pc.auth.invite

	// @ Pass Contact Back
	if offerMsg.Payload == md.Payload_CONTACT {
		// Create Accept Response
		card := dt.NewCardFromContact(peer, contact, md.TransferCard_REPLY)
		resp := &md.AuthReply{
			IsRemote: offerMsg.IsRemote,
			From:     peer,
			Type:     md.AuthReply_Contact,
			Card:     &card,
		}
		// Send to Channel
		pc.auth.respCh <- resp
	} else {
		// Prepare for Transfer
		if decision {
			pc.NewIncoming(offerMsg)
		}

		// Create Accept Response
		resp := &md.AuthReply{
			IsRemote: offerMsg.IsRemote,
			From:     peer,
			Type:     md.AuthReply_Transfer,
			Decision: decision,
		}
		// Send to Channel
		pc.auth.respCh <- resp
	}
}

// ^ Send Authorize transfer on RPC ^ //
func (pc *TransferController) Cancel(peer *md.Peer) {
	if pc.auth.invite != nil {
		// Create Cancel Reply
		reply := &md.AuthReply{
			From:     peer,
			Type:     md.AuthReply_None,
			Decision: false,
		}

		// Send to Channel
		pc.auth.respCh <- reply

		// Clear Current Invite
		pc.auth.invite = nil
	}
}

// @ Helper Method to Handle Reply
func (pc *TransferController) handleReply(data []byte) (bool, *md.Peer) {
	// Received Message
	resp := md.AuthReply{}
	err := proto.Unmarshal(data, &resp)
	if err != nil {
		pc.call.Error(err, "handleReply")
		sentry.CaptureException(err)
		return false, nil
	}
	return resp.Decision && resp.Type == md.AuthReply_Transfer, resp.From
}
