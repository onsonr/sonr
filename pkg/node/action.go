package node

import (
	"errors"
	"math"
	"time"

	sentry "github.com/getsentry/sentry-go"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multihash"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ^ User Node Info ^ //
// @ ID Returns Peer ID
func (n *Node) ID() peer.ID {
	return n.host.ID()
}

// @ Peer returns Current Peer Info
func (n *Node) Peer() *md.Peer {
	return n.peer
}

// @ Peer returns Current Peer Info as Buffer
func (n *Node) PeerBuf() []byte {
	// Convert to bytes
	buf, err := proto.Marshal(n.peer)
	if err != nil {
		sentry.CaptureException(err)
		return nil
	}
	return buf
}

// @ Peer returns Current Peer Info as Content ID
func (n *Node) PeerCID() (cid.Cid, error) {
	// Convert to bytes
	buf, err := proto.Marshal(n.peer)
	if err != nil {
		sentry.CaptureException(err)
		return cid.Undef, err
	}

	// Encode Multihash
	mhash, err := multihash.EncodeName(buf, n.peer.Id.Peer)
	if err != nil {
		sentry.CaptureException(err)
		return cid.Undef, err
	}

	// Return Key
	key := cid.NewCidV1(cid.DagProtobuf, mhash)
	return key, nil
}

// ^ Update proximity/direction and Notify Lobby ^ //
func (n *Node) Update(facing float64, heading float64) {
	// Update User Values
	var faceDir float64
	var faceAnpd float64
	var headDir float64
	var headAnpd float64
	faceDir = math.Round(facing*100) / 100
	headDir = math.Round(heading*100) / 100
	desg := int((facing / 11.25) + 0.25)

	// Find Antipodal
	if facing > 180 {
		faceAnpd = math.Round((facing-180)*100) / 100
	} else {
		faceAnpd = math.Round((facing+180)*100) / 100
	}

	// Find Antipodal
	if heading > 180 {
		headAnpd = math.Round((heading-180)*100) / 100
	} else {
		headAnpd = math.Round((heading+180)*100) / 100
	}

	// Set Position
	n.peer.Position = &md.Position{
		Facing:           faceDir,
		FacingAntipodal:  faceAnpd,
		Heading:          headDir,
		HeadingAntipodal: headAnpd,
		Designation:      md.Position_Designation(desg % 32),
	}

	// Inform Lobby
	err := n.local.Update(n.peer)
	if err != nil {
		sentry.CaptureException(err)
	}
}

// ^ Send Direct Message to Peer in Lobby ^ //
func (n *Node) Message(content string, to string) {
	if n.HasPeer(n.local, to) {
		// Inform Lobby
		err := n.local.Message(content, to, n.peer)
		if err != nil {
			sentry.CaptureException(err)
		}
	}
}

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

// ^ Join Remote File with Words ^ //
func (n *Node) JoinRemote(data string) {
	// Validate
	_, err := n.transfer.JoinRemote(data)
	if err != nil {
		// Lobby non-existent
		sentry.CaptureException(err)
		n.error(err, "Join Remote")
	}
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

// ^ Link with a QR Code ^ //
func (n *Node) LinkDevice(json string) {
	// Convert String to Bytes
	request := md.LinkRequest{}

	// Convert to Peer Protobuf
	err := protojson.Unmarshal([]byte(json), &request)
	if err != nil {
		sentry.CaptureException(err)
	}

	// Link Device
	err = n.fs.SaveDevice(request.Device)
	if err != nil {
		sentry.CaptureException(err)
	}
}

// ^ Link with a QR Code ^ //
func (n *Node) LinkRequest(name string) *md.LinkRequest {
	// Set Device
	device := n.device
	device.Name = name

	// Create Expiry - 1min 30s
	timein := time.Now().Local().Add(
		time.Minute*time.Duration(1) +
			time.Second*time.Duration(30))

	// Return Request
	return &md.LinkRequest{
		Device: device,
		Peer:   n.Peer(),
		Expiry: int32(timein.Unix()),
	}
}

// ^ Updates Current Contact Card ^
func (n *Node) SetContact(newContact *md.Contact) {

	// Set Node Contact
	n.contact = newContact

	// Update Peer Profile
	n.peer.Profile = &md.Profile{
		FirstName: newContact.GetFirstName(),
		LastName:  newContact.GetLastName(),
		Picture:   newContact.GetPicture(),
	}

	// Set User Contact
	err := n.fs.SaveContact(newContact)
	if err != nil {
		sentry.CaptureException(err)
	}
}

// ^ Close Ends All Network Communication ^
func (n *Node) Pause() {
	// Check if Response Is Invited
	if n.status == md.Status_INVITED {
		n.transfer.Cancel(n.peer)
	}
	md.GetState().Pause()
}

// ^ Close Ends All Network Communication ^
func (n *Node) Resume() {
	md.GetState().Resume()
}

// ^ Close Ends All Network Communication ^
func (n *Node) Stop() {
	// Check if Response Is Invited
	if n.status == md.Status_INVITED {
		n.transfer.Cancel(n.peer)
	}
	n.host.Close()
}
