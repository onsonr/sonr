package node

import (
	"log"
	"math"
	"time"

	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ^ Update proximity/direction and Notify Lobby ^ //
func (sn *Node) Update(facing float64, heading float64) {
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
	sn.peer.Position = &md.Position{
		Facing:           faceDir,
		FacingAntipodal:  faceAnpd,
		Heading:          headDir,
		HeadingAntipodal: headAnpd,
		Designation:      md.Position_Designation(desg % 32),
	}

	// Inform Lobby
	err := sn.lobby.Update()
	if err != nil {
		sn.error(err, "Update")
	}
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (sn *Node) Invite(req *md.InviteRequest) {
	// @ 2. Check Transfer Type
	if req.Type == md.InviteRequest_Contact || req.Type == md.InviteRequest_URL {
		// @ 3. Send Invite to Peer
		// Set Contact
		req.Contact = sn.contact
		invMsg := md.NewInviteFromRequest(req, sn.peer)

		// Get PeerID and Check error
		id, _, err := sn.lobby.Find(req.To.Id.Peer)
		if err != nil {
			sn.error(err, "InviteWithContact")
		}

		// Run Routine
		go func(inv *md.AuthInvite) {
			// Convert Protobuf to bytes
			msgBytes, err := proto.Marshal(inv)
			if err != nil {
				sn.error(err, "Marshal")
			}

			sn.peerConn.Request(sn.host, id, msgBytes)
		}(&invMsg)
	} else {
		// File Transfer
		sn.queue.AddFromRequest(req)
	}

	// Update Status
	sn.Status = md.Status_PENDING
}

// ^ Respond to an Invitation ^ //
func (sn *Node) Respond(decision bool) {
	// Send Response on PeerConnection
	sn.peerConn.Authorize(decision, sn.contact, sn.peer)

	// Update Status
	if decision {
		sn.Status = md.Status_INPROGRESS
	} else {
		sn.Status = md.Status_AVAILABLE
	}
}

// ^ Info returns ALL Peer Data as Bytes^
func (sn *Node) Info() []byte {
	// Convert to bytes to view in plugin
	data, err := proto.Marshal(sn.peer)
	if err != nil {
		log.Println("Error Marshaling Lobby Data ", err)
		return nil
	}
	return data
}

// ^ Link with a QR Code ^ //
func (sn *Node) LinkDevice(json string) {
	// Convert String to Bytes
	request := md.LinkRequest{}

	// Convert to Peer Protobuf
	err := protojson.Unmarshal([]byte(json), &request)
	if err != nil {
		sn.error(err, "LinkDevice")
	}

	// Link Device
	err = sn.fs.AddDevice(request.Device)
	if err != nil {
		sn.error(err, "LinkDevice")
	}
}

// ^ Link with a QR Code ^ //
func (sn *Node) LinkRequest(name string) *md.LinkRequest {
	// Set Device
	device := sn.device
	device.Directories = sn.directories
	device.Name = name

	// Create Expiry - 1min 30s
	timein := time.Now().Local().Add(
		time.Minute*time.Duration(1) +
			time.Second*time.Duration(30))

	// Return Request
	return &md.LinkRequest{
		Device: device,
		Peer:   sn.Peer(),
		Expiry: int32(timein.Unix()),
	}
}

// ^ Peer returns Current Peer Info ^
func (sn *Node) Peer() *md.Peer {
	return sn.peer
}

// ^ Updates Current Contact Card ^
func (sn *Node) SetContact(newContact *md.Contact) {

	// Set Node Contact
	sn.contact = newContact

	// Update Peer Profile
	sn.peer.Profile = &md.Profile{
		FirstName: newContact.GetFirstName(),
		LastName:  newContact.GetLastName(),
		Picture:   newContact.GetPicture(),
	}

	// Set User Contact
	err := sn.fs.UpdateContact(newContact)
	if err != nil {
		sn.error(err, "SetContact")
	}
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Pause() {
	// Check if Response Is Invited
	if sn.Status == md.Status_INVITED {
		sn.peerConn.Cancel(sn.peer)
	}
	err := sn.lobby.Standby()
	if err != nil {
		sn.error(err, "Pause")
	}
	md.GetState().Pause()
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Resume() {
	err := sn.lobby.Resume()
	if err != nil {
		sn.error(err, "Resume")
	}
	md.GetState().Resume()
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Stop() {
	// Check if Response Is Invited
	if sn.Status == md.Status_INVITED {
		sn.peerConn.Cancel(sn.peer)
	}
	sn.host.Close()
}
