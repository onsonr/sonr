package transfer

import (
	"context"
	"log"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
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
	call      md.TransferCallback
	respCh    chan *md.AuthReply
	inviteMsg *md.AuthInvite
}

// ^ Calls Invite on Remote Peer ^ //
func (as *AuthService) Invited(ctx context.Context, args AuthArgs, reply *AuthResponse) error {
	// Set Current Message
	as.setInvite(args.Data)

	// Hold Select for Invite Type
	if !as.inviteMsg.IsDirect {
		select {
		// Received Auth Channel Message
		case m := <-as.respCh:

			// Convert Protobuf to bytes
			msgBytes, err := proto.Marshal(m)
			if err != nil {
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

	// Begin Direct Transfer
	log.Println("Direct Transfer")
	return nil
}

// ^ Send Request to a Peer ^ //
func (pc *TransferController) Request(h host.Host, id peer.ID, msgBytes []byte) {
	// Initialize Data
	rpcClient := gorpc.NewClient(h, protocol.ID("/sonr/transfer/auth"))
	var reply AuthResponse
	var args AuthArgs
	args.Data = msgBytes

	// Call to Peer
	done := make(chan *gorpc.Call, 1)
	err := rpcClient.Go(id, "AuthService", "Invited", args, &reply, done)

	// Await Response
	call := <-done
	if call.Error != nil {
		pc.call.Error(err)
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
	// Generate Reply
	reply := md.NewReplyFromDecision(md.AuthOpts{
		Decision: decision,
		Peer:     peer,
		Contact:  contact,
		Invite:   pc.auth.inviteMsg,
	})

	// Send to Channel
	pc.auth.respCh <- &reply

	// Clear Current Invite
	pc.auth.clear()
}

// ^ Send Authorize transfer on RPC ^ //
func (pc *TransferController) Cancel(peer *md.Peer) {
	// Create Cancel Reply
	reply := md.NewReplyFromDecision(md.AuthOpts{
		Peer:     peer,
		IsCancel: true,
	})

	// Send to Channel
	pc.auth.respCh <- &reply

	// Clear Current Invite
	pc.auth.clear()
}

// @ Helper Method to Handle Reply
func (pc *TransferController) handleReply(data []byte) (bool, *md.Peer) {
	// Received Message
	resp := md.AuthReply{}
	err := proto.Unmarshal(data, &resp)
	if err != nil {
		pc.call.Error(err)
		return false, nil
	}
	return resp.Decision && resp.Type == md.AuthReply_Transfer, resp.From
}

// @ Helper Method Clears Current Invite
func (as *AuthService) clear() {
	as.inviteMsg = nil
}

// @ Helper Method Sets Current Invite
func (as *AuthService) setInvite(data []byte) error {
	// Send Callback
	as.call.Invited(data)

	// Received Message
	receivedMessage := &md.AuthInvite{}
	err := proto.Unmarshal(data, receivedMessage)
	if err != nil {
		return err
	}

	// Set Current Message
	as.inviteMsg = receivedMessage
	return nil
}
