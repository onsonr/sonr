package sonr

import (
	"log"
	"math"
	"time"

	sf "github.com/sonr-io/core/internal/file"
	"github.com/sonr-io/core/internal/lifecycle"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

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
	err = addDevice(request.Device, sn.directories.Documents)
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
func (sn *Node) SetContact(conBytes []byte) {
	// Unmarshal Data
	newContact := &md.Contact{}
	err := proto.Unmarshal(conBytes, newContact)
	if err != nil {
		log.Println(err)
	}

	// Set Node Contact
	sn.contact = newContact

	// Set User Contact
	err = updateContact(newContact, sn.directories.Documents)
	if err != nil {
		sn.error(err, "SetContact")
	}
}

// ^ Update proximity/direction and Notify Lobby ^ //
func (sn *Node) Update(direction float64) {
	// ** Initialize ** //
	// Update User Values
	var dir float64
	var anpd float64
	dir = math.Round(direction*100) / 100

	// Find Antipodal
	if direction > 180 {
		anpd = math.Round((direction-180)*100) / 100
	} else {
		anpd = math.Round((direction+180)*100) / 100
	}

	// Set Position
	sn.peer.Position = &md.Position{
		Direction: dir,
		Antipodal: anpd,
	}

	// Inform Lobby
	err := sn.lobby.Update()
	if err != nil {
		sn.error(err, "Update")
	}
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (sn *Node) Invite(reqBytes []byte) {
	// @ 1. Initialize from Request
	req := &md.InviteRequest{}
	err := proto.Unmarshal(reqBytes, req)
	if err != nil {
		log.Println(err)
	}

	// @ 2. Check Transfer Type
	if req.Type == md.InviteRequest_File {
		// Single File Transfer
		safeFile := sf.NewProcessedFile(req, sn.peer.Profile, sn.queued, sn.error)
		sn.files = append(sn.files, safeFile)
	} else if req.Type == md.InviteRequest_MultiFiles {
		// Batch File Transfer
		safeFiles := sf.NewBatchProcessFiles(req, sn.peer.Profile, sn.queued, sn.error)
		sn.files = safeFiles
	} else if req.Type == md.InviteRequest_Contact || req.Type == md.InviteRequest_URL {
		// @ 3. Send Invite to Peer
		// Set Contact
		req.Contact = sn.contact
		invMsg := sf.NewInviteFromRequest(req, sn.peer)

		// Get PeerID and Check error
		id, _, err := sn.lobby.Find(req.To.Id)
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
	}

	// Update Status
	sn.status = md.Status_PENDING
}

// ^ Respond to an Invitation ^ //
func (sn *Node) Respond(decision bool) {
	// Send Response on PeerConnection
	sn.peerConn.Authorize(decision, sn.contact, sn.peer)

	// Update Status
	if decision {
		sn.status = md.Status_INPROGRESS
	} else {
		sn.status = md.Status_AVAILABLE
	}
}

// ^ Reset Current Queued File Metadata ^ //
func (sn *Node) ResetFile() {
	// Reset Files Slice
	sn.files = nil
	sn.files = make([]*sf.ProcessedFile, maxFileBufferSize)
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Pause() {
	// Check if Response Is Invited
	if sn.status == md.Status_INVITED {
		sn.Respond(false)
	}
	err := sn.lobby.Standby()
	if err != nil {
		sn.error(err, "Pause")
	}
	lifecycle.GetState().Pause()
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Resume() {
	err := sn.lobby.Resume()
	if err != nil {
		sn.error(err, "Resume")
	}
	lifecycle.GetState().Resume()
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Stop() {
	// Check if Response Is Invited
	if sn.status == md.Status_INVITED {
		sn.Respond(false)
	}
	sn.host.Close()
}
