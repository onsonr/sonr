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
		sn.error(err, "LinkDevice")
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
		sn.error(err, "Update")
	}
}

// ^ Process adds generates Preview with Thumbnail ^ //
func (sn *Node) Process(procBytes []byte) {
	// Initialize from Info
	request := &md.ProcessRequest{}
	err := proto.Unmarshal(procBytes, request)
	if err != nil {
		log.Println(err)
	}

	// Create Preview
	safeFile := sf.NewProcessedFile(request, sn.peer.Profile, lf.ProcessCallbacks{CallQueued: sn.call.OnQueued, CallError: sn.error})
	sn.files = append(sn.files, safeFile)
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
