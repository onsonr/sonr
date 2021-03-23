package node

import (
	"errors"

	sentry "github.com/getsentry/sentry-go"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (n *Node) Invite(req *md.InviteRequest) {
	// @ 2. Check Transfer Type
	if req.Type == md.InviteRequest_Contact || req.Type == md.InviteRequest_URL {
		// @ 3. Send Invite to Peer
		// Set Contact
		req.Contact = n.contact
		invMsg := md.NewInviteFromRequest(req, n.peer)

		if req.IsRemote {
			// Start Remote Point
			err := n.transfer.StartRemote(&invMsg)
			if err != nil {
				n.error(err, "StartRemotePoint")
			}
		} else {
			if n.HasPeer(n.local, req.To.Id.Peer) {
				// Get PeerID and Check error
				id, _, err := n.GetPeer(n.local, req.To.Id.Peer)
				if err != nil {
					sentry.CaptureException(err)
				}

				// Run Routine
				go func(inv *md.AuthInvite) {
					// Convert Protobuf to bytes
					msgBytes, err := proto.Marshal(inv)
					if err != nil {
						sentry.CaptureException(err)
					}

					n.transfer.RequestInvite(n.host, id, msgBytes)
				}(&invMsg)
			} else {
				n.error(errors.New("Invalid Peer"), "Invite")
			}
		}

	} else {
		// File Transfer
		n.fs.AddFromRequest(req)
	}

	// Update Status
	n.status = md.Status_PENDING
}

// ^ Respond to an Invitation ^ //
func (n *Node) Respond(decision bool) {
	// Send Response on PeerConnection
	n.transfer.Authorize(decision, n.contact, n.peer)

	// Update Status
	if decision {
		n.status = md.Status_INPROGRESS
	} else {
		n.status = md.Status_AVAILABLE
	}
}
