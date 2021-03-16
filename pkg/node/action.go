package node

import (
	"math"
	"time"

	sentry "github.com/getsentry/sentry-go"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

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
	err := n.lobby.Update()
	if err != nil {
		sentry.CaptureException(err)
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

		// Get PeerID and Check error
		id, _, err := n.lobby.Find(req.To.Id.Peer)
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

			n.peerConn.Request(n.host, id, msgBytes)
		}(&invMsg)
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
	n.peerConn.Authorize(decision, n.contact, n.peer)

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
