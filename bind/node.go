package sonr

import (
	"log"
	"math"

	sf "github.com/sonr-io/core/internal/file"
	lf "github.com/sonr-io/core/internal/lifecycle"
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
func (sn *Node) LinkDevice(peerString string) {
	// Convert String to Bytes
	peer := md.Peer{}

	// Convert to Peer Protobuf
	err := protojson.Unmarshal([]byte(peerString), &peer)
	if err != nil {
		sn.Error(err, "LinkDevice")
	}
}

// ^ Peer returns Current Peer Info ^
func (sn *Node) Peer() *md.Peer {
	return sn.peer
}

// ^ Updates Current Contact Card ^
func (sn *Node) SetContact(conBytes []byte) {
	newContact := &md.Contact{}
	err := proto.Unmarshal(conBytes, newContact)
	if err != nil {
		log.Println(err)
	}
	sn.contact = newContact
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
		sn.Error(err, "Update")
	}
}

// ^ Invite Processes Data and Sens Invite to Peer ^ //
func (sn *Node) Invite(reqBytes []byte) {
	// @ 1. Initialize from Request
	req := &md.InviteRequest{}
	err := proto.Unmarshal(reqBytes, req)
	if err != nil {
		log.Println(err)
	}

	// Get PeerID and Check error
	id, _, err := sn.lobby.Find(req.To.Id)
	if err != nil {
		sn.Error(err, "InviteWithContact")
	}

	// @ 2. Check Transfer Type
	// Process the File
	if req.Type == md.InviteRequest_File {
		safeFile := sf.NewProcessedFile(req, sn.peer.Profile, lf.ProcessCallbacks{CallQueued: sn.Queued, CallError: sn.Error})
		sn.files = append(sn.files, safeFile)
	}

	// Contact Type Attach User Contact
	if req.Type == md.InviteRequest_MultiFiles {
		safeFiles := sf.NewBatchProcessFiles(req, sn.peer.Profile, lf.ProcessCallbacks{CallQueued: sn.Queued, CallError: sn.Error})
		sn.files = safeFiles
	}

	// Contact Type Attach User Contact
	if req.Type == md.InviteRequest_Contact {
		// Set Contact
		req.Contact = sn.contact
	}

	// @ 3. Send Invite to Peer
	invMsg := sf.NewInviteFromRequest(req, sn.peer)

	// Check if ID in PeerStore
	go func(inv *md.AuthInvite) {
		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(inv)
		if err != nil {
			sn.Error(err, "Marshal")
		}

		sn.peerConn.Request(sn.host, id, msgBytes)
	}(&invMsg)
}

// ^ Respond to an Invitation ^ //
func (sn *Node) Respond(decision bool) {
	// @ Check Decision

	// Send Response on PeerConnection
	sn.peerConn.Authorize(decision, sn.contact, sn.peer)
}

// ^ Reset Current Queued File Metadata ^ //
func (sn *Node) ResetFile() {
	// Reset Files Slice
	sn.files = nil
	sn.files = make([]*sf.ProcessedFile, maxFileBufferSize)
}
