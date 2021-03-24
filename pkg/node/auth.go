package node

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	msgio "github.com/libp2p/go-msgio"
	sf "github.com/sonr-io/core/internal/file"
	dt "github.com/sonr-io/core/pkg/data"

	md "github.com/sonr-io/core/internal/models"
	tr "github.com/sonr-io/core/pkg/transfer"
	"google.golang.org/protobuf/proto"
)

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (n *Node) InviteLink(req *md.InviteRequest, p *md.Peer) error {
	// @ 3. Send Invite to Peer
	if n.HasPeer(n.local, req.To.Id.Peer) {
		// Get PeerID and Check error
		id, _, err := n.FindPeerInTopic(n.local, req.To.Id.Peer)
		if err != nil {
			return err
		}

		// Get URL Data
		urlInfo, err := dt.GetPageInfoFromUrl(req.Url)
		if err != nil {
			log.Println(err)
			urlInfo = &md.URLLink{
				Link: req.Url,
			}
		}

		// Build Invite Message
		invMsg := md.AuthInvite{
			IsRemote: req.IsRemote,
			From:     p,
			Payload:  md.Payload_URL,
			Card: &md.TransferCard{
				// SQL Properties
				Payload:  md.Payload_URL,
				Received: int32(time.Now().Unix()),
				Platform: p.Platform,

				// Transfer Properties
				Status: md.TransferCard_DIRECT,

				// Owner Properties
				Username:  p.Profile.Username,
				FirstName: p.Profile.FirstName,
				LastName:  p.Profile.LastName,

				// Data Properties
				Url: urlInfo,
			},
		}
		// Run Routine
		go n.handleAuthInviteResponse(id, &invMsg, p, nil)
	} else {
		return errors.New("Invalid Peer")
	}
	return nil
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (n *Node) InviteContact(req *md.InviteRequest, p *md.Peer, c *md.Contact) error {
	// @ 3. Send Invite to Peer
	if n.HasPeer(n.local, req.To.Id.Peer) {
		// Get PeerID and Check error
		id, _, err := n.FindPeerInTopic(n.local, req.To.Id.Peer)
		if err != nil {
			return err
		}

		// Build Invite Message
		invMsg := md.AuthInvite{
			IsRemote: req.IsRemote,
			From:     p,
			Payload:  md.Payload_CONTACT,
			Card: &md.TransferCard{
				// SQL Properties
				Payload:  md.Payload_CONTACT,
				Received: int32(time.Now().Unix()),
				Preview:  p.Profile.Picture,
				Platform: p.Platform,

				// Transfer Properties
				Status: md.TransferCard_DIRECT,

				// Owner Properties
				Username:  p.Profile.Username,
				FirstName: p.Profile.FirstName,
				LastName:  p.Profile.LastName,

				// Data Properties
				Contact: c,
			},
		}

		// Run Routine
		go n.handleAuthInviteResponse(id, &invMsg, p, nil)
	} else {
		return errors.New("Invalid Peer")
	}
	return nil
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (n *Node) InviteFile(card *md.TransferCard, req *md.InviteRequest, p *md.Peer, cf *sf.ProcessedFile) error {
	card.Status = md.TransferCard_INVITE

	// Create Invite Message
	invMsg := md.AuthInvite{
		From:    p,
		Payload: card.Payload,
		Card:    card,
	}

	// @ Check for Remote

	// Get PeerID
	id, _, err := n.FindPeerInTopic(n.local, req.To.Id.Peer)
	if err != nil {
		return err
	}

	// Run Routine
	go n.handleAuthInviteResponse(id, &invMsg, p, cf)
	return nil
}

// ^ Respond to an Invitation ^ //
func (n *Node) Respond(decision bool, fs *sf.FileSystem, p *md.Peer, c *md.Contact) {
	// Check Decision
	if decision {
		n.host.SetStreamHandler(n.router.Transfer(), n.handleTransferIncoming)
		n.incoming = tr.CreateIncomingFile(n.auth.invite, fs, n.call)
	}

	// @ Pass Contact Back
	if n.auth.invite.Payload == md.Payload_CONTACT {
		// Create Accept Response
		resp := &md.AuthReply{
			IsRemote: n.auth.invite.IsRemote,
			From:     p,
			Type:     md.AuthReply_Contact,
			Card: &md.TransferCard{
				// SQL Properties
				Payload:  md.Payload_CONTACT,
				Received: int32(time.Now().Unix()),
				Preview:  p.Profile.Picture,
				Platform: p.Platform,

				// Transfer Properties
				Status: md.TransferCard_REPLY,

				// Owner Properties
				Username:  p.Profile.Username,
				FirstName: p.Profile.FirstName,
				LastName:  p.Profile.LastName,

				// Data Properties
				Contact: c,
			},
		}
		// Send to Channel
		n.auth.respCh <- resp
	} else {
		// Create Accept Response
		resp := &md.AuthReply{
			IsRemote: n.auth.invite.IsRemote,
			From:     p,
			Type:     md.AuthReply_Transfer,
			Decision: decision,
		}
		// Send to Channel
		n.auth.respCh <- resp
	}
}

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
	call   dt.NodeCallback
	respCh chan *md.AuthReply
	invite *md.AuthInvite
}

// ^ Calls Invite on Remote Peer ^ //
func (as *AuthService) Invited(ctx context.Context, args AuthArgs, reply *AuthResponse) error {
	// Received Message
	receivedMessage := md.AuthInvite{}
	err := proto.Unmarshal(args.Data, &receivedMessage)
	if err != nil {
		log.Println(err)
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
			log.Println(err)
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

// ^ User has accepted, Begin Sending Transfer ^ //
func (n *Node) NewOutgoingTransfer(id peer.ID, peer *md.Peer, pf *sf.ProcessedFile) {
	// Create New Auth Stream
	stream, err := n.host.NewStream(n.ctx, id, n.router.Transfer())
	if err != nil {
		n.call.Error(err, "StartOutgoing")
	}

	outFile := tr.CreateOutgoingFile(pf, n.call)

	// Initialize Writer
	writer := msgio.NewWriter(stream)

	// Start Routine
	go outFile.WriteBase64(writer, peer)
}
