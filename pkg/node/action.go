package node

import (
	"math"

	sentry "github.com/getsentry/sentry-go"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multihash"
	dt "github.com/sonr-io/core/internal/data"
	md "github.com/sonr-io/core/internal/models"
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

// ^ Join Remote File with Words ^ //
func (n *Node) JoinRemote(data string) {
	// Validate
	_, err := n.transfer.JoinRemote(data)
	if err != nil {
		// Lobby non-existent
		sentry.CaptureException(err)
		n.call.Error(err, "Join Remote")
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

	// // Set User Contact
	// err := n.fs.SaveContact(newContact)
	// if err != nil {
	// 	sentry.CaptureException(err)
	// }
}

// ^ Close Ends All Network Communication ^
func (n *Node) Pause() {
	// Check if Response Is Invited
	n.transfer.Cancel(n.peer)
	dt.GetState().Pause()
}

// ^ Close Ends All Network Communication ^
func (n *Node) Resume() {
	dt.GetState().Resume()
}

// ^ Close Ends All Network Communication ^
func (n *Node) Close() {
	n.transfer.Cancel(n.peer)
	n.host.Close()
}
